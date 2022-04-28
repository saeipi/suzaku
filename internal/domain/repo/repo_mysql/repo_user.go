package repo_mysql

import (
	"gorm.io/gorm"
	"suzaku/internal/domain/po_mysql"
	"suzaku/pkg/common/mysql"
)

type UserRepository interface {
	UserRegister(user *po_mysql.User) (err error)
	GetUserByUserID(userID string) (user *po_mysql.User, err error)
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
	if db, err = mysql.GormDB();err != nil{
		return
	}
	err = db.Save(user).Error
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
