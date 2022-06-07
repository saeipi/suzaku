package main

import (
	"suzaku/examples/zaplog/logger"
	"time"
)

func main() {
	for i := 0; i < 100; i++ {
		time.Sleep(2 * time.Second)
		logger.Debug("log Debug:", "Debug")
		logger.Infof("%s", "log Infof")
		logger.Info("log Info")
		logger.Warn("log Warn")
		logger.Error("log Error")
		logger.Infow("log Infow",
			"url", "http://www.baidu.com",
			"attempt", 3,
			"backoff", time.Second,
		)
	}
}
