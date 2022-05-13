package api_chat

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc"
	"suzaku/internal/interface/dto/dto_api"
	"suzaku/pkg/common/config"
	"suzaku/pkg/constant"
	"suzaku/pkg/factory"
	"suzaku/pkg/http"
	pb_chat "suzaku/pkg/proto/chart"
	"suzaku/pkg/proto/pb_ws"
	"suzaku/pkg/utils"
)

/*
在线字符或文本转二进制工具: https://tooltt.com/txt2bin/
在线字符串转数组(Array):http://tools.bugscaner.com/text/string2array.html
*/

func newUserSendMsg(token string, params *dto_api.SendMsgReq) *pb_chat.SendMsgReq {
	pbData := pb_chat.SendMsgReq{
		Token:       token,
		OperationId: params.OperationID,
		MsgData: &pb_ws.MsgData{
			SendId:           params.SendID,
			RecvId:           params.Data.RecvID,
			GroupId:          params.Data.GroupID,
			ClientMsgId:      params.Data.ClientMsgID,
			SenderPlatformId: params.SenderPlatformID,
			SenderNickname:   params.SenderNickName,
			SenderAvatarUrl:  params.SenderAvatarUrl,
			SessionType:      params.Data.SessionType,
			SessionId:        params.Data.SessionId,
			MsgFrom:          params.Data.MsgFrom,
			ContentType:      params.Data.ContentType,
			Content:          params.Data.Content,
			CreatedTs:        params.Data.CreatedTs,
			Options:          params.Data.Options,
			OfflinePushInfo:  params.Data.OffLineInfo,
		},
	}
	return &pbData
}

func SendMessage(c *gin.Context) {
	var (
		params     *dto_api.SendMsgReq
		token      string
		sendMsg    *pb_chat.SendMsgReq
		clientConn *grpc.ClientConn
		client     pb_chat.ChatClient
		reply      *pb_chat.SendMsgResp
		resp       *dto_api.SendMsgResp
		err        error
	)

	if err = c.BindJSON(&params); err != nil {
		http.Error(c, err, http.ErrorCodeHttpReqDeserializeFailed)
		return
	}
	token = c.Request.Header.Get(constant.HttpKeyToken)
	sendMsg = newUserSendMsg(token, params)

	clientConn = factory.ClientConn(config.Config.RPCRegisterName.OfflineMessageName)
	client = pb_chat.NewChatClient(clientConn)

	reply, err = client.SendMsg(context.Background(), sendMsg)
	if reply != nil && reply.Code > 0 {
		http.Err(c, reply.Msg, reply.Code)
		return
	}

	resp = &dto_api.SendMsgResp{}
	utils.CopyStructFields(resp, reply)
	http.Success(c, resp)
}

func HistoryMessages(c *gin.Context) {
	var (
		params     dto_api.HistoryMessagesReq
		req        *pb_chat.GetHistoryMessagesReq
		clientConn *grpc.ClientConn
		client     pb_chat.ChatClient
		reply      *pb_chat.GetHistoryMessagesResp
		resp       *dto_api.HistoryMessagesResp
		err        error
	)
	if err = c.ShouldBindQuery(&params); err != nil {
		http.Error(c, err, http.ErrorCodeHttpReqDeserializeFailed)
		return
	}
	req = new(pb_chat.GetHistoryMessagesReq)
	copier.Copy(req, params)

	clientConn = factory.ClientConn(config.Config.RPCRegisterName.OfflineMessageName)
	client = pb_chat.NewChatClient(clientConn)
	reply, _ = client.GetHistoryMessages(context.Background(), req)

	if reply == nil {
		http.Error(c, http.ErrorHttpServiceFailure, http.ErrorCodeHttpServiceFailure)
		return
	}
	if reply.Common.Code > 0 {
		http.Err(c, reply.Common.Msg, reply.Common.Code)
		return
	}
	resp = &dto_api.HistoryMessagesResp{}
	copier.Copy(resp, reply)
	http.Success(c, resp)
}
