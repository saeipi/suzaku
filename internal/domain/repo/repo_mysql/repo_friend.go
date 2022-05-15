package repo_mysql

import (
	"gorm.io/gorm"
	"suzaku/internal/domain/do"
	"suzaku/internal/domain/po_mysql"
	"suzaku/pkg/common/mysql"
	"suzaku/pkg/common/snowflake"
	pb_friend "suzaku/pkg/proto/friend"
	"suzaku/pkg/utils"
)

type FriendRepository interface {
	SaveFriendRequest(req *po_mysql.FriendRequest) (err error)
	GetFriendRequestList(query *do.MysqlQuery) (list []*po_mysql.FriendRequest, totalRows int64, err error)
	IsFriend(userId1, userId2 string) (friend *po_mysql.Friend, err error)

	UpdateFriendRequest(req *pb_friend.HandleFriendRequestReq) (err error)
	ApproveFriendRequest(req *pb_friend.HandleFriendRequestReq) (err error)

	FriendList(req *pb_friend.FriendListReq) (friends []*do.FriendInfo, totalRows int64, err error)
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
	req.ReqId = snowflake.SnowflakeID()
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
	friend = new(po_mysql.Friend)
	if db, err = mysql.GormDB(); err != nil {
		return
	}
	err = db.Where("owner_user_id=? AND friend_user_id=?", userId1, userId2).Find(friend).Error
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
	updates["handle_result"] = req.HandleResult
	updates["handle_msg"] = req.HandleMsg
	updates["handled_ts"] = utils.GetCurrentTimestampByMill()
	if req.HandleUserId != "" {
		updates["handle_user_id"] = req.HandleUserId
	} else {
		updates["handle_user_id"] = req.UserId
	}
	err = tx.Where("req_id=?", req.ReqId).Find(&friendRequest).Updates(updates).Error
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
			OwnerUserId:    req.UserId,
			FriendUserId:   req.FromUserId,
			OperatorUserId: "",
			SessionId:      utils.GetSessionId(req.FromUserId, req.UserId),
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
		if terr != nil {
			return
		}
		friend.OwnerUserId = req.FromUserId
		friend.FriendUserId = req.UserId
		terr = tx.Create(&friend).Error
		return
	})
	return
}

func (r *friendRepository) FriendList(req *pb_friend.FriendListReq) (friends []*do.FriendInfo, totalRows int64, err error) {
	var (
		db *gorm.DB
	)
	friends = make([]*do.FriendInfo, 0)
	if db, err = mysql.GormDB(); err != nil {
		return
	}
	err = db.Table("friends").
		Select("friends.session_id,users.*").
		Joins("LEFT JOIN users ON users.user_id=friends.owner_user_id").
		Where("owner_user_id=?", req.UserId).
		Count(&totalRows).
		Limit(int(req.PageSize)).
		Offset(int((req.Page - 1) * req.PageSize)).
		Find(&friends).Error

	/*
		err = db.Table("(? UNION ALL ?) tb",
			db.Table("friends").Select("friends.session_id,users.*").Joins("LEFT JOIN users ON users.user_id=friends.owner_user_id").Where("owner_user_id=?", req.UserId),
			db.Table("friends").Select("friends.session_id,users.*").Joins("LEFT JOIN users ON users.user_id=friends.friend_user_id").Where("friend_user_id=?", req.UserId)).
			Select("*").
			Count(&totalRows).
			Limit(int(req.PageSize)).
			Offset(int((req.Page - 1) * req.PageSize)).
			Find(&friends).Error
	*/
	return
}
