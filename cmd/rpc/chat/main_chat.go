package main

import (
	"flag"
	"suzaku/internal/rpc/rpc_chat"
)

func main() {
	rpcPort := flag.Int("port", 10300, "RpcChatPort default listen port 10300")
	flag.Parse()
	rpcServer := rpc_chat.NewChatRpcServer(*rpcPort)
	rpcServer.Run()
}
