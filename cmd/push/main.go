package main

import (
	"flag"
	"suzaku/internal/push/logic"
	"sync"
)

func main() {
	rpcPort := flag.Int("port", 10700, "rpc listening port")
	flag.Parse()
	var wg sync.WaitGroup
	wg.Add(1)

	logic.Initialize(*rpcPort)
	logic.Run()

	wg.Wait()
}
