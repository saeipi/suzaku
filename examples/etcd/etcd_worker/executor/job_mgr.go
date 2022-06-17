package executor

import (
	"context"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"suzaku/examples/etcd/common"
	"suzaku/examples/etcd/etcd_worker/cfg"
	"time"
)

// 任务管理器
type JobMgr struct {
	client  *clientv3.Client
	kv      clientv3.KV
	lease   clientv3.Lease
	watcher clientv3.Watcher
}

var (
	SG_JOBMGR *JobMgr
)

// 初始化管理器
func InitJobMgr(cfg *cfg.Etcd) (err error) {
	var (
		config  clientv3.Config
		client  *clientv3.Client
		kv      clientv3.KV
		lease   clientv3.Lease
		watcher clientv3.Watcher
	)
	// 初始化配置
	config = clientv3.Config{
		Endpoints:   cfg.Endpoints,                                     // 集群地址
		DialTimeout: time.Duration(cfg.DialTimeout) * time.Millisecond, // 连接超时
	}
	// 建立连接
	if client, err = clientv3.New(config); err != nil {
		return
	}
	// 得到KV和Lease的API子集
	kv = clientv3.NewKV(client)
	lease = clientv3.NewLease(client)
	watcher = clientv3.NewWatcher(client)

	// 赋值单例
	SG_JOBMGR = &JobMgr{
		client:  client,
		kv:      kv,
		lease:   lease,
		watcher: watcher,
	}
	// 启动任务监听
	SG_JOBMGR.watchJobs()
	// 启动监听killer
	SG_JOBMGR.watchKiller()
	return
}

// 创建任务执行锁
func (j *JobMgr) CreateJobLock(jobName string) (jobLock *JobLock) {
	jobLock = InitJobLock(jobName, j.kv, j.lease)
	return
}

// 监听任务变化
func (j *JobMgr) watchJobs() (err error) {
	var (
		getResp            *clientv3.GetResponse
		kvpair             *mvccpb.KeyValue
		job                *common.Job
		watchStartRevision int64
		watchChan          clientv3.WatchChan
		watchResp          clientv3.WatchResponse
		watchEvent         *clientv3.Event
		jobName            string
		jobEvent           *common.JobEvent
	)

	// 1, get一下/cron/jobs/目录下的所有任务，并且获知当前集群的revision
	if getResp, err = j.kv.Get(context.TODO(), common.JOB_SAVE_DIR, clientv3.WithPrefix()); err != nil {
		return
	}
	// 当前有哪些任务
	for _, kvpair = range getResp.Kvs {
		// 反序列化json得到Job
		if job, err = common.DeserializeJob(kvpair.Value); err == nil {
			jobEvent = common.BuildJobEvent(common.JOB_EVENT_SAVE, job)
			// 同步给scheduler(调度协程)
			SG_SCHEDULER.PushJobEvent(jobEvent)
		}
	}
	// 2, 从该revision向后监听变化事件
	go func() {
		// 从GET时刻的后续版本开始监听变化
		watchStartRevision = getResp.Header.Revision + 1
		// 监听/cron/jobs/目录的后续变化
		watchChan = j.watcher.Watch(context.TODO(), common.JOB_SAVE_DIR, clientv3.WithRev(watchStartRevision), clientv3.WithPrefix())
		// 处理监听事件
		for watchResp = range watchChan {
			for _, watchEvent = range watchResp.Events {
				switch watchEvent.Type {
				case mvccpb.PUT: // 任务保存事件
					if job, err = common.DeserializeJob(watchEvent.Kv.Value); err != nil {
						continue
					}
					// 构建一个更新Event
					jobEvent = common.BuildJobEvent(common.JOB_EVENT_SAVE, job)
				case mvccpb.DELETE: // 任务被删除了
					// Delete /cron/jobs/job10
					jobName = common.ExtractJobName(string(watchEvent.Kv.Key))
					job = &common.Job{Name: jobName}
					// 构建一个删除Event
					jobEvent = common.BuildJobEvent(common.JOB_EVENT_DELETE, job)
				}
				SG_SCHEDULER.PushJobEvent(jobEvent)
			}
		}
	}()
	return
}

// 监听强杀任务通知
func (j *JobMgr) watchKiller() {
	var (
		watchChan  clientv3.WatchChan
		watchResp  clientv3.WatchResponse
		watchEvent *clientv3.Event
		jobEvent   *common.JobEvent
		jobName    string
		job        *common.Job
	)
	// 监听/cron/killer目录
	go func() {
		// 监听/cron/killer/目录的变化
		watchChan = j.watcher.Watch(context.TODO(), common.JOB_KILLER_DIR, clientv3.WithPrefix())
		// 处理监听事件
		for watchResp = range watchChan {
			for _, watchEvent = range watchResp.Events {
				switch watchEvent.Type {
				case mvccpb.PUT:
					jobName = common.ExtractKillerName(string(watchEvent.Kv.Key))
					job = &common.Job{Name: jobName}
					jobEvent = common.BuildJobEvent(common.JOB_EVENT_KILL, job)
					// 事件推给scheduler
					SG_SCHEDULER.PushJobEvent(jobEvent)
				case mvccpb.DELETE: // killer标记过期, 被自动删除

				}
			}
		}
	}()
}
