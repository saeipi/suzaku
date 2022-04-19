package facade_user

import (
	"gorm.io/gorm"
	"suzaku/pkg/common/mysql"
	"suzaku/pkg/common/repository/mysql_model"
)

func UserRegister(user *mysql_model.User) (err error) {
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
