package rpc_group

import (
	"google.golang.org/grpc"
	"suzaku/internal/rpc/rpc_category"
	"suzaku/pkg/common/config"
	"suzaku/pkg/factory"
	pb_group "suzaku/pkg/proto/group"
)

type groupRpcServer struct {
	pb_group.UnimplementedGroupServer
	rpc_category.Rpc
}

func NewGroupRpcServer(port int) (r *groupRpcServer) {
	return &groupRpcServer{
		Rpc: rpc_category.NewRpcServer(port, config.Config.RPCRegisterName.GroupName),
	}
}

func (rpc *groupRpcServer) Run() {
	var (
		server *grpc.Server
	)
	server = factory.NewRPCNewServer()
	pb_group.RegisterGroupServer(server, rpc)
	rpc.Rpc.RunServer(server)
}
