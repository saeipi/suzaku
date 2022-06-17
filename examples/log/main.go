package main

import (
	"fmt"
	"suzaku/pkg/common/log"
)

func main() {
	log.NewLogger("suzaku", "./logs/api.log")
	for i := 0; i < 100; i++ {
		log.Info("1991", "日志信息")
	}
	var input int
	fmt.Scan(&input)
}
