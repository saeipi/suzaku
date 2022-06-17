package main

import (
	"context"
	"suzaku/internal/rpc/rpc_group"
	pb_group "suzaku/pkg/proto/group"
)

func main() {
	groupId := "666888"
	userId := "1001"
	server := rpc_group.NewGroupRpcServer(1001)

	req1 := &pb_group.GetGroupMemberListReq{GroupId: groupId,UserId:userId,Page: 2,PageSize: 100}
	server.GetGroupMemberList(context.Background(),req1)

	req := &pb_group.GetGroupAllMemberBasicReq{GroupId: groupId}
	server.GetGroupAllMemberBasic(context.Background(), req)
}
