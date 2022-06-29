package server

import (
	"flag"
	"suzaku/application/user/internal/config"
	"suzaku/application/user/internal/server/grpc"
	"suzaku/pkg/common/logger"
	"suzaku/pkg/utils"
)

var appConfig = flag.String("a", "./configs/user.yaml", "the config file")
var loggerConfig = flag.String("l", "./configs/logger.yaml", "the logger config file")

type Server struct {
	rpcSvr *grpc.UserRpcServer
}

func New() *Server {
	return new(Server)
}

func (s *Server) Initialize() (err error) {
	var (
		svrCfg = new(config.Config)
	)
	err = utils.YamlToStruct(*appConfig, svrCfg)
	if err != nil {
		panic(err)
	}
	logger.New(*loggerConfig, svrCfg.Name)
	s.rpcSvr = grpc.NewUserRpcServer(svrCfg.RPCServer)
	return
}

func (s *Server) RunLoop() {
	s.rpcSvr.Run()
}

func (s *Server) Destroy() {

}
