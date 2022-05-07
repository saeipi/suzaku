package rpc_group

import (
	"context"
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
	key = "group:" + req.GroupId
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
