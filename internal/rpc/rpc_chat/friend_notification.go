package rpc_chat

import (
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"suzaku/internal/domain/do"
	"suzaku/internal/domain/repo/repo_mysql"
	"suzaku/pkg/constant"
	pb_friend "suzaku/pkg/proto/friend"
	pb_ws "suzaku/pkg/proto/pb_ws"
)

func FriendApplicationNotification(req *pb_friend.AddFriendReq) {
	var (
		tips *pb_ws.FriendApplicationTips
	)
	tips = &pb_ws.FriendApplicationTips{FromToUserId: &pb_ws.FromToUserID{
		FromUserId: req.FromUserId,
		ToUserId:   req.ToUserId,
	}}
	friendNotification(req, constant.FriendApplicationNotification, tips)
}

func friendNotification(req *pb_friend.AddFriendReq, contentType int32, msg proto.Message) {
	var (
		tips             pb_ws.TipsComm
		marshaler        jsonpb.Marshaler
		fromUserNickname string
		toUserNickname   string
		notification     *do.NotificationMsg
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
	notification = &do.NotificationMsg{
		SendID:      req.FromUserId,
		RecvID:      req.ToUserId,
		Content:     nil,
		MsgFrom:     constant.SysMsgType,
		ContentType: contentType,
		SessionType: constant.SingleChatType,
		OperationID: req.OperationId,
	}
	notification.Content, err = proto.Marshal(&tips)
	if err != nil {
		//TODO:错误处理
		return
	}
	Notification(notification)
}
