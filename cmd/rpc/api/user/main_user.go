package main

import (
	"flag"
	"suzaku/internal/rpc/api/rpc_user"
)

func main() {
	rpcPort := flag.Int("port", 10100, "rpc listening port")
	flag.Parse()
	rpcServer := rpc_user.NewUserRpcServer(*rpcPort)
	rpcServer.Run()
}
