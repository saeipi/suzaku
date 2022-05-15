package rpc_cache

import (
	"suzaku/internal/domain/po_mysql"
	"suzaku/pkg/common/redis"
)

func (rpc *cacheRpcServer) GetFriendsIdList(userId string) {
	var (
		err     error
		fids    []string
		friends []*po_mysql.User
		user    *po_mysql.User
	)
	friends = make([]*po_mysql.User, 0)
	fids = redis.GetFriendsIdList(userId)
	for _, fid := range fids {
		user, err = redis.GetUserInfoFromCache(fid)
		if err != nil {
			return
		}
		if user == nil {
			continue
		}
		friends = append(friends, user)
	}
}
