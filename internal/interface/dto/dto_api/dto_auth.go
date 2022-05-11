package dto_api

// 1、注册
type UserRegisterReq struct {
	Secret           string `json:"secret" binding:"required,max=32"`           // 秘钥，最大长度 32 字符
	PlatformId       int32  `json:"platform_id" binding:"required,min=1,max=7"` // 平台类型 iOS 1, Android 2, Windows 3, OSX 4, WEB 5, 小程序 6，linux 7
	Mobile           string `json:"mobile" binding:"omitempty,max=32"`          // 用户 mobile，最大长度 32 字符，非中国大陆手机号码需要填写国家代码(如美国：+1-xxxxxxxxxx)或地区代码(如香港：+852-xxxxxxxx)，可设置为空字符串
	UDID             string `json:"udid" binding:"required,max=40"`             // 设备唯一标识
	VerificationCode string `json:"verification_code" binding:"required,max=6"` // 验证码
	Password         string `json:"password" binding:"required,min=32,max=32"`  // 密码
}

type UserRegisterResp struct {
	PlatformId int32  `json:"platform_id"` // 平台
	UserId     string `json:"user_id"`     // 用户ID
	Token      string `json:"token"`       // token
	Expire     int64  `json:"expire"`      // token 有效期
}

type UserTokenReq struct {
	Secret     string `json:"secret" binding:"required,max=32"`           // 秘钥，最大长度 32 字符
	PlatformId int32  `json:"platform_id" binding:"required,min=1,max=7"` // 平台类型 iOS 1, Android 2, Windows 3, OSX 4, WEB 5, 小程序 6，linux 7
	Mobile     string `json:"mobile" binding:"omitempty,max=32"`          // 用户 mobile，最大长度 32 字符，非中国大陆手机号码需要填写国家代码(如美国：+1-xxxxxxxxxx)或地区代码(如香港：+852-xxxxxxxx)，可设置为空字符串
}

type UserLoginReq struct {
	LoginId          string `json:"login_id" binding:"required,max=40"`         // 用户ID
	Password         string `json:"password" binding:"required,min=32,max=32"`  // 密码
	VerificationCode string `json:"verification_code" binding:"required,max=6"` // 验证码 TODO:暂未处理
	PlatformId       int32  `json:"platform_id" binding:"required,min=1,max=7"` // 平台类型ID
	LoginType        int32  `json:"login_type" binding:"required,min=1,max=7"`  // 1:szk_id登录
}

type UserToken struct {
	Token  string `json:"token"`  // token
	Expire int64  `json:"expire"` // token 有效期
}

type UserLoginResp struct {
	UserInfo UserInfo  `json:"user_info"`
	Token    UserToken `json:"token"`
}
