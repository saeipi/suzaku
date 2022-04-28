package api_friend

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"suzaku/internal/interface/dto/dto_api"
	"suzaku/pkg/common/config"
	"suzaku/pkg/factory"
	"suzaku/pkg/http"
	pb_friend "suzaku/pkg/proto/friend"
	"suzaku/pkg/utils"
)

func AddFriend(c *gin.Context) {
	var (
		userId     string
		ok         bool
		params     dto_api.AddFriendReq
		err        error
		req        *pb_friend.AddFriendReq
		clientConn *grpc.ClientConn
		client     pb_friend.FriendClient
		reply      *pb_friend.AddFriendResp
	)
	userId, _, ok = utils.RequestIdentity(c)
	if ok == false {
		return
	}
	if err = c.BindJSON(&params); err != nil {
		http.Error(c, err, http.ErrorCodeHttpReqDeserializeFailed)
		return
	}

	req = &pb_friend.AddFriendReq{}
	utils.CopyStructFields(req, params)
	req.UserId = userId

	clientConn = factory.ClientConn(config.Config.RPCRegisterName.FriendName)
	client = pb_friend.NewFriendClient(clientConn)
	reply, _ = client.AddFriend(context.Background(), req)
	if reply.Common != nil && reply.Common.Code > 0 {
		http.Err(c, reply.Common.Msg, reply.Common.Code)
		return
	}
	http.Success(c, reply)
}
