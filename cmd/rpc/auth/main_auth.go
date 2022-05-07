package main

import (
	"flag"
	"suzaku/internal/rpc/rpc_auth"
)

func main() {
	rpcPort := flag.Int("port", 10600, "RpcToken default listen port 10800")
	flag.Parse()
	rpcServer := rpc_auth.NewAuthRpcServer(*rpcPort)
	rpcServer.Run()
}
