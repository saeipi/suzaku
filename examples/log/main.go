package main

import (
	"fmt"
	"suzaku/pkg/common/log"
)

func main() {
	log.NewLogger("suzaku", "/Users/saeipi/Desktop/logs/api.log")
	for i := 0; i < 100; i++ {
		log.Info("load config: ", "日志信息")
	}
	var input int
	fmt.Scan(&input)
}
