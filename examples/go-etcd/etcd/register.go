package etcd

import (
	"context"
	"encoding/json"
	uuid "github.com/satori/go.uuid"
	"go.etcd.io/etcd/api/v3/v3rpc/rpctypes"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"strings"
	"time"
)

// Register 定义一个注册器
type Register struct {
	client3   *clientv3.Client
	stop      chan bool
	interval  time.Duration
	leaseTime int64
	leaseID   clientv3.LeaseID
	node      *NodeInfo
}

// RandomStr 随机生成一个字符串用户表示服务的nodeID
func RandomStr(len int) string {
	nUid := uuid.NewV4().String()
	str := strings.Replace(nUid, "-", "", -1)
	if len < 0 || len >= 32 {
		return str
	}
	return str[:len]
}

// NewRegister 注册器构造函数
func NewRegister(svrName string, cli *clientv3.Client, ipAddr string, port int) *Register {
	return &Register{
		client3:   cli,
		interval:  3 * time.Second,
		leaseTime: 5,
		stop:      make(chan bool, 1),
		node: &NodeInfo{
			SvrID:   RandomStr(32),
			SvrName: svrName,
			IpAddr:  ipAddr,
			Port:    port,
		},
	}
}

// Reg 服务注册 v1版本 配合租约一起使用，每次心跳检测时候，是都新建一个租约将key绑定到新的租约上
func (r *Register) Reg() chan error {
	// 启动协程检测心跳, 定期续租
	errChan := make(chan error, 1)
	go func() {
		t := time.NewTicker(r.interval)
		err := r.doReg()
		if err != nil {
			errChan <- err
			log.Println("注册失败，退出")
			return
		}
		for {
			select {
			//case ttl := <-t.C:
			//	r.doReg()
			//	log.Printf("heartbeat check for etcd k:%v t:[%v]\n", r.node.SvrName, ttl)
			case <-t.C:
				r.doReg()
				//log.Printf("heartbeat check for etcd k:%v t:[%v]\n", r.node.SvrName, ttl)
			case <-r.stop:
				log.Println("程序退出，心跳检测协程结束")
				return
			}
		}
	}()
	return errChan
}

// v1 版本，心跳续约如果失败，直接返回，如果网络一旦恢复，可以直接恢复注册 一次心跳检测需要 新建一个租约新put一次 2次io
func (r *Register) doReg() error {
	key := getRegKey(r.node.SvrName, r.node.SvrID)
	cli := r.client3
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	resp, err := cli.Grant(ctx, r.leaseTime)
	if err != nil {
		log.Printf("get grant err:%v\n", err)
		return err
	}
	//log.Printf("new lease:%v ttl:%v\n", resp.ID, resp.TTL)
	r.leaseID = resp.ID
	nodeBy, err := json.Marshal(r.node)
	if err != nil {
		log.Printf("data:%v to json string err:%v", r.node, err)
		return err
	}
	log.Printf("key:%v,node:%v, val:%v\n", key, r.node, string(nodeBy))
	_, err = cli.Put(ctx, key, string(nodeBy), clientv3.WithLease(resp.ID))
	if err != nil {
		log.Printf("put register key err:%v\n", err)
		return err
	}
	return nil
}

// UnReg 服务卸载
func (r *Register) UnReg() {
	r.stop <- true
	r.stop = make(chan bool, 1)
	key := getRegKey(r.node.SvrName, r.node.SvrID)
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	if _, err := r.client3.Delete(ctx, key); err != nil {
		log.Fatalln(err)
	}
	log.Printf("服务：%v 摘除成功\n", r.node.SvrName)
}

// RegV2 服务注册v2版本 服务put kv的时候配合租约一起使用，心跳检测只是去续约租约
func (r *Register) RegV2() error {
	err := r.doReg()
	if err != nil {
		log.Println("注册失败，退出")
		return err
	}
	go func() {
		t := time.NewTicker(r.interval)
		for {
			select {
			case ttl := <-t.C:
				ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
				_, err := r.client3.KeepAliveOnce(ctx, r.leaseID)
				if err != nil && err == rpctypes.ErrLeaseNotFound {
					log.Println("某次超时，导致租约丢失，重新注册")
					r.doReg()
				} else if err != nil {
					log.Printf("租约续期失败：%v t:[%v]\n", err, ttl)
				}
				//log.Printf("heartbeat check for etcd r:%+v t:[%v]\n", r, ttl)
			case <-r.stop:
				log.Println("程序退出，心跳检测协程结束")
				return
			}
		}
	}()
	return nil
}

