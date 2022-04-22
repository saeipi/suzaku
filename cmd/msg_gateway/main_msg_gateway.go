package main

import (
	"flag"
	"suzaku/internal/msg_gateway"
	"sync"
)

func main() {
	rpcPort := flag.Int("rpc_port", 10400, "rpc listening port")
	wsPort := flag.Int("ws_port", 17778, "ws listening port")
	flag.Parse()

	var wg sync.WaitGroup
	wg.Add(1)
	msg_gateway.Init(*wsPort, *rpcPort)
	msg_gateway.Run()
	wg.Wait()
}
