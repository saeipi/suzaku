package repo_mysql

import (
	"gorm.io/gorm"
	"suzaku/internal/domain/po_mysql"
	"suzaku/pkg/common/mysql"
)

type UserRepository interface {
	UserRegister(user *po_mysql.User) (err error)
	GetUserByUserID(userID string) (user *po_mysql.User, err error)
	TxGetUserByUserID(userID string, tx *gorm.DB) (user *po_mysql.User, err error)
	TxGetAvatarByUserID(userID string, tx *gorm.DB) (avatar *po_mysql.UserAvatar, err error)
	GetUserBySzkID(szkID string) (user *po_mysql.User, err error)
	GetFromToUserNickname(fromUserID, toUserID string) (fromUserNickname string, toUserNickname string, err error)
}

var UserRepo UserRepository

type userRepository struct {
}

func init() {
	UserRepo = new(userRepository)
}

func (r *userRepository) UserRegister(user *po_mysql.User) (err error) {
	var (
		db *gorm.DB
	)
	if db, err = mysql.GormDB(); err != nil {
		return
	}
	err = db.Create(user).Error
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
	db, err = mysql.GormDB()
	if err != nil {
		return
	}
	err = db.Where("szk_id=?", szkID).Find(&user).Error
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
