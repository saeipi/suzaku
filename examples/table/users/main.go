package main

import (
	"fmt"
	"gorm.io/gorm"
	"strconv"
	"suzaku/pkg/common/mysql"
	"suzaku/pkg/common/snowflake"
	"suzaku/pkg/utils"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	var (
		i        int
		user     User
		register Register
		//friend   Friend
		cv   int
		city string

		ulist []User
		rlist []Register
		//flist []Friend

		db  *gorm.DB
		err error
	)
	for i = 20001000; i < 60000000; i++ {
		cv = i % 5
		switch cv {
		case 0: //广州
			city = "广州"
		case 1: //上海
			city = "上海"
		case 2: //北京
			city = "北京"
		case 3: //成都
			city = "成都"
		case 4: //深圳
			city = "深圳"
		}
		user = User{
			UserId:     snowflake.SnowflakeID(),
			SzkId:      strconv.Itoa(i),
			Nickname:   "",
			Gender:     i % 2,
			PlatformId: i % 6,
			City:       city,
		}
		register = Register{
			UserId:   user.UserId,
			Password: utils.MD5(user.UserId),
		}

		ulist = append(ulist, user)
		rlist = append(rlist, register)
		if len(ulist) == 1000 {
			db, err = mysql.GormDB()
			if err != nil {
				fmt.Println("mysql error:", err)
				continue
			}
			err = db.Create(&ulist).Error
			if err != nil {
				fmt.Println("users error:", err)
				continue
			}
			err = db.Create(&rlist).Error
			if err != nil {
				fmt.Println("registers error:", err)
				continue
			}
			ulist = nil
			rlist = nil
		}
	}
	wg.Wait()
}

type User struct {
	UserId     string    `gorm:"column:user_id;primary_key" json:"user_id"`       // 用户ID 系统生成
	SzkId      string    `gorm:"column:szk_id" json:"szk_id"`                     // 账户ID 用户设置
	Nickname   string    `gorm:"column:nickname" json:"nickname"`                 // 昵称
	Gender     int       `gorm:"column:gender;default:0" json:"gender"`           // 性别
	Birth      int       `gorm:"column:birth;default:0" json:"birth"`             // 生日 时间戳
	Email      string    `gorm:"column:email" json:"email"`                       // Email
	Mobile     string    `gorm:"column:mobile" json:"mobile"`                     // 手机号
	PlatformId int       `gorm:"column:platform_id;default:0" json:"platform_id"` // 平台
	AvatarUrl  string    `gorm:"column:avatar_url" json:"avatar_url"`             // 头像
	Country    string    `gorm:"column:country" json:"country"`                   // 国家
	City       string    `gorm:"column:city" json:"city"`                         // 城市
	Ex         string    `gorm:"column:ex" json:"ex"`                             // 备注
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt  time.Time `gorm:"column:deleted_at;default:NULL" json:"deleted_at"`
}

type Register struct {
	UserId    string    `gorm:"column:user_id;primary_key" json:"user_id"` // 用户ID 系统生成
	Password  string    `gorm:"column:password" json:"password"`           // 密码
	Ex        string    `gorm:"column:ex" json:"ex"`                       // 备注
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

type Friend struct {
	OwnerUserId  string    `gorm:"column:owner_user_id;primary_key" json:"owner_user_id"` // 添加好友发起者ID
	FriendUserId string    `gorm:"column:friend_user_id" json:"friend_user_id"`           // 好友ID
	Source       int       `gorm:"column:source;default:0" json:"source"`                 // 添加源
	Remark       string    `gorm:"column:remark" json:"remark"`                           // 备注
	Ex           string    `gorm:"column:ex" json:"ex"`                                   // 扩展字段
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
	DeletedAt    time.Time `gorm:"column:deleted_at;default:NULL" json:"deleted_at"`
}
