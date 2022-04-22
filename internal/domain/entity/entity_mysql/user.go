package entity_mysql

import (
	"suzaku/pkg/constant"
)

type User struct {
	constant.GormModel
	UserId     string `gorm:"column:user_id" json:"user_id"`                   // 用户ID
	Mobile     string `gorm:"column:mobile" json:"mobile"`                     // 手机
	PlatformId int32  `gorm:"column:platform_id;default:0" json:"platform_id"` // 平台
	Gender     int32  `gorm:"column:gender;default:0" json:"gender"`           // 性别
}
