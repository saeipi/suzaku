package repo_mysql

import (
	"gorm.io/gorm"
	"suzaku/internal/domain/po_mysql"
	"suzaku/pkg/common/mysql"
	"suzaku/pkg/common/snowflake"
	pb_auth "suzaku/pkg/proto/auth"
	pb_user "suzaku/pkg/proto/pb_user"
)

type UserRepository interface {
	UserRegister(req *pb_auth.UserRegisterReq) (user *po_mysql.User, err error)
	GetUserByUserID(userID string) (user *po_mysql.User, err error)
	TxGetUserByUserID(userID string, tx *gorm.DB) (user *po_mysql.User, err error)
	TxGetAvatarByUserID(userID string, tx *gorm.DB) (avatar *po_mysql.UserAvatar, err error)
	GetUserBySzkID(szkID string) (user *po_mysql.User, err error)
	GetFromToUserNickname(fromUserID, toUserID string) (fromUserNickname string, toUserNickname string, err error)
	EditUserInfo(req *pb_user.EditUserInfoReq) (err error)
}

var UserRepo UserRepository

type userRepository struct {
}

func init() {
	UserRepo = new(userRepository)
}

func (r *userRepository) UserRegister(req *pb_auth.UserRegisterReq) (user *po_mysql.User, err error) {
	err = mysql.Transaction(func(tx *gorm.DB) (terr error) {
		var (
			avatar   *po_mysql.UserAvatar
			register *po_mysql.Register
		)

		avatar = new(po_mysql.UserAvatar)
		//1
		user = &po_mysql.User{
			UserId:     snowflake.SnowflakeID(),
			Mobile:     req.Mobile,
			PlatformId: int(req.PlatformId),
		}
		terr = tx.Save(user).Error
		if terr != nil {
			return
		}
		//2
		register = &po_mysql.Register{
			UserId:   user.UserId,
			Password: req.Password,
		}
		terr = tx.Save(register).Error
		if terr != nil {
			return
		}
		//3
		avatar.UserId = user.UserId
		terr = tx.Save(avatar).Error
		return
	})
	return
}

func (r *userRepository) GetUserByUserID(userID string) (user *po_mysql.User, err error) {
	var (
		db *gorm.DB
	)
	db, err = mysql.GormDB()
	if err != nil {
		return
	}
	err = db.Where("user_id=?", userID).Find(&user).Error
	return
}
func (r *userRepository) TxGetUserByUserID(userID string, tx *gorm.DB) (user *po_mysql.User, err error) {
	err = tx.Where("user_id=?", userID).Find(&user).Error
	return
}

func (r *userRepository) TxGetAvatarByUserID(userID string, tx *gorm.DB) (avatar *po_mysql.UserAvatar, err error) {
	err = tx.Where("user_id=?", userID).Find(&avatar).Error
	return
}

func (r *userRepository) GetUserBySzkID(szkID string) (user *po_mysql.User, err error) {
	var (
		db *gorm.DB
	)
	user = new(po_mysql.User)
	db, err = mysql.GormDB()
	if err != nil {
		return
	}
	err = db.Where("szk_id=?", szkID).Find(user).Error
	return
}

func (r *userRepository) GetFromToUserNickname(fromUserID, toUserID string) (fromUserNickname string, toUserNickname string, err error) {
	var (
		fromUser *po_mysql.User
		toUser   *po_mysql.User
	)
	fromUser, err = UserRepo.GetUserByUserID(fromUserID)
	if err != nil {
		return
	}
	toUser, err = UserRepo.GetUserByUserID(toUserID)
	if err != nil {
		return
	}
	fromUserNickname = fromUser.Nickname
	toUserNickname = toUser.Nickname
	return
}

func (r *userRepository) EditUserInfo(req *pb_user.EditUserInfoReq) (err error) {
	var (
		updates map[string]interface{}
		db      *gorm.DB
	)
	db, err = mysql.GormDB()
	if err != nil {
		return
	}
	updates = make(map[string]interface{})
	updates["szk_id"] = req.SzkId
	updates["nickname"] = req.Nickname
	updates["gender"] = req.Gender
	updates["birth_ts"] = req.BirthTs
	updates["email"] = req.Email
	updates["mobile"] = req.Mobile
	updates["avatar_url"] = req.AvatarUrl
	updates["city_id"] = req.CityId
	err = db.Model(po_mysql.User{}).Where("user_id=?", req.UserId).Updates(updates).Error
	return
}