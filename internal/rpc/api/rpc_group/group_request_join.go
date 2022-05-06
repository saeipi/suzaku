package rpc_group

import (
	"context"
	pb_group "suzaku/pkg/proto/group"
)

func (rpc *groupRpcServer) RequestJoinGroup(_ context.Context, req *pb_group.RequestJoinGroupReq) (resp *pb_group.RequestJoinGroupResp, _ error) {
	var (
		err error
	)
	err = err
	return
}
