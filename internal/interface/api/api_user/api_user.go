package api_user

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"suzaku/internal/dto/dto_api"
	"suzaku/pkg/common/config"
	"suzaku/pkg/factory"
	"suzaku/pkg/http"
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
	//clientConn = getcdv3.GetConn(config.Config.Etcd.Schema, strings.Join(config.Config.Etcd.Address, ","), config.Config.RPCRegisterName.AuthName)
	clientConn = factory.ClientConn(config.Config.RPCRegisterName.AuthName)
	client = pb_user.NewUserClient(clientConn)
	reply, _ = client.UserInfo(context.Background(), req)
	if reply.Common != nil && reply.Common.Code > 0 {
		http.Err(c, reply.Common.Msg, reply.Common.Code)
		return
	}
	resp = &dto_api.UserInfoResp{}
	utils.CopyStructFields(resp, reply)
	http.Success(c, resp)
}
