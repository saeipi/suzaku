package main

import (
	"database/sql"
	"gorm.io/gorm"
	"suzaku/pkg/common/mysql"
	"suzaku/pkg/constant"
)

func main() {
	var (
		db  *gorm.DB
		err error
	)
	if db, err = mysql.GormDB(); err != nil {
		return
	}

	//for i := 0; i < 10; i++ {
	//	second := Second{Second: int64(i)}
	//	err = db.Save(&second).Error
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//}

	//for i := 0; i < 10; i++ {
	//	err = db.Model(Second{}).Where("id=?",i+1).Update("second",10+1).Error
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//}

	user := User{}
	constant.Create(&user.GormTs)
	db.Save(&user)
}

type Second struct {
	Id        int          `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Second    int64        `gorm:"column:second;default:0" json:"second"`
	CreatedTs sql.NullTime `gorm:"column:created_ts;default:CURRENT_TIMESTAMP" json:"created_ts"`
	UpdatedTs sql.NullTime `gorm:"column:updated_ts;default:CURRENT_TIMESTAMP" json:"updated_ts"`
}

type User struct {
	constant.GormTs
	UserId     string `gorm:"column:user_id;primary_key" json:"user_id"`       // 用户ID 系统生成
	SzkId      string `gorm:"column:szk_id" json:"szk_id"`                     // 账户ID 用户设置
	Udid       string `gorm:"column:udid" json:"udid"`                         // 设备唯一标识
	Status     int32  `gorm:"column:status;default:0" json:"status"`           // 用户状态
	Nickname   string `gorm:"column:nickname" json:"nickname"`                 // 昵称
	Gender     int32  `gorm:"column:gender;default:0" json:"gender"`           // 性别
	BirthTs    int64  `gorm:"column:birth_ts;default:0" json:"birth_ts"`       // 生日
	Email      string `gorm:"column:email" json:"email"`                       // Email
	Mobile     string `gorm:"column:mobile" json:"mobile"`                     // 手机号
	PlatformId int    `gorm:"column:platform_id;default:0" json:"platform_id"` // 注册平台
	AvatarUrl  string `gorm:"column:avatar_url" json:"avatar_url"`             // 头像
	CityId     int    `gorm:"column:city_id;default:0" json:"city_id"`         // 城市ID
	Ex         string `gorm:"column:ex" json:"ex"`                             // 扩展字段
}