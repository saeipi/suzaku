package main

import (
	"flag"
	"suzaku/internal/rpc/rpc_chat"
)

func main() {
	rpcPort := flag.Int("port", 10300, "rpc listening port")
	flag.Parse()
	rpcServer := rpc_chat.NewChatRpcServer(*rpcPort)
	rpcServer.Run()
}
