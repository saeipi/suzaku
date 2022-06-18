package commands

import (
	"errors"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"runtime"
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

	fmt.Println("instance initialize...")
	err := inst.Initialize()
	fmt.Println("inited")
	if err != nil {
		panic(err)
		return
	}

	GMainInst = inst

	fmt.Println("instance run_loop...")
	go inst.RunLoop()

	GSignal = make(chan os.Signal, 1)
	signal.Notify(GSignal, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-GSignal
		fmt.Println("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			fmt.Println("instance exit...")
			inst.Destroy()
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
