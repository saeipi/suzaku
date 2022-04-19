package rpc_server

import (
	"bytes"
	"context"
	"encoding/gob"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"suzaku/internal/msg_gateway/protocol"
	"suzaku/internal/msg_gateway/ws_server"
	"suzaku/internal/rpc/rpc_category"
	"suzaku/pkg/common/config"
	"suzaku/pkg/constant"
	pb_relay "suzaku/pkg/proto/relay"
)

type RPCServer struct {
	pb_relay.UnimplementedOnlineMessageRelayServiceServer
	rpc_category.Rpc
	wsSvr *ws_server.WServer
}

func NewRPCServer(port int, wsSvr *ws_server.WServer) *RPCServer {
	return &RPCServer{
		Rpc:   rpc_category.NewRpcServer(port, config.Config.RPCRegisterName.OnlineMessageRelayName),
		wsSvr: wsSvr,
	}
}

func (rpc *RPCServer) Run() {
	var (
		server *grpc.Server
	)
	server = grpc.NewServer()
	pb_relay.RegisterOnlineMessageRelayServiceServer(server, rpc)
	rpc.Rpc.RunServer(server)
}

func (rpc *RPCServer) OnlinePushMsg(ctx context.Context, req *pb_relay.OnlinePushMsgReq) (resp *pb_relay.OnlinePushMsgResp, err error) {
	var (
		msgResp  protocol.MessageResp
		msgBytes []byte

		replyBytes bytes.Buffer
		encoder    *gob.Encoder
		sendResult *pb_relay.SingleMsgToUser
		platformID int32
		ok         bool
	)

	msgBytes, err = proto.Marshal(req.MsgData)
	if err != nil {
		//TODO:错误处理
		return
	}
	msgResp = protocol.MessageResp{
		ReqIdentifier: constant.WSPushMsg,
		OperationID:   req.OperationId,
		Data:          msgBytes,
	}

	encoder = gob.NewEncoder(&replyBytes)
	err = encoder.Encode(msgResp)
	if err != nil {
		//TODO:错误处理
		return
	}
	msgBytes = replyBytes.Bytes()

	resp = &pb_relay.OnlinePushMsgResp{Resp: make([]*pb_relay.SingleMsgToUser, 0)}
	ok = rpc.wsSvr.Send(req.PushToUserId, msgBytes)
	if ok == false {
		// 离线
		sendResult = &pb_relay.SingleMsgToUser{
			ResultCode:     constant.WSStatusOffline,
			RecvId:         req.PushToUserId,
			RecvPlatFormId: platformID,
		}
		resp.Resp = append(resp.Resp, sendResult)
	}

	// 在线
	sendResult = &pb_relay.SingleMsgToUser{
		ResultCode:     constant.WSStatusOnline,
		RecvId:         req.PushToUserId,
		RecvPlatFormId: platformID,
	}
	resp.Resp = append(resp.Resp, sendResult)
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
			ps.Status = constant.WSStatusOnline
			sr.Status = constant.WSStatusOnline
			sr.DetailPlatformStatus = append(sr.DetailPlatformStatus, ps)
		}

		if sr.Status == constant.WSStatusOnline {
			resp.SuccessResult = append(resp.SuccessResult, sr)
		}
	}
	return
}
