package msg_handler

import (
	"bytes"
	"context"
	"encoding/gob"
	"github.com/go-playground/validator/v10"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"suzaku/internal/msg_gateway/protocol"
	ws "suzaku/internal/msg_gateway/ws_server"
	"suzaku/internal/rpc/grpc_client"
	"suzaku/pkg/common/config"
	"suzaku/pkg/constant"
	pb_chat "suzaku/pkg/proto/chart"
	"suzaku/pkg/proto/pb_ws"
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
		buffer  *bytes.Buffer
		decoder *gob.Decoder
		req     protocol.MessageReq
		err     error
	)
	req = protocol.MessageReq{}
	buffer = bytes.NewBuffer(msg.Body)
	decoder = gob.NewDecoder(buffer)
	err = decoder.Decode(&req)
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
	default:
	}
}

func (h *MsgHandler) getNewestSeq(client *ws.Client, req *protocol.MessageReq) {
	var (
		rpcReq     pb_chat.GetMaxAndMinSeqReq
		reply      *pb_chat.GetMaxAndMinSeqResp
		clientConn *grpc.ClientConn
		chatClient pb_chat.ChatClient
		rpcReply   *pb_chat.GetMaxAndMinSeqResp
		err        error
	)
	rpcReq = pb_chat.GetMaxAndMinSeqReq{}
	reply = new(pb_chat.GetMaxAndMinSeqResp)
	rpcReq.UserId = req.SendID
	rpcReq.OperationId = req.OperationID

	clientConn = grpc_client.ClientConn(config.Config.RPCRegisterName.OfflineMessageName)
	if clientConn == nil {
		//TODO: error
		return
	}
	chatClient = pb_chat.NewChatClient(clientConn)
	rpcReply, err = chatClient.GetMaxAndMinSeq(context.Background(), &rpcReq)
	if err == nil {
		//TODO: error
		reply.ErrCode = 500
		reply.ErrMsg = ""
		h.getSeqResp(client, req, reply)
	} else {
		h.getSeqResp(client, req, rpcReply)
	}
}

func (h *MsgHandler) getSeqResp(client *ws.Client, req *protocol.MessageReq, pb *pb_chat.GetMaxAndMinSeqResp) {
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
		buffer bytes.Buffer
		encode *gob.Encoder
		err    error
	)
	encode = gob.NewEncoder(&buffer)
	err = encode.Encode(data)
	if err != nil {
		//TODO :错误
		return
	}
	client.Send(buffer.Bytes())
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
		clientConn = grpc_client.ClientConn(config.Config.RPCRegisterName.OfflineMessageName)
		chatClient = pb_chat.NewChatClient(clientConn)
		reply, err = chatClient.SendMsg(context.Background(), &reqReq)
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
	var mReplyData pb_ws.UserSendMsgResp
	mReplyData.ClientMsgId = reply.GetClientMsgId()
	mReplyData.ServerMsgId = reply.GetServerMsgId()
	mReplyData.SendTime = reply.GetSendTime()
	b, _ := proto.Marshal(&mReplyData)
	mReply := protocol.MessageResp{
		ReqIdentifier: req.ReqIdentifier,
		MsgIncr:       req.MsgIncr,
		ErrCode:       reply.GetErrCode(),
		ErrMsg:        reply.GetErrMsg(),
		OperationID:   req.OperationID,
		Data:          b,
	}
	h.sendMessage(client, mReply)
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
