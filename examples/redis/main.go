package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
)

func main() {
	var ctx = context.Background()
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{
			"127.0.0.1:7001",
			"127.0.0.1:7002",
			"127.0.0.1:7003",
			"127.0.0.1:7004",
		},
	})
	err := client.Ping(ctx).Err()
	if err != nil {
		fmt.Print(err.Error())
	}
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("k%d", i)
		val := fmt.Sprintf("v%d", i)
		err = client.Set(ctx, key, val, 0).Err()
		if err != nil {
			fmt.Print(err.Error())
		}
		val, err = client.Get(ctx, key).Result()
		if err != nil {
			fmt.Print(err.Error())
		}
		fmt.Println("key:", key, "val:", val)
	}
}
