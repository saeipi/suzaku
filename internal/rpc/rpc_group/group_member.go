package rpc_group

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"suzaku/internal/domain/po_mysql"
	"suzaku/internal/domain/repo/repo_mysql"
	"suzaku/pkg/common/redis"
	pb_group "suzaku/pkg/proto/group"
	"suzaku/pkg/proto/pb_com"
	"suzaku/pkg/utils"
)

func (rpc *groupRpcServer) GetGroupAllMemberBasic(_ context.Context, req *pb_group.GetGroupAllMemberBasicReq) (resp *pb_group.GetGroupAllMemberBasicResp, _ error) {
	var (
		key     string
		valJson string
		members []*po_mysql.GroupMember
		err     error
	)
	resp = &pb_group.GetGroupAllMemberBasicResp{Common: &pb_com.CommonResp{}}
	key = fmt.Sprintf(redis.RedisKeyGroup, req.GroupId)
	valJson, err = redis.Get(key)
	if err == nil && valJson != "" {
		utils.JsonToObj(valJson, &resp.MemberList)
		return
	}
	if err == redis.ErrorRedisNil {
		err = nil
	}
	if err != nil {
		// TODO:错误
		resp.Common.Code = 101
		resp.Common.Msg = err.Error()
		return
	}
	if valJson == "" {
		members, err = repo_mysql.GroupRepo.AllMember(req.GroupId)
		if err != nil {
			// TODO:错误
			resp.Common.Code = 101
			resp.Common.Msg = err.Error()
			return
		}
		copier.Copy(&resp.MemberList, &members)
		if len(members) == 0 {
			return
		}
		valJson, _ = utils.ObjToJson(resp.MemberList)
		err = redis.Set(key, valJson, 0)
		if err != nil {
			// TODO:错误
			resp.Common.Code = 101
			resp.Common.Msg = err.Error()
			return
		}
	}
	return
}

func (rpc *groupRpcServer) GetGroupMemberList(_ context.Context, req *pb_group.GetGroupMemberListReq) (resp *pb_group.GetGroupMemberListResp, _ error) {
	var (
		isJoined  bool
		members   []*po_mysql.GroupMember
		totalRows int64
		err       error
	)
	resp = &pb_group.GetGroupMemberListResp{Common: &pb_com.CommonResp{}}
	isJoined, _ = rpc.IsJoined(req.GroupId, req.UserId)
	if isJoined == false {
		//TODO:非群成员
		return
	}
	members, totalRows, err = repo_mysql.GroupRepo.MemberList(req)
	if err != nil {
		//TODO:错误
		return
	}
	resp.TotalRows = totalRows
	copier.Copy(&resp.MemberList, &members)
	return
}

func (rpc *groupRpcServer) IsJoined(groupId string, userId string) (isJoined bool, err error) {
	var (
		member *po_mysql.GroupMember
		val    string
		key    string
	)
	key = fmt.Sprintf(redis.RedisKeyGroupMember, groupId, userId)
	val, err = redis.Get(key)
	if err == nil && val != "" {
		isJoined = true
		return
	}

	if err == nil && val == "" {
		member, err = repo_mysql.GroupRepo.IsJoined(groupId, userId)
		if err != nil {
			//TODO:错误
			return
		}
		if member.UserId != "" {
			isJoined = true
			redis.Set(key, member.UserId, 0)
		}
		return
	}
	return
}
