package rpc_user

import (
	"context"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc"
	"suzaku/internal/domain/po_mysql"
	"suzaku/internal/domain/repo/repo_mysql"
	"suzaku/internal/rpc/rpc_category"
	"suzaku/pkg/common/config"
	"suzaku/pkg/factory"
	pb_com "suzaku/pkg/proto/pb_com"
	pb_user "suzaku/pkg/proto/user"
)

const (
	ErrorCodeUserIdNotExist = 2001
)

const (
	ErrorUserIdNotExist = "user id does not exist"
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
	server = factory.NewRPCNewServer()
	pb_user.RegisterUserServer(server, rpc)
	rpc.Rpc.RunServer(server)
}

func (rpc *userRpcServer) UserInfo(ctx context.Context, req *pb_user.UserInfoReq) (resp *pb_user.UserInfoResp, err error) {
	var (
		common = &pb_com.CommonResp{}
		user   *po_mysql.User
	)
	resp = &pb_user.UserInfoResp{Common: common,UserInfo:&pb_user.UserInfo{}}
	user, err = repo_mysql.UserRepo.GetUserByUserID(req.UserId)
	if err != nil {
		common.Code = ErrorCodeUserIdNotExist
		common.Msg = ErrorUserIdNotExist
		return
	}
	copier.Copy(resp.UserInfo, user)
	return
}

func (rpc *userRpcServer) EditUserInfo(_ context.Context, req *pb_user.EditUserInfoReq) (resp *pb_com.CommonResp, _ error) {
	var (
		err error
	)
	resp = &pb_com.CommonResp{}
	err = repo_mysql.UserRepo.EditUserInfo(req)
	if err != nil {
		resp.Code = 777
		resp.Msg = err.Error()
		return
	}
	return
}
