package rpc_friend

import (
	"google.golang.org/grpc"
	"suzaku/internal/rpc/rpc_category"
	"suzaku/pkg/common/config"
	"suzaku/pkg/factory"
	pb_friend "suzaku/pkg/proto/friend"
)

const (
	ErrorCodeUserIdNotExist     = 2001
	ErrorCodeSaveDatabaseFailed = 2002
)

const (
	ErrorUserIdNotExist     = "user id does not exist"
	ErrorSaveDatabaseFailed = "failed to save to database"
)

type friendRpcServer struct {
	pb_friend.UnimplementedFriendServer
	rpc_category.Rpc
}

func NewFriendRpcServer(port int) (r *friendRpcServer) {
	return &friendRpcServer{
		Rpc: rpc_category.NewRpcServer(port, config.Config.RPCRegisterName.FriendName),
	}
}

func (rpc *friendRpcServer) Run() {
	var (
		server *grpc.Server
	)
	server = factory.NewRPCNewServer()
	pb_friend.RegisterFriendServer(server, rpc)
	rpc.Rpc.RunServer(server)
}


