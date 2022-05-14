package rpc_cache

import (
	"context"
	"github.com/jinzhu/copier"
	"suzaku/internal/domain/po_mysql"
	"suzaku/internal/domain/repo/repo_mysql"
	"suzaku/pkg/common/redis"
	"suzaku/pkg/proto/pb_com"
	"suzaku/pkg/proto/pb_user"
)

func (rpc *cacheRpcServer) GetUserInfo(_ context.Context, req *pb_user.UserInfoReq) (resp *pb_user.UserInfoResp, _ error) {
	var (
		err  error
		user *po_mysql.User
	)
	resp = &pb_user.UserInfoResp{Common: &pb_com.CommonResp{}, UserInfo: &pb_user.UserInfo{}}
	user, err = redis.GetUserInfoFromCache(req.UserId)
	if err != nil {
		resp.Common.Msg = err.Error()
		resp.Common.Code = 777
		return
	}
	if user != nil && user.UserId != "" {
		copier.Copy(resp.UserInfo, user)
		return
	}
	user, err = repo_mysql.UserRepo.GetUserByUserID(req.UserId)
	if err != nil {
		resp.Common.Msg = err.Error()
		resp.Common.Code = 777
		return
	}
	resp.UserInfo = new(pb_user.UserInfo)
	copier.Copy(resp.UserInfo, user)
	redis.SetUserInfoToCache(user)
	return
}
