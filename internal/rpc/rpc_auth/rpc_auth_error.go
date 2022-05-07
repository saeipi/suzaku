package rpc_auth

import "errors"

var (
	ErrCodeRpcRegisterFailed int32 = 100001
)

var (
	ErrRpcRegisterFailed = errors.New("注册失败")
)
