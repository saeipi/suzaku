package rpc_friend

import (
	"context"
	"github.com/jinzhu/copier"
	"suzaku/internal/domain/do"
	"suzaku/internal/domain/po_mysql"
	"suzaku/internal/domain/repo/repo_mysql"
	pb_friend "suzaku/pkg/proto/friend"
	"suzaku/pkg/proto/pb_com"
)

func (rpc *friendRpcServer) GetFriendRequestList(_ context.Context, req *pb_friend.GetFriendRequestListReq) (resp *pb_friend.GetFriendRequestListResp, err error) {
	var (
		query     *do.MysqlQuery
		list      []*po_mysql.FriendRequest
		totalRows int64
	)
	resp = &pb_friend.GetFriendRequestListResp{Common: &pb_com.CommonResp{}}
	query = do.NewMysqlQuery()

	switch req.Role {
	case pb_friend.FRIEND_REQUEST_ROLE_SPONSOR: // 自己是发起者
		query.Condition += " AND from_user_id=?"
		query.Params = append(query.Params, req.UserId)
	case pb_friend.FRIEND_REQUEST_ROLE_INVITED:
		query.Condition += " AND to_user_id=?"
		query.Params = append(query.Params, req.UserId)
	default:
		return
	}
	list, totalRows, err = repo_mysql.FriendRepo.GetFriendRequestList(query)
	if err != nil {
		//TODO: error
		return
	}
	resp.TotalRows = totalRows
	copier.Copy(&resp.List, list)
	return
}
