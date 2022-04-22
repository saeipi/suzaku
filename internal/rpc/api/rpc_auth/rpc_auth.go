package rpc_auth

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"suzaku/internal/domain/entity/entity_mysql"
	"suzaku/internal/domain/repository/repository_mysql"
	"suzaku/internal/rpc/rpc_category"
	"suzaku/pkg/common/config"
	"suzaku/pkg/common/jwt_auth"
	"suzaku/pkg/common/redis"
	"suzaku/pkg/common/snowflake"
	pb_auth "suzaku/pkg/proto/auth"
	pb_com "suzaku/pkg/proto/pb_com"
)

type authRpcServer struct {
	pb_auth.UnimplementedAuthServer
	rpc_category.Rpc
}

func NewAuthRpcServer(port int) (r *authRpcServer) {
	return &authRpcServer{
		Rpc: rpc_category.NewRpcServer(port, config.Config.RPCRegisterName.AuthName),
	}
}

func (rpc *authRpcServer) Run() {
	var (
		server *grpc.Server
	)
	server = grpc.NewServer()
	pb_auth.RegisterAuthServer(server, rpc)
	rpc.Rpc.RunServer(server)
}

func (rpc *authRpcServer) UserRegister(ctx context.Context, req *pb_auth.UserRegisterReq) (resp *pb_auth.UserRegisterResp, err error) {
	var (
		user   *entity_mysql.User
		common = &pb_com.CommonResp{}
	)
	resp = &pb_auth.UserRegisterResp{Common: common}

	user = &entity_mysql.User{
		UserId:     snowflake.SnowflakeID(),
		Mobile:     req.Mobile,
		PlatformId: req.PlatformId,
	}

	err = repository_mysql.UserRepo.UserRegister(user)
	if err != nil {
		common.Msg = err.Error()
		common.Code = ErrCodeRpcRegisterFailed
		return
	}
	resp.PlatformId = user.PlatformId
	resp.Id = user.ID
	resp.UserId = user.UserId
	return
}

func (rpc *authRpcServer) UserToken(ctx context.Context, req *pb_auth.UserTokenReq) (resp *pb_auth.UserTokenResp, err error) {
	var (
		token  string
		expire int64
	)
	token, expire = jwt_auth.CreateJwtToken(req.UserId, req.PlatformId)
	resp = &pb_auth.UserTokenResp{
		Token:  token,
		Expire: expire,
	}
	// TODO:调试用
	if expire > 0 {
		err = redis.Set(fmt.Sprintf(redis.RedisKeyJwtUserTokenKey, req.UserId, req.PlatformId), token, int(expire)*1000)
	}
	return
}
