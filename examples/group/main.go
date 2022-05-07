package main

import (
	"context"
	"suzaku/internal/rpc/api/rpc_group"
	pb_group "suzaku/pkg/proto/group"
)

func main() {
	server := rpc_group.NewGroupRpcServer(1001)
	req := &pb_group.GetGroupAllMemberBasicReq{GroupId: "123123"}
	server.GetGroupAllMemberBasic(context.Background(), req)
}
