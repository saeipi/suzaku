package api_user

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"suzaku/internal/interface/dto/dto_api"
	"suzaku/pkg/common/config"
	"suzaku/pkg/factory"
	"suzaku/pkg/http"
	"suzaku/pkg/proto/pb_com"
	pb_user "suzaku/pkg/proto/user"
	"suzaku/pkg/utils"
)

func SelfInfo(c *gin.Context) {
	var (
		userId     string
		ok         bool
		req        *pb_user.UserInfoReq
		clientConn *grpc.ClientConn
		client     pb_user.UserClient
		reply      *pb_user.UserInfoResp
		resp       *dto_api.UserInfoResp
	)
	userId, _, ok = utils.RequestIdentity(c)
	if ok == false {
		return
	}
	req = &pb_user.UserInfoReq{UserId: userId}
	clientConn = factory.ClientConn(config.Config.RPCRegisterName.UserName)
	client = pb_user.NewUserClient(clientConn)
	reply, _ = client.UserInfo(context.Background(), req)
	if reply == nil {
		http.Error(c, http.ErrorHttpServiceFailure, http.ErrorCodeHttpServiceFailure)
		return
	}
	if reply.Common.Code > 0 {
		http.Err(c, reply.Common.Msg, reply.Common.Code)
		return
	}
	resp = &dto_api.UserInfoResp{}
	utils.CopyStructFields(resp, reply)
	http.Success(c, resp)
}

func EditInfo(c *gin.Context) {
	var (
		userId     string
		ok         bool
		params     dto_api.EditUserInfoReq
		req        *pb_user.EditUserInfoReq
		clientConn *grpc.ClientConn
		client     pb_user.UserClient
		reply      *pb_com.CommonResp
		err        error
	)
	userId, _, ok = utils.RequestIdentity(c)
	if ok == false {
		return
	}
	if err = c.BindJSON(&params); err != nil {
		http.Error(c, err, http.ErrorCodeHttpReqDeserializeFailed)
		return
	}

	req = &pb_user.EditUserInfoReq{}
	utils.CopyStructFields(req, params)
	req.UserId = userId

	clientConn = factory.ClientConn(config.Config.RPCRegisterName.UserName)
	client = pb_user.NewUserClient(clientConn)
	reply, _ = client.EditUserInfo(context.Background(), req)
	if reply == nil {
		http.Error(c, http.ErrorHttpServiceFailure, http.ErrorCodeHttpServiceFailure)
		return
	}
	if reply.Code > 0 {
		http.Err(c, reply.Msg, reply.Code)
		return
	}
	http.Success(c)
}
