package repo_mysql

import (
	"gorm.io/gorm"
	"suzaku/internal/domain/po_mysql"
	"suzaku/pkg/common/mysql"
	pb_chat "suzaku/pkg/proto/chart"
	"suzaku/pkg/utils"
)

type ChatRepository interface {
	SaveMessage(msg *po_mysql.Message) (err error)
	HistoryMessages(req *pb_chat.GetHistoryMessagesReq) (messages []*po_mysql.Message, err error)
}

var ChatRepo ChatRepository

type chatRepository struct {
}

func init() {
	ChatRepo = new(chatRepository)
}

func (r *chatRepository) SaveMessage(msg *po_mysql.Message) (err error) {
	var (
		db *gorm.DB
	)
	if db, err = mysql.GormDB(); err != nil {
		return
	}
	err = db.Create(msg).Error
	return
}

func (r *chatRepository) HistoryMessages(req *pb_chat.GetHistoryMessagesReq) (messages []*po_mysql.Message, err error) {
	var (
		db    *gorm.DB
		query string
		args  []interface{}
	)
	messages = make([]*po_mysql.Message, 0)
	if db, err = mysql.GormDB(); err != nil {
		return
	}
	query = "session_type=? AND session_id=?"
	args = []interface{}{req.SessionType, req.SessionId}

	if req.CreatedTs == 0 {
		if req.Back == true {
			query += " AND created_ts<?"
			args = append(args, utils.GetCurrentTimestampByMill())
		}
	} else {
		if req.Back == true {
			query += " AND created_ts<?"

		} else {
			query += " AND created_ts>?"
		}
		args = append(args, req.CreatedTs)
	}
	err = db.Where(query, args...).Find(&messages).Error
	return
}
