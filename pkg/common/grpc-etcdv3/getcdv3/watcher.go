package getcdv3

import (
	"context"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"strings"
	"sync"
	"time"
)

type Watcher struct {
	rwLock     sync.RWMutex
	client     *clientv3.Client
	kv         clientv3.KV
	watcher    clientv3.Watcher
	catalog    string
	kvs        map[string]string
	allService []string
	schema     string
	etcdAddr   string
}

func NewWatcher(catalog string, endpoints []string, schema string, etcdAddr string) (w *Watcher, err error) {
	var (
		config  clientv3.Config
		client  *clientv3.Client
		kv      clientv3.KV
		watcher clientv3.Watcher
	)

	config = clientv3.Config{
		Endpoints:   endpoints,                              // 集群地址
		DialTimeout: time.Duration(5000) * time.Millisecond, // 连接超时
	}
	// 1、建立连接
	if client, err = clientv3.New(config); err != nil {
		return
	}
	// 2、得到KV和观察者
	kv = clientv3.NewKV(client)
	watcher = clientv3.NewWatcher(client)

	w = &Watcher{
		client:     client,
		kv:         kv,
		watcher:    watcher,
		catalog:    catalog,
		kvs:        make(map[string]string),
		allService: make([]string, 0),
		schema:     schema,
		etcdAddr:   etcdAddr,
	}
	// 3、进行监听
	w.listening()
	return
}

// 监听任务变化
func (j *Watcher) listening() (err error) {
	var (
		resp               *clientv3.GetResponse
		kvpair             *mvccpb.KeyValue
		watchStartRevision int64
		watchChan          clientv3.WatchChan
		watchResp          clientv3.WatchResponse
		watchEvent         *clientv3.Event
		key                string
		value              string
	)

	// 1、get目录下的所有键值对，并且获知当前集群的revision
	if resp, err = j.kv.Get(context.TODO(), j.catalog, clientv3.WithPrefix()); err != nil {
		return
	}
	for _, kvpair = range resp.Kvs {
		key = string(kvpair.Key)
		value = string(kvpair.Value)
		j.kvs[key] = value
	}
	j.updateServices()

	// 2、从该revision向后监听变化事件
	go func() {
		// 从GET时刻的后续版本开始监听变化
		watchStartRevision = resp.Header.Revision + 1
		// 监听目录的后续变化
		watchChan = j.watcher.Watch(context.TODO(), j.catalog, clientv3.WithRev(watchStartRevision), clientv3.WithPrefix())
		// 处理监听事件
		for watchResp = range watchChan {
			for _, watchEvent = range watchResp.Events {
				switch watchEvent.Type {
				case mvccpb.PUT: // 任务保存事件
					j.rwLock.Lock()

					key = string(watchEvent.Kv.Key)
					value = string(watchEvent.Kv.Value)
					j.kvs[key] = value
					j.updateServices()

					j.rwLock.Unlock()
				case mvccpb.DELETE: // 任务被删除了
					j.rwLock.Lock()

					key = string(watchEvent.Kv.Key)
					delete(j.kvs, key)
					j.updateServices()

					j.rwLock.Unlock()
				}
			}
		}
	}()
	return
}

func (j *Watcher) updateServices() {
	var (
		maps        map[string]string
		key         string
		serviceName string
	)
	j.allService = make([]string, 0)
	maps = make(map[string]string)
	for key, _ = range j.kvs {
		serviceName = getServiceName(key)
		if _, ok := maps[serviceName]; ok == true {
			continue
		}
		maps[serviceName] = serviceName
		j.allService = append(j.allService, serviceName)
	}
}

func getServiceName(key string) (name string) {
	var (
		index int
		str   string
	)
	index = strings.LastIndex(key, "///")
	str = key[index+len("///"):]
	index = strings.Index(str, "/")
	name = str[:index]
	return
}

func (j *Watcher) GetAllConns() (conns []*grpc.ClientConn) {
	var (
		services   []string
		service    string
		clientConn *grpc.ClientConn
	)
	j.rwLock.RLock()
	services = j.allService
	j.rwLock.RUnlock()

	conns = make([]*grpc.ClientConn, 0)
	for _, service = range services {
		clientConn = GetConn(j.schema, j.etcdAddr, service)
		if clientConn == nil {
			//TODO: error
			continue
		}
		conns = append(conns, clientConn)
	}
	return
}
