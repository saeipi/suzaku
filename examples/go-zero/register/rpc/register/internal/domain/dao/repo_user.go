package dao

import (
	"gorm.io/gorm"
	"suzaku/examples/go-zero/proto/szkproto"
	"suzaku/examples/go-zero/register/rpc/register/internal/domain/po"
	"suzaku/pkg/common/mysql"
	"suzaku/pkg/common/snowflake"
)

type UserRepository interface {
	UserRegister(req *szkproto.UserRegisterReq) (user *po.User, err error)
}

var UserRepo UserRepository

type userRepository struct {
}

func init() {
	UserRepo = new(userRepository)
}

func (r *userRepository) UserRegister(req *szkproto.UserRegisterReq) (user *po.User, err error) {
	err = mysql.Transaction(func(tx *gorm.DB) (terr error) {
		user = &po.User{
			UserId:     snowflake.SnowflakeID(),
			SzkId:      req.SzkId,
			Udid:       req.Udid,
			Status:     1,
			Nickname:   req.Nickname,
			Gender:     0,
			BirthTs:    0,
			Email:      "",
			Mobile:     req.Mobile,
			PlatformId: int(req.PlatformId),
			AvatarUrl:  req.AvatarUrl,
			CityId:     0,
			Ex:         "",
		}
		terr = tx.Create(user).Error
		return
	})
	return
}
