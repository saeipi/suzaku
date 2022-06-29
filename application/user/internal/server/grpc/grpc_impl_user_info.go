package grpc

import (
	"context"
	"github.com/jinzhu/copier"
	"suzaku/internal/domain/po_mysql"
	"suzaku/internal/domain/repo/repo_mysql"
	"suzaku/internal/rpc/rpc_cache"
	"suzaku/pkg/proto/pb_com"
	"suzaku/pkg/proto/pb_user"
)

func (rpc *UserRpcServer) UserInfo(ctx context.Context, req *pb_user.UserInfoReq) (resp *pb_user.UserInfoResp, err error) {
	var (
		common = &pb_com.CommonResp{}
		user   *po_mysql.User
	)
	resp = &pb_user.UserInfoResp{Common: common, UserInfo: &pb_user.UserInfo{}}
	user, err = repo_mysql.UserRepo.GetUserByUserID(req.UserId)
	if err != nil {
		common.Code = ErrorCodeUserIdNotExist
		common.Msg = ErrorUserIdNotExist
		return
	}
	copier.Copy(resp.UserInfo, user)
	return
}

func (rpc *UserRpcServer) EditUserInfo(_ context.Context, req *pb_user.EditUserInfoReq) (resp *pb_com.CommonResp, _ error) {
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
	go rpc_cache.UpdateUserInfoToCacheOnUserId(req.UserId)
	return
}
