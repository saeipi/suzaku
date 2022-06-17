package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

// chan 套娃 不确定谁会收到

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover 捕获到了异常", err)
		}
	}()
	ch1 := make(chan bool, 10)
	ch2 := make(chan bool, 10)
	ch1 <- true
	ch1 <- false
	ch2 <- true
	ch2 <- false
	fmt.Println(<-ch1)
	close(ch1)
	fmt.Println(<-ch1)
	close(ch2)
	ch2 <- true

	return

	var signalChan = make(chan bool)
	wg.Add(1)
	go worker(signalChan)
	time.Sleep(time.Second * 5)
	signalChan <- true
	wg.Wait()

	var input int
	fmt.Scan(&input)
}

func worker(signalChan chan bool) {
	defer wg.Done()

	go subWorker(signalChan)
	time.Sleep(time.Second * 1)

	var (
		isExit bool
	)
	for {
		select {
		case isExit = <-signalChan:
			if isExit {
				goto EXIT
			}
		default:
		}
		/*
			isExit = <-signalChan
			if isExit {
				break EXIT
			}
		*/
		time.Sleep(time.Second * 1)
		fmt.Println("01 循环中……")
	}
EXIT:
	fmt.Println("01 循环中 【退出】")
}

func subWorker(signalChan chan bool) {
	var (
		isExit bool
	)
	for {
		select {
		case isExit = <-signalChan:
			if isExit {
				goto EXIT
			}
		default:
		}
		/*
			isExit = <-signalChan
			if isExit {
				break EXIT
			}
		*/
		time.Sleep(time.Second * 1)
		fmt.Println("02 循环中……")
	}
EXIT:
	fmt.Println("02 循环中 【退出】")
}
