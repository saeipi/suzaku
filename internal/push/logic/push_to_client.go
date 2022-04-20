package logic

import (
	"context"
	"google.golang.org/grpc"
	"strings"
	"suzaku/pkg/common/config"
	"suzaku/pkg/common/grpc-etcdv3/getcdv3"
	"suzaku/pkg/constant"
	pb_push "suzaku/pkg/proto/push"
	pb_relay "suzaku/pkg/proto/relay"
	"suzaku/pkg/utils"
)

func MsgToUser(pushMsg *pb_push.PushMsgReq) {
	var (
		wsResult      []*pb_relay.SingleMsgToUser
		isOfflinePush bool
		grpcCons      []*grpc.ClientConn
		conn          *grpc.ClientConn
		msgClient     pb_relay.OnlineMessageRelayServiceClient
		req           *pb_relay.OnlinePushMsgReq
		reply         *pb_relay.OnlinePushMsgResp
		result        *pb_relay.SingleMsgToUser
		platformID    int32
		err           error
	)
	wsResult = make([]*pb_relay.SingleMsgToUser, 0)
	isOfflinePush = utils.GetSwitchFromOptions(pushMsg.MsgData.Options, constant.IsOfflinePush)
	grpcCons = getcdv3.GetConn4Unique(config.Config.Etcd.Schema,
		strings.Join(config.Config.Etcd.Address, ","),
		config.Config.RPCRegisterName.OnlineMessageRelayName)
	for _, conn = range grpcCons {
		req = &pb_relay.OnlinePushMsgReq{
			OperationId:  pushMsg.OperationId,
			MsgData:      pushMsg.MsgData,
			PushToUserId: pushMsg.PushToUserId,
		}
		msgClient = pb_relay.NewOnlineMessageRelayServiceClient(conn)
		reply, err = msgClient.OnlinePushMsg(context.Background(), req)
		if err != nil {
			continue
		}
		if reply != nil && reply.Resp != nil {
			wsResult = append(wsResult, reply.Resp...)
		}
	}
	if isOfflinePush && pushMsg.PushToUserId != pushMsg.MsgData.SendId {
		for _, result = range wsResult {
			if result.ResultCode == 0 {
				continue
			}
			for _, platformID = range pushTerminal {
				if result.RecvPlatFormId == platformID {
					//TODO:
				}
			}
		}
	}
}
