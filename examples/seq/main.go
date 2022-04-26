package main

import (
	"fmt"
	"suzaku/pkg/common/redis"
)

func main() {
	var (
		seq    uint64
		minSeq uint64
		maxSeq uint64
		userID string
		err    error
	)
	userID = "1001"
	for i := 0; i < 10; i++ {
		seq, err = redis.IncrUserSeq(userID)
		if err != nil {
			fmt.Println("err:", err)
			continue
		}
		maxSeq, err = redis.GetUserMaxSeq(userID)
		minSeq, err = redis.GetUserMinSeq(userID)

		fmt.Println(seq, minSeq, maxSeq)
	}
}
