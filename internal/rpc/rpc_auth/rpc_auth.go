package rpc_auth

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc"
	"suzaku/internal/domain/po_mysql"
	"suzaku/internal/domain/repo/repo_mysql"
	"suzaku/internal/rpc/rpc_category"
	"suzaku/pkg/common/config"
	"suzaku/pkg/common/jwt_auth"
	"suzaku/pkg/common/redis"
	"suzaku/pkg/factory"
	pb_auth "suzaku/pkg/proto/auth"
	pb_com "suzaku/pkg/proto/pb_com"
	"suzaku/pkg/proto/pb_user"
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
	server = factory.NewRPCNewServer()
	pb_auth.RegisterAuthServer(server, rpc)
	rpc.Rpc.RunServer(server)
}

func (rpc *authRpcServer) UserRegister(ctx context.Context, req *pb_auth.UserRegisterReq) (resp *pb_auth.UserRegisterResp, _ error) {
	var (
		common = &pb_com.CommonResp{}
		user   *po_mysql.User
		err    error
	)
	resp = &pb_auth.UserRegisterResp{
		Common:   common,
		UserInfo: &pb_user.UserInfo{},
		Token:    &pb_auth.UserToken{},
	}
	user, err = repo_mysql.UserRepo.GetUserBySzkID(req.SzkId)
	if err != nil {
		//TODO: Error
		common.Msg = err.Error()
		common.Code = 777
		return
	}
	if user.UserId != "" {
		//TODO: szk_id 已经存在
		common.Code = 777
		return
	}
	user, err = repo_mysql.UserRepo.UserRegister(req)
	if err != nil {
		common.Msg = err.Error()
		common.Code = ErrCodeRpcRegisterFailed
		return
	}
	copier.Copy(resp.UserInfo, user)

	resp.Token.Token, resp.Token.Expire = jwt_auth.CreateJwtToken(user.UserId, req.PlatformId)

	// TODO:调试用
	if resp.Token.Expire > 0 {
		err = redis.Set(fmt.Sprintf(redis.RedisKeyJwtUserToken, user.UserId, req.PlatformId), resp.Token.Token, int(resp.Token.Expire)*1000)
	}
	return
}

func (rpc *authRpcServer) UserToken(ctx context.Context, req *pb_auth.UserTokenReq) (resp *pb_auth.UserTokenResp, err error) {
	var (
		token  string
		expire int64
	)
	token, expire = jwt_auth.CreateJwtToken(req.UserId, req.PlatformId)
	resp = &pb_auth.UserTokenResp{
		Common: &pb_com.CommonResp{},
		Token:  token,
		Expire: expire,
	}
	// TODO:调试用
	if expire > 0 {
		err = redis.Set(fmt.Sprintf(redis.RedisKeyJwtUserToken, req.UserId, req.PlatformId), token, int(expire)*1000)
	}
	return
}

func (rpc *authRpcServer) UserLogin(_ context.Context, req *pb_auth.UserLoginReq) (resp *pb_auth.UserLoginResp, _ error) {
	var (
		register *po_mysql.Register
		user     *po_mysql.User
		err      error
	)
	resp = &pb_auth.UserLoginResp{
		Common:   &pb_com.CommonResp{},
		UserInfo: &pb_user.UserInfo{},
		Token:    &pb_auth.UserToken{},
	}
	user, err = repo_mysql.UserRepo.GetUserBySzkID(req.LoginId)
	if err != nil {
		//TODO:error
		resp.Common.Msg = err.Error()
		resp.Common.Code = 777
		return
	}
	if user.UserId == "" {
		//TODO:用户不存在
		resp.Common.Code = 777
		return
	}
	register, err = repo_mysql.AuthRepo.VerifyPassword(user.UserId, req.Password)
	if err != nil {
		//TODO:error
		resp.Common.Msg = err.Error()
		resp.Common.Code = 777
		return
	}
	if register.UserId == "" {
		//TODO:密码错误
		resp.Common.Code = 777
		return
	}

	copier.Copy(resp.UserInfo, user)
	resp.Token.Token, resp.Token.Expire = jwt_auth.CreateJwtToken(user.UserId, req.PlatformId)

	// TODO:调试用
	if resp.Token.Expire > 0 {
		err = redis.Set(fmt.Sprintf(redis.RedisKeyJwtUserToken, user.UserId, req.PlatformId), resp.Token.Token, int(resp.Token.Expire)*1000)
	}
	return
}
