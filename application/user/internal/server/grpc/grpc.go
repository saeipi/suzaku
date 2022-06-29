package grpc

import (
	"google.golang.org/grpc"
	"suzaku/application/user/internal/config"
	"suzaku/internal/rpc/rpc_category"
	"suzaku/pkg/factory"
	pb_user "suzaku/pkg/proto/pb_user"
)

const (
	ErrorCodeUserIdNotExist = 2001
)

const (
	ErrorUserIdNotExist = "user id does not exist"
)

type UserRpcServer struct {
	pb_user.UnimplementedUserServer
	rpc_category.Rpc
}

func NewUserRpcServer(cfg config.RPCServer) *UserRpcServer {
	return &UserRpcServer{
		Rpc: rpc_category.NewRpcServer(cfg.Port, cfg.Name),
	}
}

func (rpc *UserRpcServer) Run() {
	var (
		server *grpc.Server
	)
	server = factory.NewRPCNewServer()
	pb_user.RegisterUserServer(server, rpc)
	rpc.Rpc.RunServer(server)
}
