package repo_mysql

import (
	"gorm.io/gorm"
	"suzaku/internal/domain/po_mysql"
	"suzaku/pkg/common/mysql"
)

type GroupRepository interface {
	Create(group po_mysql.Group) (err error)
	GroupExist(groupID string) (group po_mysql.Group, err error)
	Join(member po_mysql.GroupMember) (err error)
	IsJoin(groupID string, userID string) (member po_mysql.GroupMember, err error)
}

var GroupRepo GroupRepository

type groupRepository struct {
}

func init() {
	GroupRepo = new(groupRepository)
}

func (r *groupRepository) Create(group po_mysql.Group) (err error) {
	var (
		db *gorm.DB
	)
	if db, err = mysql.GormDB(); err != nil {
		return
	}
	err = db.Create(&group).Error
	return
}

func (r *groupRepository) GroupExist(groupID string) (group po_mysql.Group, err error) {
	var (
		db *gorm.DB
	)
	if db, err = mysql.GormDB(); err != nil {
		return
	}
	err = db.Where("group_id = ?", groupID).Find(&group).Error
	if err == gorm.ErrRecordNotFound {
		err = nil
	}
	return
}

func (r *groupRepository) Join(member po_mysql.GroupMember) (err error) {
	var (
		db *gorm.DB
	)
	if db, err = mysql.GormDB(); err != nil {
		return
	}
	err = db.Create(&member).Error
	return
}

func (r *groupRepository) IsJoin(groupID string, userID string) (member po_mysql.GroupMember, err error) {
	var (
		db *gorm.DB
	)
	if db, err = mysql.GormDB(); err != nil {
		return
	}
	err = db.Model(po_mysql.GroupMember{}).Where("group_id=? AND user_id=?", groupID, userID).Find(&member).Error
	return
}
