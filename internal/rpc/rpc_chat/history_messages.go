package rpc_chat

import (
	"context"
	"github.com/jinzhu/copier"
	"suzaku/internal/domain/po_mongo"
	"suzaku/internal/domain/po_mysql"
	"suzaku/internal/domain/repo/repo_mongo"
	"suzaku/internal/domain/repo/repo_mysql"
	pb_chat "suzaku/pkg/proto/chart"
	"suzaku/pkg/proto/pb_com"
	"suzaku/pkg/proto/pb_ws"
	"suzaku/pkg/utils"
)

const (
	MongoTimeFrame = 1296000000 // 15天
)

func (rpc *chatRpcServer) GetHistoryMessages(_ context.Context, req *pb_chat.GetHistoryMessagesReq) (resp *pb_chat.GetHistoryMessagesResp, _ error) {
	var (
		messages []*po_mysql.Message
		msgs     []*po_mongo.Message
		err      error
	)
	resp = &pb_chat.GetHistoryMessagesResp{
		Common:  &pb_com.CommonResp{},
		MsgList: make([]*pb_ws.MsgData, 0),
	}
	if req.CreatedTs == 0 || utils.GetCurrentTimestampByMill()-req.CreatedTs < MongoTimeFrame {
		msgs, err = repo_mongo.MgChatRepo.HistoryMessages(req)
		if err != nil {
			//TODO:Error
			resp.Common.Code = 777
			return
		}
		copier.Copy(&resp.MsgList, msgs)
		// 查mongodb
		return
	}
	messages, err = repo_mysql.ChatRepo.HistoryMessages(req)
	if err != nil {
		//TODO:Error
		resp.Common.Code = 777
		return
	}
	copier.Copy(&resp.MsgList, messages)
	return
}
