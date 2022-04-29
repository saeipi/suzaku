package rpc_friend

import (
	"context"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"suzaku/internal/domain/po_mysql"
	"suzaku/internal/domain/repo/repo_mysql"
	"suzaku/pkg/constant"
	pb_friend "suzaku/pkg/proto/friend"
	"suzaku/pkg/proto/pb_com"
	pb_ws "suzaku/pkg/proto/pb_ws"
)

func (rpc *friendRpcServer) AddFriend(_ context.Context, req *pb_friend.AddFriendReq) (resp *pb_friend.AddFriendResp, err error) {
	var (
		common        = &pb_com.CommonResp{}
		friendRequest *po_mysql.FriendRequest
	)
	resp = &pb_friend.AddFriendResp{Common: common}
	if _, err = repo_mysql.UserRepo.GetUserByUserID(req.ToUserId); err != nil {
		common.Code = ErrorCodeUserIdNotExist
		common.Msg = ErrorUserIdNotExist
		return
	}
	friendRequest = &po_mysql.FriendRequest{
		FromUserId: req.FromUserId,
		ToUserId:   req.ToUserId,
		ReqMsg:     req.ReqMsg,
	}
	err = repo_mysql.FriendRepo.SaveFriendRequest(friendRequest)
	if err != nil {
		common.Code = ErrorCodeSaveDatabaseFailed
		common.Msg = ErrorSaveDatabaseFailed
		return
	}
	return
}

func (rpc *friendRpcServer) friendApplicationNotification(req *pb_friend.AddFriendReq) {
	var (
		tips *pb_ws.FriendApplicationTips
	)
	tips = &pb_ws.FriendApplicationTips{FromToUserId: &pb_ws.FromToUserID{
		FromUserId: req.FromUserId,
		ToUserId:   req.ToUserId,
	}}
	tips = tips
}

func (rpc *friendRpcServer) friendNotification(req *pb_friend.AddFriendReq, contentType int32, msg proto.Message) {
	var (
		tips             pb_ws.TipsComm
		marshaler        jsonpb.Marshaler
		fromUserNickname string
		toUserNickname   string
		err              error
	)
	tips.Detail, err = proto.Marshal(msg)
	if err != nil {
		//TODO:错误处理
		return
	}
	marshaler = jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  false,
		EmitDefaults: false,
	}
	tips.JsonDetail, _ = marshaler.MarshalToString(msg)

	fromUserNickname, toUserNickname, err = repo_mysql.UserRepo.GetFromToUserNickname(req.FromUserId, req.ToUserId)
	if err != nil {
		//TODO:错误处理
		return
	}
	switch contentType {
	case constant.FriendApplicationApprovedNotification: // add_friend_response
		tips.DefaultTips = toUserNickname + "同意了你的好友申请"
	case constant.FriendApplicationRejectedNotification: // add_friend_response
		tips.DefaultTips = toUserNickname + "拒绝了你的好友申请"
	case constant.FriendApplicationNotification: // add_friend
		tips.DefaultTips = fromUserNickname + "请求添加你为好友"
	default:
		return
	}
}
