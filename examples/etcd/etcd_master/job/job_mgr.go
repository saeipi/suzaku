package job

import (
	"context"
	"encoding/json"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"suzaku/examples/etcd/common"
	"suzaku/examples/etcd/etcd_master/cfg"
	"time"
)

type JobMgr struct {
	client *clientv3.Client
	kv     clientv3.KV
	lease  clientv3.Lease
}

var (
	SG_JOBMGR *JobMgr
)

func InitJobMgr(cfg *cfg.Etcd) (err error) {
	var (
		config clientv3.Config
		client *clientv3.Client
		kv     clientv3.KV
		lease  clientv3.Lease
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

	// 赋值单例
	SG_JOBMGR = &JobMgr{
		client: client,
		kv:     kv,
		lease:  lease,
	}
	return
}

// 保存任务
func (j *JobMgr) SaveJob(job *common.Job) (oldJob *common.Job, err error) {
	// 把任务保存到/cron/jobs/任务名 -> json
	var (
		jobKey   string
		jobValue []byte
		putResp  *clientv3.PutResponse
	)
	// etcd的保存key
	jobKey = common.JOB_SAVE_DIR + job.Name
	// 任务信息json
	if jobValue, err = json.Marshal(job); err != nil {
		return
	}
	// 保存到etcd
	if putResp, err = j.kv.Put(context.TODO(), jobKey, string(jobValue), clientv3.WithPrevKV()); err != nil {
		return
	}
	// 如果是更新, 那么返回旧值
	if putResp.PrevKv == nil {
		return
	}
	// 对旧值做一个反序列化
	if err = json.Unmarshal(putResp.PrevKv.Value, &oldJob); err != nil {
		err = nil
	}
	return
}

// 删除任务
func (j *JobMgr) DeleteJob(name string) (oldJob *common.Job, err error) {
	var (
		jobKey  string
		delResp *clientv3.DeleteResponse
	)
	// etcd中保存任务的key
	jobKey = common.JOB_SAVE_DIR + name
	// 从etcd中删除它
	if delResp, err = j.kv.Delete(context.TODO(), jobKey, clientv3.WithPrevKV()); err != nil {
		return
	}
	if len(delResp.PrevKvs) == 0 {
		return
	}
	// 解析一下旧值, 返回它
	if err = json.Unmarshal(delResp.PrevKvs[0].Value, &oldJob); err != nil {
		err = nil
	}
	return
}

// 列举任务
func (j *JobMgr) JobList() (jobList []*common.Job, err error) {
	var (
		dirKey  string
		getResp *clientv3.GetResponse
		kvPair  *mvccpb.KeyValue
		job     *common.Job
	)
	// 任务保存的目录
	dirKey = common.JOB_SAVE_DIR
	// 获取目录下所有任务信息
	if getResp, err = j.kv.Get(context.TODO(), dirKey, clientv3.WithPrefix()); err != nil {
		return
	}
	// 初始化数组空间
	jobList = make([]*common.Job, 0)
	// 遍历所有任务, 进行反序列化
	for _, kvPair = range getResp.Kvs {
		job = &common.Job{}
		if err = json.Unmarshal(kvPair.Value, job); err != nil {
			err = nil
			continue
		}
		jobList = append(jobList, job)
	}
	return
}

// 杀死任务
func (j *JobMgr) KillJob(name string) (err error) {
	// 更新一下key=/cron/killer/任务名
	var (
		killerKey      string
		leaseGrantResp *clientv3.LeaseGrantResponse
		leaseId        clientv3.LeaseID
	)
	// 通知worker杀死对应任务
	killerKey = common.JOB_KILLER_DIR + name

	// 让worker监听到一次put操作, 创建一个租约让其稍后自动过期即可（1秒之后过期）
	if leaseGrantResp, err = j.lease.Grant(context.TODO(), 1); err != nil {
		return
	}
	// 租约ID
	leaseId = leaseGrantResp.ID
	// 设置killer标记
	if _, err = j.kv.Put(context.TODO(), killerKey, "", clientv3.WithLease(leaseId)); err != nil {
		return
	}
	return
}
