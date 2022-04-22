package http

import "errors"

var (
	ErrorHttpReqDeserializeFailed   = errors.New("请求参数序列化错误")
	ErrorHttpReqParamErr            = errors.New("请求参数错误")
	ErrorHttpReqNotAuthorized       = errors.New("没有授权")
	ErrorHttpRegisterFailed         = errors.New("注册失败")
	ErrorHttpUserIdDoesNotExist     = errors.New("用户ID信息缺失")
	ErrorHttpPlatformIdDoesNotExist = errors.New("平台信息缺失")
	ErrorHttpGetUserFailed          = errors.New("获取用户信息失败")
	ErrorHttpAddFriendFailed        = errors.New("添加好友失败")
)
