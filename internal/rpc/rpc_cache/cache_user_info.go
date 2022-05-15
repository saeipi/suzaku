package rpc_cache

import (
	"context"
	"github.com/jinzhu/copier"
	"suzaku/internal/domain/po_mysql"
	"suzaku/internal/domain/repo/repo_mysql"
	"suzaku/pkg/common/redis"
	pb_cache "suzaku/pkg/proto/cache"
	"suzaku/pkg/proto/pb_com"
	"suzaku/pkg/proto/pb_user"
)

func (rpc *cacheRpcServer) GetUserInfo(_ context.Context, req *pb_user.UserInfoReq) (resp *pb_user.UserInfoResp, _ error) {
	var (
		user *po_mysql.User
		err  error
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

func (rpc *cacheRpcServer) UpdateUserInfo(_ context.Context, req *pb_cache.UpdateUserInfoReq) (resp *pb_cache.UpdateUserInfoResp, _ error) {
	var (
		user *po_mysql.User
		err  error
	)
	resp = &pb_cache.UpdateUserInfoResp{Common: &pb_com.CommonResp{}}
	user = new(po_mysql.User)
	switch req.UpdateType {
	case pb_cache.UPDATE_CACHE_TYPE_QUERY:
		user, err = repo_mysql.UserRepo.GetUserByUserID(req.UserId)
		if err != nil {
			resp.Common.Msg = err.Error()
			resp.Common.Code = 777
			return
		}
		copier.Copy(req.UserInfo, user)
		redis.SetUserInfoToCache(user)
	case pb_cache.UPDATE_CACHE_TYPE_VALUE:
		copier.Copy(user, req.UserInfo)
		err = redis.SetUserInfoToCache(user)
		if err != nil {
			resp.Common.Msg = err.Error()
			resp.Common.Code = 777
			return
		}
	}
	return
}

func UpdateUserInfoToCache(userInfo *pb_user.UserInfo) (resp *pb_cache.UpdateUserInfoResp) {
	var (
		client pb_cache.CacheClient
		req    *pb_cache.UpdateUserInfoReq
	)

	req = &pb_cache.UpdateUserInfoReq{
		UpdateType: pb_cache.UPDATE_CACHE_TYPE_VALUE,
		UserId:     "",
		UserInfo:   userInfo,
	}

	client = NewCacheClient()
	resp, _ = client.UpdateUserInfo(context.Background(), req)
	return
}

func UpdateUserInfoToCacheOnUserId(userId string) (resp *pb_cache.UpdateUserInfoResp) {
	var (
		client pb_cache.CacheClient
		req    *pb_cache.UpdateUserInfoReq
	)
	req = &pb_cache.UpdateUserInfoReq{
		UpdateType: pb_cache.UPDATE_CACHE_TYPE_QUERY,
		UserId:     userId,
		UserInfo:   nil,
	}
	client = NewCacheClient()
	resp, _ = client.UpdateUserInfo(context.Background(), req)
	return
}
