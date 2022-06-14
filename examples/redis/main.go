package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{"10.0.115.140:6381"},
	})
	err := client.Ping().Err()
	if err != nil {
		fmt.Print(err.Error())
	}
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("k%d", i)
		val := fmt.Sprintf("v%d", i)
		err = client.Set(key, val, 0).Err()
		if err != nil {
			fmt.Print(err.Error())
		}
		val, err = client.Get(key).Result()
		if err != nil {
			fmt.Print(err.Error())
		}
		fmt.Println("key:", key, "val:", val)
	}
}
