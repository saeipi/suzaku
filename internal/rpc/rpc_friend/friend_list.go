package rpc_friend

import (
	"context"
	"github.com/jinzhu/copier"
	"suzaku/internal/domain/do"
	"suzaku/internal/domain/repo/repo_mysql"
	pb_friend "suzaku/pkg/proto/friend"
	"suzaku/pkg/proto/pb_com"
)

func (rpc *friendRpcServer) GetFriendList(_ context.Context, req *pb_friend.FriendListReq) (resp *pb_friend.FriendListResp, _ error) {
	var (
		friends   []*do.FriendInfo
		totalRows int64
		err       error
	)
	resp = &pb_friend.FriendListResp{Common: &pb_com.CommonResp{}}
	friends, totalRows, err = repo_mysql.FriendRepo.FriendList(req)
	if err != nil {
		resp.Common.Code = 777
		resp.Common.Msg = err.Error()
		return
	}
	resp.TotalRows = totalRows
	copier.Copy(&resp.MemberList, &friends)
	return
}
