package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	request(ctx)

	var input int
	fmt.Scan(&input)
}

func request(ctx context.Context) {
	go func() {
		for {
			time.Sleep(time.Millisecond * 200)
			fmt.Println("---- 循环中 ----")
			select {
			case <-ctx.Done():
				fmt.Println("---- 超时 ----")
				return
			default:
			}
		}
	}()
}
