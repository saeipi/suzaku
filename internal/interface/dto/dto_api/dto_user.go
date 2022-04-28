package dto_api

import "time"

type UserInfoResp struct {
	UserId     string    `gorm:"column:user_id;primary_key" json:"user_id"`       // 用户ID 系统生成
	SzkId      string    `gorm:"column:szk_id" json:"szk_id"`                     // 账户ID 用户设置
	Nickname   string    `gorm:"column:nickname" json:"nickname"`                 // 昵称
	Gender     int       `gorm:"column:gender;default:0" json:"gender"`           // 性别
	Birth      time.Time `gorm:"column:birth" json:"birth"`                       // 生日
	Email      string    `gorm:"column:email" json:"email"`                       // Email
	Mobile     string    `gorm:"column:mobile" json:"mobile"`                     // 手机号
	PlatformId int       `gorm:"column:platform_id;default:0" json:"platform_id"` // 注册平台
	AvatarUrl  string    `gorm:"column:avatar_url" json:"avatar_url"`             // 头像
	Country    string    `gorm:"column:country" json:"country"`                   // 国家
	City       string    `gorm:"column:city" json:"city"`                         // 城市
}
