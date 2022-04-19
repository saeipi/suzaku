package rpc_auth

import (
	"context"
	"suzaku/pkg/common/config"
	"google.golang.org/grpc"
	"suzaku/internal/rpc/rpc_category"
	"suzaku/pkg/common/jwt_auth"
	"suzaku/pkg/common/repository/facade/facade_user"
	"suzaku/pkg/common/repository/mysql_model"
	"suzaku/pkg/common/snowflake"
	pb_auth "suzaku/pkg/proto/auth"
	pb_com "suzaku/pkg/proto/pb_com"
)

type authRpc struct {
	pb_auth.UnimplementedAuthServer
	rpc_category.Rpc
}

func NewRpcAuthServer(port int) (r *authRpc) {
	return &authRpc{
		Rpc: rpc_category.NewRpcServer(port,config.Config.RPCRegisterName.AuthName),
	}
}

func (rpc *authRpc) Run() {
	var (
		server *grpc.Server
	)
	server = grpc.NewServer()
	pb_auth.RegisterAuthServer(server, rpc)
	rpc.Rpc.RunServer(server)
}

func (rpc *authRpc) UserRegister(ctx context.Context, req *pb_auth.UserRegisterReq) (resp *pb_auth.UserRegisterResp, err error) {
	var (
		user   *mysql_model.User
		common = &pb_com.CommonResp{}
	)
	resp = &pb_auth.UserRegisterResp{Common: common}

	user = &mysql_model.User{
		UserId:   snowflake.SnowflakeID(),
		Mobile:   req.Mobile,
		Platform: req.Platform,
	}
	err = facade_user.UserRegister(user)
	if err != nil {
		common.Msg = err.Error()
		common.Code = ErrCodeRpcRegisterFailed
		return
	}
	resp.Platform = user.Platform
	resp.Id = user.ID
	resp.UserId = user.UserId
	return
}

func (rpc *authRpc) UserToken(ctx context.Context, req *pb_auth.UserTokenReq) (resp *pb_auth.UserTokenResp, err error) {
	var (
		token  string
		expire int64
	)
	token, expire = jwt_auth.CreateJwtToken(req.UserId, req.Platform)
	resp = &pb_auth.UserTokenResp{
		Token:  token,
		Expire: expire,
	}
	/*
		if expire > 0 {
			err = redis.Set(fmt.Sprintf(redis.RedisKeyJwtUserTokenKey, req.UserId, req.Platform), token, int(expire)*1000)
		}
	*/
	return
}
