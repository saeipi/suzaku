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

func (rpc *friendRpcServer) GetFriendRequestList(_ context.Context, req *pb_friend.GetFriendRequestListReq) (resp *pb_friend.GetFriendRequestListResp, _ error) {
	var (
		sel       *do.MysqlSelect
		list      []*po_mysql.FriendRequest
		totalRows int64
		err       error
	)
	resp = &pb_friend.GetFriendRequestListResp{Common: &pb_com.CommonResp{}}
	sel = do.NewMysqlSelect()

	switch req.Role {
	case pb_friend.FRIEND_REQUEST_ROLE_SPONSOR: // 自己是发起者
		sel.Query += " AND from_user_id=?"
		sel.Args = append(sel.Args, req.UserId)
	case pb_friend.FRIEND_REQUEST_ROLE_INVITED:
		sel.Query += " AND to_user_id=?"
		sel.Args = append(sel.Args, req.UserId)
	default:
		return
	}
	list, totalRows, err = repo_mysql.FriendRepo.GetFriendRequestList(sel)
	if err != nil {
		//TODO: error
		return
	}
	resp.TotalRows = totalRows
	copier.Copy(&resp.List, list)
	return
}
