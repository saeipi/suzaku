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
