package etcd

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"testing"
	"time"
)

func TestCreateCli(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
		Context:     ctx,
	})
	//ca()
	if err != nil {
		// handle error!
		log.Fatalf("get etcd client init error %v", err)
	}
	defer cli.Close()

	_, err = cli.Put(ctx, "k", "v")
	if err != nil {
		// handle error!
		log.Fatalf("put------------ etcd client init error %v", err)
	}
	log.Printf("完毕")
}

func TestRegister_Reg(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:         []string{"127.0.0.1:2379"},
		DialTimeout:       5 * time.Second,
		DialKeepAliveTime: 1 * time.Minute,
		Context:           ctx,
	})
	if err != nil {
		// handle error!
		log.Fatalf("get etcd client init error %v", err)
	}
	defer cli.Close()
	register := NewRegister("svr1", cli, "127.0.0.1", 8080)
	// 主进程处理业务，另起一个协程进行定时心跳
	// Reg，RegV2 共同点如果首次注册失败(项目启动)，则直接主进程退出
	// 使用第一种注册方式，是间隔3s创建一个新的租约，租约时长为5s，将kv，与新的租约进行绑定。每次心跳即续约有2次io 生成新租约与绑定租约。容错：由于每次都是创建新的租约，网络波动与否几乎不影响流程，弊端是可能存在某一时刻一个实例注册了多个结点
	//errChan = register.Reg()
	// 使用第二种注册方式，将首次注册的租约id记录下来，将kv，与租约进行绑定。后面每次心跳只需要进行租约续期，续期5s。容错：如果网络波动导致续约失败，等待网络恢复会重新注册新的租约
	err = register.RegV2()
	if err != nil {
		log.Fatalf("注册失败 error %v", err)
	}
	for true {
		// 制定一个定时器，模拟20s后摘除服务
		t := time.NewTicker(1200 * time.Second)
		select {
		case tt := <-t.C:
			register.UnReg()
			log.Printf("服务卸载，tt:%v\n", tt)
			return
		}
	}
}
