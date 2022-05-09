package api_friend

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
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
	req.FromUserId = userId

	clientConn = factory.ClientConn(config.Config.RPCRegisterName.FriendName)
	client = pb_friend.NewFriendClient(clientConn)
	reply, _ = client.AddFriend(context.Background(), req)
	if reply == nil {
		http.Error(c, http.ErrorHttpServiceFailure, http.ErrorCodeHttpServiceFailure)
		return
	}
	if reply.Common.Code > 0 {
		http.Err(c, reply.Common.Msg, reply.Common.Code)
		return
	}
	http.Success(c)
}

func FriendRequestList(c *gin.Context) {
	var (
		userId     string
		ok         bool
		params     dto_api.FriendRequestListReq
		req        *pb_friend.GetFriendRequestListReq
		clientConn *grpc.ClientConn
		client     pb_friend.FriendClient
		reply      *pb_friend.GetFriendRequestListResp
		resp       *dto_api.FriendRequestListResp
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
	req = &pb_friend.GetFriendRequestListReq{
		UserId:   userId,
		Role:     pb_friend.FRIEND_REQUEST_ROLE(params.Role),
		Page:     int32(params.Page),
		PageSize: int32(params.PageSize),
	}
	clientConn = factory.ClientConn(config.Config.RPCRegisterName.FriendName)
	client = pb_friend.NewFriendClient(clientConn)
	reply, _ = client.GetFriendRequestList(context.Background(), req)
	if reply == nil {
		http.Error(c, http.ErrorHttpServiceFailure, http.ErrorCodeHttpServiceFailure)
		return
	}
	if reply.Common.Code > 0 {
		http.Err(c, reply.Common.Msg, reply.Common.Code)
		return
	}
	resp = &dto_api.FriendRequestListResp{}
	copier.Copy(resp, reply)
	http.Success(c, resp)
}

func HandleFriendRequest(c *gin.Context) {
	var (
		userId string
		ok     bool
		params dto_api.HandleFriendRequestReq
		req    *pb_friend.HandleFriendRequestReq
		reply  *pb_friend.HandleFriendRequestResp

		clientConn *grpc.ClientConn
		client     pb_friend.FriendClient
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

	req = &pb_friend.HandleFriendRequestReq{}
	utils.CopyStructFields(req, params)
	req.UserId = userId

	clientConn = factory.ClientConn(config.Config.RPCRegisterName.FriendName)
	client = pb_friend.NewFriendClient(clientConn)
	reply, _ = client.HandleFriendRequest(context.Background(), req)
	if reply == nil {
		http.Error(c, http.ErrorHttpServiceFailure, http.ErrorCodeHttpServiceFailure)
		return
	}
	if reply.Common.Code > 0 {
		http.Err(c, reply.Common.Msg, reply.Common.Code)
		return
	}
	http.Success(c)
}
