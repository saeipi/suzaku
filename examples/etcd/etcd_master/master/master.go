package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"suzaku/examples/etcd/etcd_master/cfg"
	"suzaku/examples/etcd/etcd_master/job"
	"suzaku/examples/etcd/etcd_master/log"
	"suzaku/examples/etcd/etcd_master/server"
	"suzaku/examples/etcd/etcd_master/worker"
	"sync"
)

var (
	configFile string
)

func initEnv() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

// 解析命令行参数
func initArgs() {
	// master -config ./config/config.yaml
	// master -h
	flag.StringVar(&configFile, "config", "./examples/etcd/etcd_master/cfg/config.yaml", "指定配置文件")
	//flag.StringVar(&configFile, "config", "../../../../conf/dev.yaml", "指定配置文件")
	flag.Parse()
}

// go build main.go
func main() {
	str, _ := os.Getwd()
	println(str)

	var (
		err error
		wg  sync.WaitGroup
	)
	wg = sync.WaitGroup{}
	wg.Add(1)

	// 1 初始化命令行参数
	initArgs()

	// 2 初始化环境
	initEnv()

	// 3 加载配置
	if err = cfg.InitConfig(configFile); err != nil {
		goto ERR
	}
	// 4 初始化服务发现模块
	if err = worker.InitWorkerMgr(cfg.SG_CFG.Etcd); err != nil {
		goto ERR
	}
	// 5 日志管理器
	if err = log.InitLogMgr(cfg.SG_CFG.Mongodb); err != nil {
		goto ERR
	}
	// 6 任务管理器
	if err = job.InitJobMgr(cfg.SG_CFG.Etcd); err != nil {
		goto ERR
	}
	// 7 启动Api HTTP服务
	if err = server.InitRouter(cfg.SG_CFG.Server); err != nil {
		goto ERR
	}
	wg.Wait()
ERR:
	wg.Done()
	fmt.Println(err)
}
