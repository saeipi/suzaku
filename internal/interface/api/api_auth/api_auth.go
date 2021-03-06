package api_auth

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc"
	"suzaku/internal/interface/dto/dto_api"
	"suzaku/pkg/common/config"
	"suzaku/pkg/factory"
	"suzaku/pkg/http"
	pb_auth "suzaku/pkg/proto/auth"
	"suzaku/pkg/utils"
)

func UserRegister(c *gin.Context) {
	var (
		params     dto_api.UserRegisterReq
		err        error
		clientConn *grpc.ClientConn
		client     pb_auth.AuthClient
		req        *pb_auth.UserRegisterReq
		reply      *pb_auth.UserRegisterResp
		tokenReq   *pb_auth.UserTokenReq
		replyToken *pb_auth.UserTokenResp
		resp       *dto_api.UserRegisterResp
	)

	if err = c.BindJSON(&params); err != nil {
		http.Error(c, err, http.ErrorCodeHttpReqDeserializeFailed)
		return
	}
	if params.Secret != config.Config.Secret {
		http.Error(c, err, http.ErrorCodeHttpReqNotAuthorized)
		return
	}
	req = &pb_auth.UserRegisterReq{}
	utils.CopyStructFields(req, params)
	//clientConn = getcdv3.GetConn(config.Config.Etcd.Schema, strings.Join(config.Config.Etcd.Address, ","), config.Config.RPCRegisterName.AuthName)
	clientConn = factory.ClientConn(config.Config.RPCRegisterName.AuthName)
	client = pb_auth.NewAuthClient(clientConn)
	reply, _ = client.UserRegister(context.Background(), req)
	if reply == nil {
		http.Error(c, http.ErrorHttpServiceFailure, http.ErrorCodeHttpServiceFailure)
		return
	}
	if reply.Common.Code > 0 {
		http.Err(c, reply.Common.Msg, reply.Common.Code)
		return
	}

	tokenReq = &pb_auth.UserTokenReq{}
	utils.CopyStructFields(tokenReq, reply)
	replyToken, _ = client.UserToken(context.Background(), tokenReq)
	if replyToken == nil {
		http.Error(c, http.ErrorHttpServiceFailure, http.ErrorCodeHttpServiceFailure)
		return
	}
	if replyToken.Common.Code > 0 {
		http.Err(c, replyToken.Common.Msg, replyToken.Common.Code)
		return
	}

	resp = new(dto_api.UserRegisterResp)
	copier.Copy(resp, reply)
	http.Success(c, resp)
}

func UserToken(c *gin.Context) {

}

func Login(c *gin.Context) {
	var (
		params     dto_api.UserLoginReq
		err        error
		req        *pb_auth.UserLoginReq
		clientConn *grpc.ClientConn
		client     pb_auth.AuthClient
		reply      *pb_auth.UserLoginResp
		resp       *dto_api.UserLoginResp
	)
	if err = c.BindJSON(&params); err != nil {
		http.Error(c, err, http.ErrorCodeHttpReqDeserializeFailed)
		return
	}
	req = &pb_auth.UserLoginReq{}
	copier.Copy(req, params)

	clientConn = factory.ClientConn(config.Config.RPCRegisterName.AuthName)
	client = pb_auth.NewAuthClient(clientConn)
	reply, _ = client.UserLogin(context.Background(), req)
	if reply == nil {
		http.Error(c, http.ErrorHttpServiceFailure, http.ErrorCodeHttpServiceFailure)
		return
	}
	if reply.Common.Code > 0 {
		http.Err(c, reply.Common.Msg, reply.Common.Code)
		return
	}
	resp = new(dto_api.UserLoginResp)
	copier.Copy(resp, reply)
	http.Success(c, resp)
}
