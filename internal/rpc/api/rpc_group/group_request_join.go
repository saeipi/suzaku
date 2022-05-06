package rpc_group

import (
	"context"
	"suzaku/internal/domain/po_mysql"
	"suzaku/internal/domain/repo/repo_mysql"
	pb_group "suzaku/pkg/proto/group"
	"time"
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
		ReqMsg:       req.ReqMessage,
		ReqTs:        time.Now().Unix(),
	}

	err = repo_mysql.GroupRepo.RequestJoin(joinRequest)
	if err != nil {
		//TODO:错误
		return
	}
	return
}
