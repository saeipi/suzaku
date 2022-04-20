package rpc_friend

import (
	"context"
	"google.golang.org/grpc"
	"suzaku/internal/rpc/rpc_category"
	pb_friend "suzaku/pkg/proto/friend"
	"suzaku/pkg/common/config"
	pb_com "suzaku/pkg/proto/pb_com"
)

type friendRpcServer struct {
	pb_friend.UnimplementedFriendServer
	rpc_category.Rpc
}

func NewRpcFriendServer(port int) (r *friendRpcServer) {
	return &friendRpcServer{
		Rpc: rpc_category.NewRpcServer(port,config.Config.RPCRegisterName.FriendName),
	}
}

func (rpc *friendRpcServer) Run() {
	var (
		server *grpc.Server
	)
	server = grpc.NewServer()
	pb_friend.RegisterFriendServer(server, rpc)
	rpc.Rpc.RunServer(server)
}

func (rpc *friendRpcServer) AddFriend(_ context.Context, req *pb_friend.AddFriendReq) (resp *pb_friend.AddFriendResp, err error) {
	var (
		common = &pb_com.CommonResp{}
	)
	resp = &pb_friend.AddFriendResp{Common: common}
	return
}
