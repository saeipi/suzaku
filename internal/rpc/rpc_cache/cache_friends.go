package rpc_cache

import (
	"context"
	"suzaku/pkg/common/redis"
	pb_cache "suzaku/pkg/proto/cache"
	"suzaku/pkg/proto/pb_com"
)

func SyncFriendInfo() {

}

func (rpc *cacheRpcServer) UpdateFriendsId(_ context.Context, req *pb_cache.UpdateFriendsIdReq) (resp *pb_cache.UpdateFriendsIdResp, _ error) {
	var (
		err error
	)
	resp = &pb_cache.UpdateFriendsIdResp{Common: &pb_com.CommonResp{}}
	switch req.ActionType {
	case pb_com.RPC_ACTION_TYPE_ADD:
		err = redis.AddFriendsId(req.UserId, req.FromUserId)
		if err != nil {
			resp.Common.Msg = err.Error()
			resp.Common.Code = 777
			return
		}
		err = redis.AddFriendsId(req.FromUserId, req.UserId)
		if err != nil {
			resp.Common.Msg = err.Error()
			resp.Common.Code = 777
			return
		}
	case pb_com.RPC_ACTION_TYPE_DELETE:
		err = redis.DeleteFriendsId(req.UserId, req.FromUserId)
		if err != nil {
			resp.Common.Msg = err.Error()
			resp.Common.Code = 777
			return
		}
		err = redis.DeleteFriendsId(req.FromUserId, req.UserId)
		if err != nil {
			resp.Common.Msg = err.Error()
			resp.Common.Code = 777
			return
		}
	}
	return
}
