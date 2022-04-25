package rpc_server

import (
	"context"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"suzaku/internal/msg_gateway/protocol"
	"suzaku/internal/msg_gateway/ws_server"
	"suzaku/internal/rpc/rpc_category"
	"suzaku/pkg/common/config"
	"suzaku/pkg/constant"
	"suzaku/pkg/factory"
	pb_relay "suzaku/pkg/proto/relay"
	"suzaku/pkg/utils"
)

type RPCServer struct {
	pb_relay.UnimplementedOnlineMessageRelayServiceServer
	rpc_category.Rpc
	wsSvr       *ws_server.WServer
	platformIds []int32
}

func NewRPCServer(port int, wsSvr *ws_server.WServer) *RPCServer {
	return &RPCServer{
		Rpc:         rpc_category.NewRpcServer(port, config.Config.RPCRegisterName.OnlineMessageRelayName),
		wsSvr:       wsSvr,
		platformIds: getPlatformIds(),
	}
}

func (rpc *RPCServer) Run() {
	var (
		server *grpc.Server
	)
	server = factory.NewRPCNewServer()
	pb_relay.RegisterOnlineMessageRelayServiceServer(server, rpc)
	rpc.Rpc.RunServer(server)
}

func (rpc *RPCServer) OnlinePushMsg(ctx context.Context, req *pb_relay.OnlinePushMsgReq) (resp *pb_relay.OnlinePushMsgResp, err error) {
	var (
		msgResp protocol.MessageResp
		msgBuf  []byte
		respBuf []byte

		sendResult *pb_relay.SingleMsgToUser
		platformID int32
		resultCode int
	)
	resp = &pb_relay.OnlinePushMsgResp{Resp: make([]*pb_relay.SingleMsgToUser, 0)}

	msgBuf, err = proto.Marshal(req.MsgData)
	if err != nil {
		//TODO:错误处理
		return
	}
	msgResp = protocol.MessageResp{
		ReqIdentifier: constant.WSPushMsg,
		OperationID:   req.OperationId,
		Data:          msgBuf,
	}
	respBuf, err = utils.ObjEncode(msgResp)
	if err != nil {
		//TODO:错误处理
		return
	}
	// TODO:发送给全平台目标用户
	for _, platformID = range rpc.platformIds {
		resultCode, _ = rpc.wsSvr.SendMessage(req.PushToUserId, platformID, respBuf)
		// 发送成功
		sendResult = &pb_relay.SingleMsgToUser{
			ResultCode:     int64(resultCode), // 成功标识 0
			RecvId:         req.PushToUserId,
			RecvPlatFormId: platformID,
		}
		resp.Resp = append(resp.Resp, sendResult)
	}
	return
}

func (rpc *RPCServer) GetUsersOnlineStatus(ctx context.Context, req *pb_relay.UsersOnlineStatusReq) (resp *pb_relay.UsersOnlineStatusResp, err error) {
	var (
		platformID int32
		userID     string
		ps         *pb_relay.SuccessDetail
		sr         *pb_relay.SuccessResult
	)
	resp = new(pb_relay.UsersOnlineStatusResp)

	for _, userID = range req.UserIdList {
		sr = new(pb_relay.SuccessResult)
		if rpc.wsSvr.IsOnline(userID) == true {
			ps = new(pb_relay.SuccessDetail)
			ps.PlatformId = platformID
			ps.Status = constant.OnlineStatus
			sr.Status = constant.OnlineStatus
			sr.DetailPlatformStatus = append(sr.DetailPlatformStatus, ps)
		}

		if sr.Status == constant.OnlineStatus {
			resp.SuccessResult = append(resp.SuccessResult, sr)
		}
	}
	return
}

func getPlatformIds() (ids []int32) {
	var (
		platformID int
	)
	ids = make([]int32, 0)
	for platformID = 1; platformID <= constant.LinuxPlatformID; platformID++ {
		ids = append(ids, int32(platformID))
	}
	return
}
