package handler

import (
	"context"
	"suzaku/examples/go-zero/proto/szkproto"
	"suzaku/examples/go-zero/register/rpc/register/internal/domain/dao"
)

var (
	RegisterHandler *registerHandler
)

type registerHandler struct {
}

func init() {
	RegisterHandler = new(registerHandler)
}

func (r *registerHandler) UserRegister(ctx context.Context, req *szkproto.UserRegisterReq) (resp *szkproto.UserRegisterResp, _ error) {
	var (
		common = &szkproto.CommonResp{}
		err    error
	)
	resp = &szkproto.UserRegisterResp{Common: common}
	_, err = dao.UserRepo.UserRegister(req)
	if err != nil {
		//TODO: Error
		common.Msg = err.Error()
		common.Code = 777
		return
	}
	return
}
