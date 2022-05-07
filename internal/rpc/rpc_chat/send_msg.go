package rpc_chat

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"suzaku/pkg/common/config"
	"suzaku/pkg/common/redis"
	"suzaku/pkg/constant"
	"suzaku/pkg/factory"
	pb_chat "suzaku/pkg/proto/chart"
	pb_group "suzaku/pkg/proto/group"
	"suzaku/pkg/proto/pb_ws"
	"suzaku/pkg/utils"
)

type MsgCallBackReq struct {
	SendID       string `json:"send_id"`
	RecvID       string `json:"recv_id"`
	Content      string `json:"content"`
	SendTs       int64  `json:"send_ts"`
	MsgFrom      int32  `json:"msg_from"`
	ContentType  int32  `json:"content_type"`
	SessionType  int32  `json:"session_type"`
	PlatformID   int32  `json:"sender_platform_id"`
	MsgID        string `json:"msg_id"`
	IsOnlineOnly bool   `json:"is_online_only"`
}

type MsgCallBackResp struct {
	ErrCode         int32  `json:"err_code"`
	ErrMsg          string `json:"err_msg"`
	ResponseErrCode int32  `json:"response_err_code"`
	ResponseResult  struct {
		ModifiedMsg string `json:"modified_msg"`
		Ext         string `json:"ext"`
	}
}

func (rpc *chatRpcServer) encapsulateMsgData(msg *pb_ws.MsgData) {
	msg.ServerMsgId = GetMsgID(msg.SendId)
	msg.SendTs = utils.GetCurrentTimestampByMill()
	switch msg.ContentType {
	case constant.Text:
		fallthrough
	case constant.Picture:
		fallthrough
	case constant.Voice:
		fallthrough
	case constant.Video:
		fallthrough
	case constant.File:
		fallthrough
	case constant.AtText:
		fallthrough
	case constant.Merger:
		fallthrough
	case constant.Card:
		fallthrough
	case constant.Location:
		fallthrough
	case constant.Custom:
		fallthrough
	case constant.Quote: // 引用
		utils.SetSwitchFromOptions(msg.Options, constant.IsConversationUpdate, true)
		utils.SetSwitchFromOptions(msg.Options, constant.IsUnreadCount, true)
		utils.SetSwitchFromOptions(msg.Options, constant.IsSenderSync, true)
	case constant.Revoke: // 撤销
		utils.SetSwitchFromOptions(msg.Options, constant.IsUnreadCount, false)
		utils.SetSwitchFromOptions(msg.Options, constant.IsOfflinePush, false)
	case constant.HasReadReceipt: // 已读回执
		//log.Info("", "this is a test start", msg, msg.Options)
		utils.SetSwitchFromOptions(msg.Options, constant.IsConversationUpdate, false)
		utils.SetSwitchFromOptions(msg.Options, constant.IsUnreadCount, false)
		utils.SetSwitchFromOptions(msg.Options, constant.IsOfflinePush, false)
		//log.Info("", "this is a test end", msg, msg.Options)
	case constant.Typing: // 打字
		utils.SetSwitchFromOptions(msg.Options, constant.IsHistory, false)
		utils.SetSwitchFromOptions(msg.Options, constant.IsPersistent, false)
		utils.SetSwitchFromOptions(msg.Options, constant.IsSenderSync, false)
		utils.SetSwitchFromOptions(msg.Options, constant.IsConversationUpdate, false)
		utils.SetSwitchFromOptions(msg.Options, constant.IsUnreadCount, false)
		utils.SetSwitchFromOptions(msg.Options, constant.IsOfflinePush, false)
	}
}

func (rpc *chatRpcServer) SendMsg(_ context.Context, pb *pb_chat.SendMsgReq) (resp *pb_chat.SendMsgResp, err error) {
	var (
		msgToMQ   pb_chat.MsgDataToMQ
		isHistory bool
		req       MsgCallBackReq
		canSend   bool
		replay    pb_chat.SendMsgResp
		isSend    bool

		//群消息
		clientConn        *grpc.ClientConn
		client            pb_group.GroupClient
		groupAllMemberReq *pb_group.GetGroupAllMemberBasicReq
		reply             *pb_group.GetGroupAllMemberBasicResp
		memberUserIdList  []string
	)
	replay = pb_chat.SendMsgResp{}
	rpc.encapsulateMsgData(pb.MsgData)
	msgToMQ = pb_chat.MsgDataToMQ{
		Token:       pb.Token,
		OperationId: pb.OperationId,
	}
	// 是否是历史消息
	isHistory = utils.GetSwitchFromOptions(pb.MsgData.Options, constant.IsHistory)

	req = MsgCallBackReq{
		SendID:      pb.MsgData.SendId,
		RecvID:      pb.MsgData.RecvId,
		Content:     string(pb.MsgData.Content),
		SendTs:      pb.MsgData.SendTs,
		MsgFrom:     pb.MsgData.MsgFrom,
		ContentType: pb.MsgData.ContentType,
		SessionType: pb.MsgData.SessionType,
		PlatformID:  pb.MsgData.SenderPlatformId,
		MsgID:       pb.MsgData.ClientMsgId,
	}
	if isHistory == false {
		req.IsOnlineOnly = true
	}
	canSend, err = callbackWordFilter(pb)
	if err != nil {
		//return
	}
	// TODO: 调试强制发送
	func() {
		canSend = true
	}()
	if canSend == false {
		return returnMsg(&replay, pb, 201, "callbackWordFilter result stop rpc and return", "", 0)
	}

	switch pb.MsgData.SessionType {
	case constant.SingleChatType:
		canSend, err = callbackBeforeSendSingleMsg(pb)
		// TODO: 调试强制发送
		func() {
			canSend = true
		}()
		if err != nil {
			//return
		}
		if !canSend {
			return returnMsg(&replay, pb, 201, "callbackBeforeSendSingleMsg result stop rpc and return", "", 0)
		}
		isSend = modifyMessageByUserMessageReceiveOpt(pb.MsgData.RecvId, pb.MsgData.SendId, constant.SingleChatType, pb)
		// TODO: 调试强制发送
		func() {
			isSend = true
		}()
		if isSend {
			msgToMQ.MsgData = pb.MsgData
			// 消息只要成功落入MQ中，就可以视为发送成功，消息发送的可靠性依赖于MQ集群。
			err = rpc.sendMsgToKafka(&msgToMQ, msgToMQ.MsgData.RecvId)
			if err != nil {
				// TODO: kafka发送消息失败
				return returnMsg(&replay, pb, 201, "kafka send msg err", "", 0)
			}
		}
		if msgToMQ.MsgData.SendId != msgToMQ.MsgData.RecvId { //Filter messages sent to yourself
			// 消息只要成功落入MQ中，就可以视为发送成功，消息发送的可靠性依赖于MQ集群。
			err = rpc.sendMsgToKafka(&msgToMQ, msgToMQ.MsgData.SendId)
			if err != nil {
				// TODO: kafka发送消息失败
				return returnMsg(&replay, pb, 201, "kafka send msg err", "", 0)
			}
		}
		// callback
		if err = callbackAfterSendSingleMsg(pb); err != nil {
			// TODO:错误
		}
		return returnMsg(&replay, pb, 0, "", msgToMQ.MsgData.ServerMsgId, msgToMQ.MsgData.SendTs)
	case constant.GroupChatType:
		// callback
		canSend, err = callbackBeforeSendGroupMsg(pb)
		if err != nil {
			// TODO:ERROR
		}
		// TODO: 调试强制发送
		func() {
			canSend = true
		}()
		if canSend == false {
			return returnMsg(&replay, pb, 201, "callbackBeforeSendGroupMsg result stop rpc and return", "", 0)
		}
		clientConn = factory.ClientConn(config.Config.RPCRegisterName.GroupName)
		client = pb_group.NewGroupClient(clientConn)
		groupAllMemberReq = &pb_group.GetGroupAllMemberBasicReq{
			GroupId:     pb.MsgData.GroupId,
			OperationId: pb.OperationId,
		}
		reply, err = client.GetGroupAllMemberBasic(context.Background(), groupAllMemberReq)
		if err != nil {
			// TODO:ERROR
			return returnMsg(&replay, pb, 201, err.Error(), "", 0)
		}
		if reply.Common.Code != 0 {
			return returnMsg(&replay, pb, reply.Common.Code, reply.Common.Msg, "", 0)
		}
		memberUserIdList = func(members []*pb_group.GroupMemberInfo) (userIds []string) {
			userIds = make([]string, 0)
			for _, v := range members {
				userIds = append(userIds, v.UserId)
			}
			return
		}(reply.MemberList)

		//TODO:消息分类处理
		switch pb.MsgData.ContentType {
		case constant.MemberKickedNotification:
		case constant.MemberQuitNotification:
		case constant.AtText:
		}
		for _, userId := range memberUserIdList {
			pb.MsgData.RecvId = userId
			isSend = modifyMessageByUserMessageReceiveOpt(userId, pb.MsgData.GroupId, constant.GroupChatType, pb)
			// TODO: 调试强制发送
			func() {
				isSend = true
			}()
			if isSend == true {
				msgToMQ.MsgData = pb.MsgData
				err = rpc.sendMsgToKafka(&msgToMQ, userId)
				if err != nil {
					return returnMsg(&replay, pb, 201, "kafka send msg err", "", 0)
				}
			}
		}
	}
	return returnMsg(&replay, pb, 0, "", msgToMQ.MsgData.ServerMsgId, msgToMQ.MsgData.SendTs)
}

func returnMsg(replay *pb_chat.SendMsgResp, pb *pb_chat.SendMsgReq, errCode int32, errMsg, serverMsgID string, sendTime int64) (*pb_chat.SendMsgResp, error) {
	replay.ErrCode = errCode
	replay.ErrMsg = errMsg
	replay.ServerMsgId = serverMsgID
	replay.ClientMsgId = pb.MsgData.ClientMsgId
	replay.SendTime = sendTime
	return replay, nil
}

func modifyMessageByUserMessageReceiveOpt(userID, sourceID string, sessionType int, pb *pb_chat.SendMsgReq) bool {
	conversationID := utils.GetConversationIDBySessionType(sourceID, sessionType)
	opt, err := redis.GetSingleConversationMsgOpt(userID, conversationID)
	if err != nil && err != redis.ErrorRedisNil {
		return true
	}
	switch opt {
	case constant.ReceiveMessage:
		return true
	case constant.NotReceiveMessage:
		return false
	case constant.ReceiveNotNotifyMessage:
		if pb.MsgData.Options == nil {
			pb.MsgData.Options = make(map[string]bool, 10)
		}
		utils.SetSwitchFromOptions(pb.MsgData.Options, constant.IsOfflinePush, false)
		return true
	}
	return true
}

func (rpc *chatRpcServer) sendMsgToKafka(m *pb_chat.MsgDataToMQ, key string) (err error) {
	var (
		pid    int32
		offset int64
	)
	pid, offset, err = rpc.producer.SendMessage(m, key)
	if err != nil {
		fmt.Println("kafka send failed", m.OperationId, "send data", m.String(), "pid", pid, "offset", offset, "err", err.Error(), "key", key)
	}
	return
}
