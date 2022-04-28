package msg_handler

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"suzaku/internal/msg_gateway/protocol"
	ws "suzaku/internal/msg_gateway/ws_server"
	"suzaku/pkg/common/config"
	"suzaku/pkg/constant"
	"suzaku/pkg/factory"
	pb_chat "suzaku/pkg/proto/chart"
	"suzaku/pkg/proto/pb_ws"
	"suzaku/pkg/utils"
)

type MsgHandler struct {
	validate     *validator.Validate
	sendMsgCount uint64
}

func NewMsgHandler() *MsgHandler {
	return &MsgHandler{validate: validator.New()}
}

func (h *MsgHandler) MessageCallback(msg *ws.Message) {
	var (
		req protocol.MessageReq
		err error
	)
	req = protocol.MessageReq{}
	err = utils.BufferDecode(msg.Body, &req)
	if err != nil {
		// TODO :错误信息
		h.sendErrMsg(msg.Client, 200, err.Error(), 3001, "", "")
		msg.Client.Close()
		return
	}
	if err = h.validate.Struct(req); err != nil {
		// TODO :错误信息
		h.sendErrMsg(msg.Client, 201, err.Error(), req.ReqIdentifier, req.MsgIncr, req.OperationID)
		return
	}
	switch req.ReqIdentifier {
	case constant.WSGetNewestSeq:
		h.getNewestSeq(msg.Client, &req)
	case constant.WSSendMsg:
		h.sendMsgReq(msg.Client, &req)
	case constant.WSPullMsgBySeqList:
		h.pullMsgBySeqListReq(msg.Client, &req)
	default:
	}
}

func (h *MsgHandler) getNewestSeq(client *ws.Client, req *protocol.MessageReq) {
	var (
		rpcReq     pb_chat.GetMinMaxSeqReq
		reply      *pb_chat.GetMinMaxSeqResp
		clientConn *grpc.ClientConn
		chatClient pb_chat.ChatClient
		rpcReply   *pb_chat.GetMinMaxSeqResp
		err        error
	)
	rpcReq = pb_chat.GetMinMaxSeqReq{}
	reply = new(pb_chat.GetMinMaxSeqResp)
	rpcReq.UserId = req.SendID
	rpcReq.OperationId = req.OperationID

	clientConn = factory.ClientConn(config.Config.RPCRegisterName.OfflineMessageName)
	if clientConn == nil {
		//TODO: error
		return
	}
	chatClient = pb_chat.NewChatClient(clientConn)
	rpcReply, err = chatClient.GetMinMaxSeq(context.Background(), &rpcReq)
	if err == nil {
		//TODO: error
		reply.ErrCode = 500
		reply.ErrMsg = ""
		h.getSeqResp(client, req, reply)
	} else {
		h.getSeqResp(client, req, rpcReply)
	}
}

func (h *MsgHandler) getSeqResp(client *ws.Client, req *protocol.MessageReq, pb *pb_chat.GetMinMaxSeqResp) {
	var (
		replyData pb_ws.GetMaxAndMinSeqResp
		buffer    []byte
		reply     *protocol.MessageResp
		err       error
	)
	replyData.MaxSeq = pb.GetMaxSeq()
	replyData.MinSeq = pb.GetMinSeq()
	buffer, err = proto.Marshal(&replyData)
	if err != nil {
		//TODO:错误
		return
	}
	reply = &protocol.MessageResp{
		ReqIdentifier: req.ReqIdentifier,
		MsgIncr:       req.MsgIncr,
		ErrCode:       pb.GetErrCode(),
		ErrMsg:        pb.GetErrMsg(),
		OperationID:   req.OperationID,
		Data:          buffer,
	}
	h.sendMessage(client, reply)
}

func (h *MsgHandler) sendMessage(client *ws.Client, data interface{}) {
	var (
		buf []byte
		err error
	)
	buf, err = utils.ObjEncode(data)
	if err != nil {
		//TODO :错误
		return
	}
	client.Send(buf)
}

func (h *MsgHandler) sendErrMsg(client *ws.Client, errCode int32, errMsg string, reqIdentifier int32, msgIncr string, operationID string) {
	var (
		reply *protocol.MessageResp
	)
	reply = &protocol.MessageResp{
		ReqIdentifier: reqIdentifier,
		MsgIncr:       msgIncr,
		ErrCode:       errCode,
		ErrMsg:        errMsg,
		OperationID:   operationID,
	}
	h.sendMessage(client, reply)
}

func (h *MsgHandler) sendMsgReq(client *ws.Client, req *protocol.MessageReq) {
	var (
		reply      *pb_chat.SendMsgResp
		isPass     bool
		errCode    int32
		errMsg     string
		data       interface{}
		msgData    pb_ws.MsgData
		reqReq     pb_chat.SendMsgReq
		clientConn *grpc.ClientConn
		chatClient pb_chat.ChatClient
		err        error
	)

	h.sendMsgCount++
	reply = new(pb_chat.SendMsgResp)
	isPass, errCode, errMsg, data = h.argsValidate(req, constant.WSSendMsg)
	if isPass {
		msgData = data.(pb_ws.MsgData)
		reqReq = pb_chat.SendMsgReq{
			Token:       req.Token,
			OperationId: req.OperationID,
			MsgData:     &msgData,
		}
		clientConn = factory.ClientConn(config.Config.RPCRegisterName.OfflineMessageName)
		chatClient = pb_chat.NewChatClient(clientConn)
		reply, err = chatClient.SendMsg(context.Background(), &reqReq)
		if reply == nil {
			//TODO: error
			return
		}
		if err != nil {
			reply.ErrCode = 200
			reply.ErrMsg = err.Error()
			h.sendMsgResp(client, req, reply)
		} else {
			h.sendMsgResp(client, req, reply)
		}
		return
	}
	reply.ErrCode = errCode
	reply.ErrMsg = errMsg
	h.sendMsgResp(client, req, reply)
}

func (h *MsgHandler) sendMsgResp(client *ws.Client, req *protocol.MessageReq, reply *pb_chat.SendMsgResp) {
	var (
		replyData pb_ws.UserSendMsgResp
		buf       []byte
		resp      protocol.MessageResp
		err       error
	)
	replyData.ClientMsgId = reply.GetClientMsgId()
	replyData.ServerMsgId = reply.GetServerMsgId()
	replyData.SendTime = reply.GetSendTime()
	buf, err = proto.Marshal(&replyData)
	if err != nil {
		//TODO: error
		return
	}
	resp = protocol.MessageResp{
		ReqIdentifier: req.ReqIdentifier,
		MsgIncr:       req.MsgIncr,
		ErrCode:       reply.GetErrCode(),
		ErrMsg:        reply.GetErrMsg(),
		OperationID:   req.OperationID,
		Data:          buf,
	}
	h.sendMessage(client, resp)
}

func (h *MsgHandler) pullMsgBySeqListReq(client *ws.Client, req *protocol.MessageReq) {
	var (
		reply      *pb_ws.PullMessageBySeqListResp
		isPass     bool
		errCode    int32
		errMsg     string
		data       interface{}
		rpcReq     pb_ws.PullMessageBySeqListReq
		clientConn *grpc.ClientConn
		chatClient pb_chat.ChatClient
		err        error
	)
	reply = new(pb_ws.PullMessageBySeqListResp)
	isPass, errCode, errMsg, data = h.argsValidate(req, constant.WSPullMsgBySeqList)
	if isPass {
		rpcReq = pb_ws.PullMessageBySeqListReq{}
		rpcReq.SeqList = data.(pb_ws.PullMessageBySeqListReq).SeqList
		rpcReq.UserId = req.SendID
		rpcReq.OperationId = req.OperationID

		clientConn = factory.ClientConn(config.Config.RPCRegisterName.OfflineMessageName)
		chatClient = pb_chat.NewChatClient(clientConn)
		reply, err = chatClient.PullMessageBySeqList(context.Background(), &rpcReq)
		if reply == nil {
			//TODO: error
			return
		}
		if err != nil {
			reply.ErrCode = 200
			reply.ErrMsg = err.Error()
			h.pullMsgBySeqListResp(client, req, reply)
		} else {
			h.pullMsgBySeqListResp(client, req, reply)
		}
		return
	}
	reply.ErrCode = errCode
	reply.ErrMsg = errMsg
	h.pullMsgBySeqListResp(client, req, reply)
}

func (h *MsgHandler) pullMsgBySeqListResp(client *ws.Client, req *protocol.MessageReq, reply *pb_ws.PullMessageBySeqListResp) {
	var (
		buf  []byte
		resp protocol.MessageResp
		err  error
	)

	buf, err = proto.Marshal(reply)
	if err != nil {
		//TODO: error
		return
	}
	resp = protocol.MessageResp{
		ReqIdentifier: req.ReqIdentifier,
		MsgIncr:       req.MsgIncr,
		ErrCode:       reply.GetErrCode(),
		ErrMsg:        reply.GetErrMsg(),
		OperationID:   req.OperationID,
		Data:          buf,
	}
	h.sendMessage(client, resp)
}

func (h *MsgHandler) argsValidate(m *protocol.MessageReq, r int32) (isPass bool, errCode int32, errMsg string, returnData interface{}) {
	switch r {
	case constant.WSSendMsg:
		data := pb_ws.MsgData{}
		if err := proto.Unmarshal(m.Data, &data); err != nil {
			//log.ErrorByKv("Decode Data struct  err", "", "err", err.Error(), "reqIdentifier", r)
			return false, 203, err.Error(), nil
		}
		if err := h.validate.Struct(data); err != nil {
			//log.ErrorByKv("data args validate  err", "", "err", err.Error(), "reqIdentifier", r)
			return false, 204, err.Error(), nil
		}
		return true, 0, "", data
	case constant.WSPullMsgBySeqList:
		data := pb_ws.PullMessageBySeqListReq{}
		if err := proto.Unmarshal(m.Data, &data); err != nil {
			//log.ErrorByKv("Decode Data struct  err", "", "err", err.Error(), "reqIdentifier", r)
			return false, 203, err.Error(), nil
		}
		if err := h.validate.Struct(data); err != nil {
			//log.ErrorByKv("data args validate  err", "", "err", err.Error(), "reqIdentifier", r)
			return false, 204, err.Error(), nil
		}
		return true, 0, "", data
	default:
	}
	return false, 204, "args err", nil
}
