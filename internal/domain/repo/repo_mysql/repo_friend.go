package repo_mysql

import (
	"gorm.io/gorm"
	"suzaku/internal/domain/po_mysql"
	"suzaku/pkg/common/mysql"
)

type FriendRepository interface {
	SaveFriendRequest(req *po_mysql.FriendRequest) (err error)
}

var FriendRepo FriendRepository

type friendRepository struct {
}

func init() {
	FriendRepo = new(friendRepository)
}

func (r *friendRepository) SaveFriendRequest(req *po_mysql.FriendRequest) (err error) {
	var (
		db *gorm.DB
	)
	if db, err = mysql.GormDB(); err != nil {
		return
	}
	err = db.Save(req).Error
	return
}
