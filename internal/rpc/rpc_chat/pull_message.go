package rpc_chat

import (
	"context"
	"suzaku/pkg/common/redis"
	pb_chat "suzaku/pkg/proto/chart"
	"suzaku/pkg/proto/pb_ws"
)

func (rpc *chatRpcServer) GetMinMaxSeq(_ context.Context, req *pb_chat.GetMinMaxSeqReq) (resp *pb_chat.GetMinMaxSeqResp, _ error) {
	var (
		maxSeq uint64
		minSeq uint64
		err1   error
		err2   error
	)
	maxSeq, err1 = redis.GetUserMaxSeq(req.UserId)
	minSeq, err2 = redis.GetUserMinSeq(req.UserId)
	resp = new(pb_chat.GetMinMaxSeqResp)
	if err1 != nil {
		resp.MaxSeq = uint32(maxSeq)
	} else if err1 != redis.ErrorRedisNil {
		resp.MaxSeq = 0
	} else {
		//TODO:error
		resp.Code = 200
		resp.Msg = "redis get err"
	}
	if err2 == nil {
		resp.MinSeq = uint32(minSeq)
	} else if err2 == redis.ErrorRedisNil {
		resp.MinSeq = 0
	} else {
		//TODO:error
		resp.Code = 201
		resp.Msg = "redis get err"
	}
	return
}

func (rpc *chatRpcServer) PullMessageBySeqList(_ context.Context, req *pb_ws.PullMessageBySeqListReq) (resp *pb_ws.PullMessageBySeqListResp, err error) {
	var (
		seqMsg []*pb_ws.MsgData
	)
	resp = new(pb_ws.PullMessageBySeqListResp)
	seqMsg = make([]*pb_ws.MsgData, 0)
	//seqMsg, err = repo_mongo.MgChatRepo.GetMsgBySeqListMongo2(req.UserId, req.SeqList, req.OperationId)
	//if err != nil {
	//	//TODO:error
	//	resp.ErrCode = 201
	//	resp.ErrMsg = err.Error()
	//	return
	//}
	resp.Code = 0
	resp.Msg = ""
	resp.List = seqMsg
	return
}

type MsgFormats []*pb_ws.MsgData

// Implement the sort.Interface interface to get the number of elements method
func (s MsgFormats) Len() int {
	return len(s)
}

//Implement the sort.Interface interface comparison element method
func (s MsgFormats) Less(i, j int) bool {
	return s[i].SendTs < s[j].SendTs
}

//Implement the sort.Interface interface exchange element method
func (s MsgFormats) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
