package repo_mysql

import (
	"gorm.io/gorm"
	"suzaku/internal/domain/po_mysql"
	"suzaku/pkg/common/mysql"
)

type AuthRepository interface {
	VerifyPassword(userId string, password string) (register *po_mysql.Register, err error)
}

var AuthRepo AuthRepository

type authRepository struct {
}

func init() {
	AuthRepo = new(authRepository)
}

func (r *authRepository) VerifyPassword(userId string, password string) (register *po_mysql.Register, err error) {
	var (
		db *gorm.DB
	)
	register = new(po_mysql.Register)
	db, err = mysql.GormDB()
	if err != nil {
		return
	}
	err = db.Model(po_mysql.Register{}).Where("user_id=? AND password=?", userId, password).Find(register).Error
	if err == gorm.ErrRecordNotFound {
		err = nil
	}
	return
}
