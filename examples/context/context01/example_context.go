package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait()

	var input int
	fmt.Scan(&input)
}

func worker(ctx context.Context) {
	defer wg.Done()

	go subWorker(ctx)

	for {
		select {
		case <-ctx.Done():
			goto EXIT
		default:
		}
		/*
			isExit = <-signalChan
			if isExit {
				break EXIT
			}
		*/
		time.Sleep(time.Second * 1)
		fmt.Println("01循环中……")
	}
EXIT:
	fmt.Println("01循环中 【退出】")
}

func subWorker(ctx context.Context) {

	for {
		select {
		case <-ctx.Done():
			goto EXIT
		default:
		}
		/*
			isExit = <-signalChan
			if isExit {
				break EXIT
			}
		*/
		time.Sleep(time.Second * 1)
		fmt.Println("02循环中……")
	}
EXIT:
	fmt.Println("02循环中 【退出】")
}
