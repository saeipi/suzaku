package main

import (
	"flag"
	"fmt"
	"runtime"
	"suzaku/examples/etcd/etcd_worker/cfg"
	"suzaku/examples/etcd/etcd_worker/executor"
	"suzaku/examples/etcd/etcd_worker/log"
	"suzaku/examples/etcd/etcd_worker/register"
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
	// master -config ./cfg/config.yaml
	// master -h
	flag.StringVar(&configFile, "config", "./examples/etcd/etcd_worker/cfg/config.yaml", "指定配置文件")
	flag.Parse()
}

func main() {
	var (
		err error
		wg  sync.WaitGroup
	)
	wg = sync.WaitGroup{}
	wg.Add(1)

	// 初始化命令行参数
	initArgs()

	// 初始化线程
	initEnv()

	// 加载配置
	if err = cfg.InitConfig(configFile); err != nil {
		goto ERR
	}
	// 服务注册
	if err = register.InitRegister(cfg.SG_CFG.Etcd); err != nil {
		goto ERR
	}
	// 启动日志协程
	if err = log.InitLogMgr(cfg.SG_CFG.Mongodb); err != nil {
		goto ERR
	}
	// 启动执行器
	if err = executor.InitExecutor(); err != nil {
		goto ERR
	}
	// 启动调度器
	if err = executor.InitScheduler(); err != nil {
		goto ERR
	}
	// 初始化任务管理器
	if err = executor.InitJobMgr(cfg.SG_CFG.Etcd); err != nil {
		goto ERR
	}
	wg.Wait()
ERR:
	wg.Done()
	fmt.Println(err)
}
