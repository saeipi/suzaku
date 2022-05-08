package api_group

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc"
	"suzaku/internal/interface/dto/dto_api"
	"suzaku/pkg/common/config"
	"suzaku/pkg/factory"
	"suzaku/pkg/http"
	pb_group "suzaku/pkg/proto/group"
	"suzaku/pkg/utils"
)

func MemberList(c *gin.Context) {
	var (
		userId     string
		ok         bool
		params     dto_api.GroupMemberListReq
		req        *pb_group.GetGroupMemberListReq
		clientConn *grpc.ClientConn
		client     pb_group.GroupClient
		reply      *pb_group.GetGroupMemberListResp
		resp       *dto_api.GroupMemberListResp
		err        error
	)
	userId, _, ok = utils.RequestIdentity(c)
	if ok == false {
		return
	}
	if err = c.ShouldBindQuery(&params); err != nil {
		http.Error(c, err, http.ErrorCodeHttpReqDeserializeFailed)
		return
	}
	req = new(pb_group.GetGroupMemberListReq)
	copier.Copy(&req, &params)
	req.UserId = userId
	clientConn = factory.ClientConn(config.Config.RPCRegisterName.GroupName)
	client = pb_group.NewGroupClient(clientConn)
	reply, _ = client.GetGroupMemberList(context.Background(), req)

	if reply == nil {
		http.Error(c, http.ErrorHttpServiceFailure, http.ErrorCodeHttpServiceFailure)
		return
	}
	if reply.Common.Code > 0 {
		http.Err(c, reply.Common.Msg, reply.Common.Code)
		return
	}
	resp = new(dto_api.GroupMemberListResp)
	copier.Copy(resp, reply)
	http.Success(c, resp)
}
