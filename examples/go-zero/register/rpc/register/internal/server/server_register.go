package server

import (
	"google.golang.org/grpc"
	"suzaku/examples/go-zero/proto/szkproto"
	"suzaku/examples/go-zero/register/rpc/register/internal/config"
	"suzaku/examples/go-zero/rpc_server"
	"suzaku/pkg/factory"
)

type authRpcServer struct {
	szkproto.UnimplementedAuthServer
	rpc_server.RpcServer
}

func NewAuthRpcServer(cfg *config.Config) (r *authRpcServer) {
	return &authRpcServer{
		RpcServer: rpc_server.NewRpcServer(cfg.Server.Port, cfg.Server.Name, cfg.Etcd.Schema, cfg.Etcd.Address),
	}
}

func (rpc *authRpcServer) Run() {
	var (
		server *grpc.Server
	)
	server = factory.NewRPCNewServer()
	szkproto.RegisterAuthServer(server, rpc)
	rpc.RpcServer.Run(server)
}
