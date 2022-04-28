package rpc_friend

import (
	"context"
	"google.golang.org/grpc"
	"suzaku/internal/domain/repo/repo_mysql"
	"suzaku/internal/rpc/rpc_category"
	"suzaku/pkg/common/config"
	"suzaku/pkg/factory"
	pb_friend "suzaku/pkg/proto/friend"
	pb_com "suzaku/pkg/proto/pb_com"
)

const (
	ErrorCodeUserIdNotExist = 2001
)

const (
	ErrorUserIdNotExist = "user id does not exist"
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

func (rpc *friendRpcServer) AddFriend(_ context.Context, req *pb_friend.AddFriendReq) (resp *pb_friend.AddFriendResp, err error) {
	var (
		common = &pb_com.CommonResp{}
	)
	resp = &pb_friend.AddFriendResp{Common: common}
	if _, err = repo_mysql.UserRepo.GetUserByUserID(req.ToUserId); err != nil {
		common.Code = ErrorCodeUserIdNotExist
		common.Msg = ErrorUserIdNotExist
		return
	}

	return
}
