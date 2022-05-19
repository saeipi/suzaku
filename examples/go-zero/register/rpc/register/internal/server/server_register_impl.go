package server

import (
	"context"
	"suzaku/examples/go-zero/proto/szkproto"
	"suzaku/examples/go-zero/register/rpc/register/internal/handler"
)

func (rpc *authRpcServer) UserRegister(ctx context.Context, req *szkproto.UserRegisterReq) (resp *szkproto.UserRegisterResp, _ error) {
	return handler.RegisterHandler.UserRegister(ctx, req)
}
