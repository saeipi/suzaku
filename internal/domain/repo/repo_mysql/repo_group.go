package repo_mysql

import (
	"gorm.io/gorm"
	"suzaku/internal/domain/do"
	"suzaku/internal/domain/po_mysql"
	"suzaku/pkg/common/mysql"
	pb_group "suzaku/pkg/proto/group"
	"time"
)

type GroupRepository interface {
	Create(group *po_mysql.Group, avatar *po_mysql.GroupAvatar) (err error)
	GroupExist(groupID string) (res *po_mysql.Group, err error)
	RequestJoin(request *po_mysql.GroupRequest) (err error)
	HandleRequestJoin(req *pb_group.HandleRequestJoinGroupReq) (result *do.JoinGroupResult, err error)
	Join(member *po_mysql.GroupMember) (err error)
	IsJoined(groupID string, userID string) (member *po_mysql.GroupMember, err error)
	AllMember(groupId string) (members []*po_mysql.GroupMember, err error)
	MemberList(req *pb_group.GetGroupMemberListReq) (members []*po_mysql.GroupMember, totalRows int64, err error)
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
func (r *groupRepository) Create(group *po_mysql.Group, avatar *po_mysql.GroupAvatar) (err error) {
	err = mysql.Transaction(func(tx *gorm.DB) (terr error) {
		group.CreatedTs = time.Now().Unix()
		terr = tx.Save(group).Error
		if terr != nil {
			return
		}
		avatar.GroupId = group.GroupId
		avatar.UpdatedTs = group.CreatedTs
		terr = tx.Save(avatar).Error
		return
	})
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
	request.ReqTs = time.Now().Unix()
	err = db.Create(request).Error
	return
}

func (r *groupRepository) HandleRequestJoin(req *pb_group.HandleRequestJoinGroupReq) (result *do.JoinGroupResult, err error) {
	result = new(do.JoinGroupResult)
	result.HandleResult = req.HandleResult

	err = mysql.Transaction(func(tx *gorm.DB) (terr error) {
		var (
			vals         map[string]interface{}
			joinTs       int64
			groupRequest *po_mysql.GroupRequest
			user         *po_mysql.User
			member       *po_mysql.GroupMember
		)
		joinTs = time.Now().Unix()
		// 1、更新GroupRequest
		vals = make(map[string]interface{})
		vals["handle_user_id"] = req.OperationId
		vals["handle_result"] = req.HandleResult
		vals["handle_msg"] = req.HandleMsg
		vals["handle_ts"] = joinTs
		terr = tx.Model(po_mysql.GroupRequest{}).Where("group_id = ? AND user_id = ?", req.GroupId, req.UserId).Updates(vals).Error
		if terr != nil {
			return
		}
		if req.HandleResult != int32(pb_group.HANDLE_JOIN_GROUP_RESULT_APPROVE) {
			return
		}
		// 2、获取最新GroupRequest
		groupRequest, terr = r.TxGetGroupRequest(req.GroupId, req.UserId, tx)
		if terr != nil {
			return
		}
		// 3、获取User
		user, terr = UserRepo.TxGetUserByUserID(groupRequest.UserId, tx)
		if terr != nil {
			return
		}
		// 4、成为群成员
		member = &po_mysql.GroupMember{
			GroupId:        groupRequest.GroupId,
			UserId:         groupRequest.UserId,
			Nickname:       user.Nickname,
			UserAvatarUrl:  user.AvatarUrl,
			JoinedTs:       joinTs,
			JoinSource:     groupRequest.ReqSource,
			OperatorUserId: groupRequest.HandleUserId,
		}
		terr = r.TxCreateGroupMember(member, tx)
		if terr != nil {
			return
		}
		// 5、获取群信息
		result.Member = member
		result.GroupRequest = groupRequest
		result.Group, terr = r.TxGetGroup(groupRequest.GroupId, tx)
		return
	})
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

func (r *groupRepository) TxCreateGroupMember(member *po_mysql.GroupMember, tx *gorm.DB) (err error) {
	err = tx.Create(member).Error
	return
}

func (r *groupRepository) TxGetGroupRequest(groupId, userId string, tx *gorm.DB) (groupRequest *po_mysql.GroupRequest, err error) {
	err = tx.Where("group_id = ? AND user_id = ?", groupId, userId).Find(&groupRequest).Error
	return
}

func (r *groupRepository) TxGetGroup(groupId string, tx *gorm.DB) (group *po_mysql.Group, err error) {
	err = tx.Where("group_id=?", groupId).Find(&group).Error
	return
}

func (r *groupRepository) AllMember(groupId string) (members []*po_mysql.GroupMember, err error) {
	members = make([]*po_mysql.GroupMember, 0)
	var (
		db *gorm.DB
	)
	if db, err = mysql.GormDB(); err != nil {
		return
	}
	err = db.Where("group_id = ?", groupId).Find(&members).Error
	return
}

func (r *groupRepository) MemberList(req *pb_group.GetGroupMemberListReq) (members []*po_mysql.GroupMember, totalRows int64, err error) {
	members = make([]*po_mysql.GroupMember, 0)
	var (
		db *gorm.DB
	)
	if db, err = mysql.GormDB(); err != nil {
		return
	}
	err = db.Model(po_mysql.GroupMember{}).
		Where("group_id = ?", req.GroupId).
		Count(&totalRows).
		Limit(int(req.PageSize)).
		Offset(int((req.Page - 1) * req.PageSize)).
		Find(&members).Error
	return
}
