package dto_api

type UserInfoResp struct {
	UserId     string `gorm:"column:user_id;primary_key" json:"user_id"`       // 用户ID 系统生成
	SzkId      string `gorm:"column:szk_id" json:"szk_id"`                     // 账户ID 用户设置
	Udid       string `gorm:"column:udid" json:"udid"`                         // 设备唯一标识
	Nickname   string `gorm:"column:nickname" json:"nickname"`                 // 昵称
	Gender     int    `gorm:"column:gender;default:0" json:"gender"`           // 性别
	BirthTs    int64  `gorm:"column:birth_ts;default:0" json:"birth_ts"`       // 生日
	Email      string `gorm:"column:email" json:"email"`                       // Email
	Mobile     string `gorm:"column:mobile" json:"mobile"`                     // 手机号
	PlatformId int    `gorm:"column:platform_id;default:0" json:"platform_id"` // 注册平台
	AvatarUrl  string `gorm:"column:avatar_url" json:"avatar_url"`             // 头像
	CityId     int    `gorm:"column:city_id;default:0" json:"city_id"`         // 城市ID
	Ex         string `gorm:"column:ex" json:"ex"`                             // 扩展字段
}
