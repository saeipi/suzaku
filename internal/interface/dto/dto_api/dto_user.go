package dto_api

type UserInfo struct {
	UserId    string `json:"user_id"`    // 用户ID
	SzkId     string `json:"szk_id"`     // 账户ID 用户设置 (一年只允许修改一次)
	Nickname  string `json:"nickname"`   // 昵称
	Gender    int32  `json:"gender"`     // 性别
	BirthTs   int64  `json:"birth_ts"`   // 生日
	Email     string `json:"email"`      // Email
	Mobile    string `json:"mobile"`     // 手机号
	AvatarUrl string `json:"avatar_url"` // 头像
	CityId    int    `json:"city_id"`    // 城市ID
}

type UserInfoResp struct {
	UserInfo
	//TODO:尚未开发
	Country string `json:"country"`
	City string `json:"city"`
}

type EditUserInfoReq struct {
	UserInfo
}
