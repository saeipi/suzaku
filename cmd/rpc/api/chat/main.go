package main

import (
	"flag"
	"suzaku/internal/rpc/api/rpc_chat"
)

func main() {
	rpcPort := flag.Int("port", 10300, "rpc listening port")
	flag.Parse()
	rpcServer := rpc_chat.NewRpcChatServer(*rpcPort)
	rpcServer.Run()
}
