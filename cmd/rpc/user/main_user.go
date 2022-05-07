package main

import (
	"flag"
	"fmt"
	"suzaku/internal/rpc/rpc_user"
)

func main() {
	rpcPort := flag.Int("port", 10100, "rpc listening port")
	flag.Parse()
	fmt.Println("start user rpc server, port: ", *rpcPort)
	rpcServer := rpc_user.NewUserRpcServer(*rpcPort)
	rpcServer.Run()
}
