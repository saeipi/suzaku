package rpc_user

import (
	"context"
	"google.golang.org/grpc"
	"suzaku/internal/rpc/rpc_category"
	"suzaku/pkg/common/config"
	pb_com "suzaku/pkg/proto/pb_com"
	pb_user "suzaku/pkg/proto/user"
)

type userRpcServer struct {
	pb_user.UnimplementedUserServer
	rpc_category.Rpc
}

func NewUserRpcServer(port int) *userRpcServer {
	return &userRpcServer{
		Rpc: rpc_category.NewRpcServer(port, config.Config.RPCRegisterName.UserName),
	}
}

func (rpc *userRpcServer) Run() {
	var (
		server *grpc.Server
	)
	server = grpc.NewServer()
	pb_user.RegisterUserServer(server, rpc)
	rpc.Rpc.RunServer(server)
}

func (rpc *userRpcServer) UserInfo(ctx context.Context, req *pb_user.UserInfoReq) (resp *pb_user.UserInfoResp, err error) {
	var (
		common = &pb_com.CommonResp{}
	)
	resp = &pb_user.UserInfoResp{Common: common}
	resp.UserId = req.UserId
	return
}
