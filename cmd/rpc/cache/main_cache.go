package main

import (
	"fmt"
	"flag"
	"suzaku/internal/rpc/rpc_cache"
)

func main() {
	rpcPort := flag.Int("port", 10688, "RpcCachePort default listen port 10688")
	flag.Parse()
	fmt.Println("start auth rpc server, port: ", *rpcPort)
	rpcServer := rpc_cache.NewCacheRpcServer(*rpcPort)
	rpcServer.Run()

}

