package repo_mysql

import (
	"gorm.io/gorm"
	"suzaku/internal/domain/po_mysql"
	"suzaku/pkg/common/mysql"
	pb_chat "suzaku/pkg/proto/chart"
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

	if req.Seq > 0 {
		if req.Back == true {
			query += " AND seq<?"

		} else {
			query += " AND seq>?"
		}
		args = append(args, req.Seq)
	}
	err = db.Where(query, args...).
		Order("seq DESC").
		Limit(int(req.PageSize)).
		Find(&messages).Error
	return
}
