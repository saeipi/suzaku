package repo_mysql

import (
	"gorm.io/gorm"
	"suzaku/internal/domain/do"
	"suzaku/internal/domain/po_mysql"
	"suzaku/pkg/common/mysql"
	pb_friend "suzaku/pkg/proto/friend"
	"suzaku/pkg/utils"
)

type FriendRepository interface {
	SaveFriendRequest(req *po_mysql.FriendRequest) (err error)
	GetFriendRequestList(query *do.MysqlQuery) (list []*po_mysql.FriendRequest, totalRows int64, err error)
	IsFriend(userId1, userId2 string) (friend *po_mysql.Friend, err error)

	UpdateFriendRequest(req *pb_friend.HandleFriendRequestReq) (err error)
	ApproveFriendRequest(req *pb_friend.HandleFriendRequestReq) (err error)
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

func (r *friendRepository) IsFriend(userId1, userId2 string) (friend *po_mysql.Friend, err error) {
	var (
		db *gorm.DB
	)
	if db, err = mysql.GormDB(); err != nil {
		return
	}
	err = db.Where("owner_user_id=? AND friend_user_id=?", userId1, userId2).Find(&friend).Error
	if err == gorm.ErrRecordNotFound {
		err = nil
	}
	return
}

func (r *friendRepository) TxUpdateFriendRequest(req *pb_friend.HandleFriendRequestReq, tx *gorm.DB) (err error) {
	var (
		friendRequest po_mysql.FriendRequest
		updates       map[string]interface{}
	)
	updates = make(map[string]interface{})
	updates["handle_user_id"] = req.UserId
	updates["handle_result"] = req.HandleResult
	updates["handle_msg"] = req.HandleMsg
	updates["handled_ts"] = utils.GetCurrentTimestampByMill()
	if req.HandleUserId != "" {
		updates["handle_user_id"] = req.HandleUserId
	} else {
		updates["handle_user_id"] = req.UserId
	}
	err = tx.Where("from_user_id=? AND to_user_id=?", req.FromUserId, req.UserId).Find(&friendRequest).Updates(updates).Error
	return
}

func (r *friendRepository) UpdateFriendRequest(req *pb_friend.HandleFriendRequestReq) (err error) {
	err = mysql.Transaction(func(tx *gorm.DB) (terr error) {
		terr = r.TxUpdateFriendRequest(req, tx)
		return
	})
	return
}

func (r *friendRepository) ApproveFriendRequest(req *pb_friend.HandleFriendRequestReq) (err error) {
	err = mysql.Transaction(func(tx *gorm.DB) (terr error) {
		var (
			friend po_mysql.Friend
		)
		terr = r.TxUpdateFriendRequest(req, tx)
		if terr != nil {
			return
		}
		friend = po_mysql.Friend{
			OwnerUserId:    req.FromUserId,
			FriendUserId:   "",
			OperatorUserId: "",
			Source:         0, // 暂时全部为0
			Remark:         "",
			Ex:             "",
		}
		if req.HandleUserId != "" {
			friend.OperatorUserId = req.HandleUserId
		} else {
			friend.OperatorUserId = req.UserId
		}
		terr = tx.Create(&friend).Error
		return
	})
	return
}
