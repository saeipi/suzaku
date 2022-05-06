package main

import (
	"flag"
	"fmt"
	"suzaku/internal/rpc/api/rpc_group"
)

func main() {
	rpcPort := flag.Int("port", 10500, "get RpcGroupPort from cmd,default 16000 as port")
	flag.Parse()
	fmt.Println("start group rpc server, port: ", *rpcPort)
	rpcServer := rpc_group.NewGroupRpcServer(*rpcPort)
	rpcServer.Run()
}
