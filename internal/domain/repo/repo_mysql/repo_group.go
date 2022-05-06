package repo_mysql

import (
	"gorm.io/gorm"
	"suzaku/internal/domain/po_mysql"
	"suzaku/pkg/common/mysql"
)

type GroupRepository interface {
	Create(group *po_mysql.Group) (err error)
	GroupExist(groupID string) (res *po_mysql.Group, err error)
	RequestJoin(request *po_mysql.GroupRequest) (err error)
	Join(member *po_mysql.GroupMember) (err error)
	IsJoined(groupID string, userID string) (member *po_mysql.GroupMember, err error)
}

var GroupRepo GroupRepository

type groupRepository struct {
}

func init() {
	GroupRepo = new(groupRepository)
}
/*
存:传指针对象，Create时不需要&，同时会Out表中的数据
读:返回指针对象，Find时需要&
 */
func (r *groupRepository) Create(group *po_mysql.Group) (err error) {
	var (
		db *gorm.DB
	)
	if db, err = mysql.GormDB(); err != nil {
		return
	}
	err = db.Create(group).Error
	return
}

func (r *groupRepository) GroupExist(groupID string) (group *po_mysql.Group, err error) {
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

func (r *groupRepository) RequestJoin(request *po_mysql.GroupRequest) (err error) {
	var (
		db *gorm.DB
	)
	if db, err = mysql.GormDB(); err != nil {
		return
	}
	err = db.Create(request).Error
	return
}

func (r *groupRepository) Join(member *po_mysql.GroupMember) (err error) {
	var (
		db *gorm.DB
	)
	if db, err = mysql.GormDB(); err != nil {
		return
	}
	err = db.Create(member).Error
	return
}

func (r *groupRepository) IsJoined(groupID string, userID string) (member *po_mysql.GroupMember, err error) {
	var (
		db *gorm.DB
	)
	if db, err = mysql.GormDB(); err != nil {
		return
	}
	err = db.Model(po_mysql.GroupMember{}).Where("group_id=? AND user_id=?", groupID, userID).Find(&member).Error
	if err == gorm.ErrRecordNotFound {
		err = nil
	}
	return
}
