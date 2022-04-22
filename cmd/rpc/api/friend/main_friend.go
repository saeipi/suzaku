package main

import (
	"flag"
	"suzaku/internal/rpc/api/rpc_friend"
)

func main() {
	rpcPort := flag.Int("port", 10200, "get RpcFriendPort from cmd,default 12000 as port")
	flag.Parse()
	rpcServer := rpc_friend.NewFriendRpcServer(*rpcPort)
	rpcServer.Run()
}