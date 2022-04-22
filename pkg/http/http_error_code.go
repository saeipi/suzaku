package http

const (
	ErrorCodeHttpReqDeserializeFailed   = 10001 // 请求参数序列化错误
	ErrorCodeHttpReqParamErr            = 10002 // 请求参数错误
	ErrorCodeHttpReqNotAuthorized       = 10003 // 请求参数错误
	ErrorCodeHttpRegisterFailed         = 10004 // 注册失败
	ErrorCodeHttpTokenFailed            = 10005 // 获取token失败
	ErrorCodeHttpJwtTokenErr            = 10006 // token 错误
	ErrorCodeHttpUserIdDoesNotExist     = 10007 // 用户ID信息缺失
	ErrorCodeHttpPlatformIdDoesNotExist = 10008 // 平台信息缺失
	ErrorCodeHttpGetUserFailed          = 10009 // 获取用户信息失败
	ErrorCodeHttpAddFriendFailed        = 10010 // 添加好友失败
)

const (
	ErrorCodeHttp400 = 400
)
