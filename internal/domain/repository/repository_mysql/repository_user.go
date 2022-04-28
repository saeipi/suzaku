package repository_mysql

import (
	"gorm.io/gorm"
	"suzaku/internal/domain/entity/entity_mysql"
	"suzaku/pkg/common/mysql"
)

type UserRepository interface {
	UserRegister(user *entity_mysql.User) (err error)
}

var UserRepo UserRepository

type userRepository struct {
}

func init() {
	UserRepo = new(userRepository)
}

func (r *userRepository) UserRegister(user *entity_mysql.User) (err error) {
	var (
		db *gorm.DB
	)
	db, err = mysql.GormDB()
	if err != nil {
		return
	}
	err = db.Save(user).Error
	return
}

func (r *userRepository) GetUserByUserID(userID string) (err error) {
	var (
		db *gorm.DB
	)
	db, err = mysql.GormDB()
	if err != nil {
		return
	}
	//err = db.Model(&entity_mysql.User{}).Where("user_id=?",userID).Find()
	return
}
