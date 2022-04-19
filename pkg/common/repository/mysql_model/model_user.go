package mysql_model

import (
	"suzaku/pkg/constant"
)

type User struct {
	constant.GormModel
	UserId   string `gorm:"column:user_id" json:"user_id"`             // 用户ID
	Mobile   string `gorm:"column:mobile" json:"mobile"`               // 手机
	Platform int32  `gorm:"column:platform;default:0" json:"platform"` // 平台
	Gender   int32  `gorm:"column:gender;default:0" json:"gender"`     // 性别
}
