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

	for num := range loop(ctx) {
		fmt.Println(num)
		if num == 5 {
			break
		}
	}

	time.Sleep(time.Second * 2)
	cancel()
	fmt.Println("结束")

	var input int
	fmt.Scan(&input)
}

func loop(ctx context.Context) <-chan int {
	dst := make(chan int)
	num := 1
	go func() {
		for {
			time.Sleep(time.Millisecond * 200)
			select {
			case <-ctx.Done():
				return
			case dst <- num:
				fmt.Println("----", num, "----")
				num++
			default:

			}
		}
	}()
	return dst
}
