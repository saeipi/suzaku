package rpc_group

import (
	"context"
	"suzaku/internal/domain/do"
	"suzaku/internal/domain/po_mysql"
	"suzaku/internal/domain/repo/repo_mysql"
	pb_group "suzaku/pkg/proto/group"
)

func (rpc *groupRpcServer) RequestJoinGroup(_ context.Context, req *pb_group.RequestJoinGroupReq) (resp *pb_group.RequestJoinGroupResp, _ error) {
	var (
		member      *po_mysql.GroupMember
		joinRequest *po_mysql.GroupRequest
		err         error
	)
	member, err = repo_mysql.GroupRepo.IsJoined(req.GroupId, req.UserId)
	if err != nil {
		//TODO:错误
		return
	}
	if member.UserId != "" {
		//TODO:已经加入
		return
	}
	joinRequest = &po_mysql.GroupRequest{
		UserId:       req.UserId,
		GroupId:      req.GroupId,
		HandleUserId: req.OperationId,
		ReqMsg:       req.ReqMsg,
		ReqSource:    req.ReqSource,
	}
	err = repo_mysql.GroupRepo.RequestJoin(joinRequest)
	if err != nil {
		//TODO:错误
		return
	}
	return
}

func (rpc *groupRpcServer) HandleRequestJoinGroup(_ context.Context, req *pb_group.HandleRequestJoinGroupReq) (resp *pb_group.HandleRequestJoinGroupResp, _ error) {
	var (
		result *do.JoinGroupResult
		err    error
	)
	result, err = repo_mysql.GroupRepo.HandleRequestJoin(req)
	if err != nil {
		//TODO:错误
		return
	}
	switch result.HandleResult {
	case int32(pb_group.HANDLE_JOIN_GROUP_RESULT_APPROVE), int32(pb_group.HANDLE_JOIN_GROUP_RESULT_RESOLUTELY):
		//TODO:通知
	}
	return
}
