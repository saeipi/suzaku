package rpc_chat

import (
	"context"
	"google.golang.org/grpc"
	"suzaku/internal/domain/do"
	"suzaku/pkg/common/config"
	"suzaku/pkg/constant"
	"suzaku/pkg/factory"
	pb_chat "suzaku/pkg/proto/chart"
	pb_ws "suzaku/pkg/proto/pb_ws"
	"suzaku/pkg/utils"
)

func Notification(n *do.NotificationMsg) {
	var (
		req             pb_chat.SendMsgReq
		msg             pb_ws.MsgData
		offlineInfo     pb_ws.OfflinePushInfo
		title, desc, ex string
		//pushSwitch, unReadCount bool
		//reliabilityLevel int

		clientConn *grpc.ClientConn
		client     pb_chat.ChatClient
		reply      *pb_chat.SendMsgResp
	)
	req.OperationId = n.OperationID
	msg.SendId = n.SendID
	msg.RecvId = n.RecvID
	msg.Content = n.Content
	msg.MsgFrom = n.MsgFrom
	msg.ContentType = n.ContentType
	msg.SessionType = n.SessionType
	msg.CreatedTs = utils.GetCurrentTimestampByMill() // 微妙
	msg.ClientMsgId = utils.GetMsgID(n.SendID)
	msg.Options = make(map[string]bool, 7)
	switch n.SessionType {
	case constant.GroupChatType:
		msg.RecvId = ""
		msg.GroupId = n.RecvID
	}
	offlineInfo.IosBadgeCount = config.Config.IosPush.BadgeCount
	offlineInfo.IosPushSound = config.Config.IosPush.PushSound

	offlineInfo.Title = title
	offlineInfo.Desc = desc
	offlineInfo.Ex = ex
	msg.OfflinePushInfo = &offlineInfo
	req.MsgData = &msg

	clientConn = factory.ClientConn(config.Config.RPCRegisterName.OfflineMessageName)
	client = pb_chat.NewChatClient(clientConn)

	reply, _ = client.SendMsg(context.Background(), &req)
	if reply == nil {
		//TODO:错误处理
		return
	}
	if reply.Code != 0 {
		//TODO:错误处理
		return
	}
}
