package commands

import (
	"errors"
	"flag"
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"suzaku/examples/zaplog/logger"
	"syscall"
	"time"
)

var (
	GMainInst MainInstance
	GSignal   chan os.Signal
)

type MainInstance interface {
	Initialize() error
	RunLoop()
	Destroy()
}

func Run(inst MainInstance) {
	flag.Parse()

	if inst == nil {
		panic(errors.New("inst is nil, exit"))
		return
	}

	rand.Seed(time.Now().UTC().UnixNano())
	runtime.GOMAXPROCS(runtime.NumCPU())

	logger.Info("instance initialize...")
	err := inst.Initialize()
	logger.Info("inited")
	if err != nil {
		panic(err)
		return
	}

	GMainInst = inst

	logger.Info("instance run_loop...")
	go inst.RunLoop()

	GSignal = make(chan os.Signal, 1)
	signal.Notify(GSignal, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-GSignal
		logger.Info("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			logger.Info("instance exit...")
			inst.Destroy()
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
