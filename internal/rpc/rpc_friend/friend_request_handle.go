package rpc_friend

import (
	"context"
	"suzaku/internal/domain/po_mysql"
	"suzaku/internal/domain/repo/repo_mysql"
	pb_friend "suzaku/pkg/proto/friend"
	"suzaku/pkg/proto/pb_com"
)

func (rpc *friendRpcServer) HandleFriendRequest(_ context.Context, req *pb_friend.HandleFriendRequestReq) (resp *pb_friend.HandleFriendRequestResp, _ error) {
	switch req.HandleResult {
	case int32(pb_friend.HANDLE_FRIEND_REQUEST_RESULT_APPROVE):
		return rpc.approve(req)
	case int32(pb_friend.HANDLE_FRIEND_REQUEST_RESULT_RESOLUTELY):
		return rpc.resolutely(req)
	}
	return
}

func (rpc *friendRpcServer) approve(req *pb_friend.HandleFriendRequestReq) (resp *pb_friend.HandleFriendRequestResp, _ error) {
	var (
		err    error
		friend *po_mysql.Friend
	)
	resp = &pb_friend.HandleFriendRequestResp{Common: &pb_com.CommonResp{}}
	friend, err = repo_mysql.FriendRepo.IsFriend(req.UserId, req.FromUserId)
	if err != nil {
		//TODO:Error
		resp.Common.Code = 777
		resp.Common.Msg = err.Error()
		return
	}
	if friend.OwnerUserId != "" {
		//TODO:已经是好友
		return
	}
	err = repo_mysql.FriendRepo.ApproveFriendRequest(req)
	if err != nil {
		//TODO:Error
		resp.Common.Code = 777
		resp.Common.Msg = err.Error()
		return
	}
	return
}

func (rpc *friendRpcServer) resolutely(req *pb_friend.HandleFriendRequestReq) (resp *pb_friend.HandleFriendRequestResp, _ error) {
	var (
		err error
	)
	resp = &pb_friend.HandleFriendRequestResp{Common: &pb_com.CommonResp{}}
	err = repo_mysql.FriendRepo.UpdateFriendRequest(req)
	if err != nil {
		//TODO:Error
		resp.Common.Code = 777
		resp.Common.Msg = err.Error()
		return
	}
	return
}
