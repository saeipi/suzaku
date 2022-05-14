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
		messages   []*po_mysql.Message
		msgs       []*po_mongo.Message
		appendList []*pb_ws.MsgData
		lastMsg    *pb_ws.MsgData
		nowMill    int64
		err        error
	)
	resp = &pb_chat.GetHistoryMessagesResp{
		Common:  &pb_com.CommonResp{},
		MsgList: make([]*pb_ws.MsgData, 0),
	}
	nowMill = utils.GetCurrentTimestampByMill()
	if req.CreatedTs == 0 || nowMill-req.CreatedTs < MongoTimeFrame {
		msgs, err = repo_mongo.MgChatRepo.HistoryMessages(req)
		if err != nil {
			//TODO:Error
			resp.Common.Code = 777
			return
		}
		copier.Copy(&resp.MsgList, msgs)
		if len(msgs) == int(req.PageSize) {
			return
		}
	}

	if len(resp.MsgList) > 0 {
		if req.Back {
			lastMsg = resp.MsgList[len(resp.MsgList)-1]
		} else {
			lastMsg = resp.MsgList[0]
		}
		req.PageSize = int32(int(req.PageSize) - len(resp.MsgList))
		req.Seq = int64(lastMsg.Seq)
		req.CreatedTs = lastMsg.CreatedTs
	}

	messages, err = repo_mysql.ChatRepo.HistoryMessages(req)
	if err != nil {
		//TODO:Error
		resp.Common.Code = 777
		return
	}

	if len(resp.MsgList) > 0 {
		appendList = make([]*pb_ws.MsgData, 0)
		copier.Copy(&appendList, messages)
		if req.Back {
			resp.MsgList = append(resp.MsgList, appendList...)
		} else {
			resp.MsgList = append(appendList, resp.MsgList...)
		}
	} else {
		copier.Copy(&resp.MsgList, messages)
	}
	return
}
