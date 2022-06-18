package server

import (
	"flag"
	"suzaku/application/msg_gateway/internal/config"
	"suzaku/application/msg_gateway/internal/server/grpc"
	"suzaku/application/msg_gateway/internal/server/websocket/ws_server"
	"suzaku/pkg/common/logger"
	"suzaku/pkg/utils"
)

var configFile = flag.String("f", "./configs/msg_gateway.yaml", "the config file")

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
	err = utils.YamlToStruct(*configFile, svrCfg)
	if err != nil {
		panic(err)
	}
	err = utils.YamlToStruct("./configs/logger.yaml", logCfg)
	if err != nil {
		panic(err)
	}

	logCfg.Directory = svrCfg.Name
	logger.InitLogger(logCfg)

	s.wsSvr = ws_server.NewServer(svrCfg.WsServer.Port, nil)
	s.rpcSvr = grpc.NewRPCServer(svrCfg.RPCServer, s.wsSvr)

	go func() {
		s.wsSvr.Run()
		s.rpcSvr.Run()
	}()
	logger.Info("初始化完毕")
	return
}

func (s *Server) RunLoop() {

}
func (s *Server) Destroy() {

}