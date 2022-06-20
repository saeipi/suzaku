package server

import (
	"flag"
	"suzaku/application/msg_gateway/internal/config"
	"suzaku/application/msg_gateway/internal/server/grpc"
	"suzaku/application/msg_gateway/internal/server/websocket/msg_handler"
	"suzaku/application/msg_gateway/internal/server/websocket/ws_server"
	"suzaku/pkg/common/logger"
	"suzaku/pkg/utils"
)

var appConfig = flag.String("a", "./configs/msg_gateway.yaml", "the config file")
var loggerConfig = flag.String("l", "./configs/logger.yaml", "the config file")

type Server struct {
	wsSvr  *ws_server.WServer
	rpcSvr *grpc.RPCServer
}

func New() *Server {
	return new(Server)
}

func (s *Server) Initialize() (err error) {
	var (
		svrCfg = new(config.Config)
		logCfg = new(logger.Zap)
	)
	err = utils.YamlToStruct(*appConfig, svrCfg)
	if err != nil {
		panic(err)
	}
	err = utils.YamlToStruct(*loggerConfig, logCfg)
	if err != nil {
		panic(err)
	}

	logCfg.Directory = svrCfg.Name
	logger.InitLogger(logCfg)

	s.wsSvr = ws_server.NewServer(&svrCfg.WsServer, msg_handler.NewMsgHandler(&svrCfg.RPCServer))
	s.rpcSvr = grpc.NewRPCServer(svrCfg.RPCServer, s.wsSvr)
	return
}

func (s *Server) RunLoop() {
	s.wsSvr.Run()
	go s.rpcSvr.Run()
}

func (s *Server) Destroy() {

}
