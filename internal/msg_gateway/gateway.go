package msg_gateway

import (
	"suzaku/internal/msg_gateway/msg_handler"
	"suzaku/internal/msg_gateway/rpc_server"
	"suzaku/internal/msg_gateway/ws_server"
)

var (
	wsSvr   *ws_server.WServer
	rpcSvr  *rpc_server.RPCServer
	handler *msg_handler.MsgHandler
)

func Init(wsPort int, rpcPort int) {
	handler = msg_handler.NewMsgHandler()
	wsSvr = ws_server.NewWServer(wsPort, handler.MessageCallback)
	rpcSvr = rpc_server.NewRPCServer(rpcPort, wsSvr)
}

func Run() {
	go wsSvr.Run()
	go rpcSvr.Run()
}
