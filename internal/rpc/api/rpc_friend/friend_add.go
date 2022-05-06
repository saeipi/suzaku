package rpc_friend

import (
	"context"
	"suzaku/internal/domain/po_mysql"
	"suzaku/internal/domain/repo/repo_mysql"
	"suzaku/internal/rpc/api/rpc_chat"
	pb_friend "suzaku/pkg/proto/friend"
	"suzaku/pkg/proto/pb_com"
	"time"
)

func (rpc *friendRpcServer) AddFriend(_ context.Context, req *pb_friend.AddFriendReq) (resp *pb_friend.AddFriendResp, _ error) {
	var (
		common        = &pb_com.CommonResp{}
		friendRequest *po_mysql.FriendRequest
		toUser        *po_mysql.User
		err           error
	)
	resp = &pb_friend.AddFriendResp{Common: common}
	if toUser, err = repo_mysql.UserRepo.GetUserBySzkID(req.ToSzkId); err != nil {
		common.Code = ErrorCodeUserIdNotExist
		common.Msg = ErrorUserIdNotExist
		return
	}
	req.ToUserId = toUser.UserId
	friendRequest = &po_mysql.FriendRequest{
		FromUserId:   req.FromUserId,
		ToUserId:     req.ToUserId,
		ReqMsg:       req.ReqMsg,
		HandleUserId: req.OperationId,
		ReqTs:        time.Now().Unix(),
	}
	err = repo_mysql.FriendRepo.SaveFriendRequest(friendRequest)
	if err != nil {
		common.Code = ErrorCodeSaveDatabaseFailed
		common.Msg = ErrorSaveDatabaseFailed
		return
	}
	rpc_chat.FriendApplicationNotification(req)
	return
}
