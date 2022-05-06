package repo_mysql

import (
	"gorm.io/gorm"
	"suzaku/internal/domain/do"
	"suzaku/internal/domain/po_mysql"
	"suzaku/pkg/common/mysql"
)

type FriendRepository interface {
	SaveFriendRequest(req *po_mysql.FriendRequest) (err error)
	GetFriendRequestList(query *do.MysqlQuery) (list []*po_mysql.FriendRequest, totalRows int64, err error)
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
	err = db.Create(req).Error
	return
}

func (r *friendRepository) GetFriendRequestList(query *do.MysqlQuery) (list []*po_mysql.FriendRequest, totalRows int64, err error) {
	var (
		db *gorm.DB
	)
	list = make([]*po_mysql.FriendRequest, 0)
	if db, err = mysql.GormDB(); err != nil {
		return
	}
	err = db.Where(query.Condition, query.Params...).
		Find(&list).
		Count(&totalRows).
		Offset(query.Page).
		Limit((query.Page - 1) * query.PageSize).Error
	return
}
