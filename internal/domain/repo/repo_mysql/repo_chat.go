package repo_mysql

import (
	"gorm.io/gorm"
	"suzaku/internal/domain/po_mysql"
	"suzaku/pkg/common/mysql"
	"suzaku/pkg/common/redis"
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
		db  *gorm.DB
		seq uint64
	)
	if db, err = mysql.GormDB(); err != nil {
		return
	}
	seq, err = redis.IncrSeqID(msg.SessionId)
	if err != nil {
		return
	}
	msg.Seq = int64(seq)
	err = db.Create(msg).Error
	return
}

func (r *chatRepository) HistoryMessages(req *pb_chat.GetHistoryMessagesReq) (messages []*po_mysql.Message, err error) {
	var (
		db         *gorm.DB
		query      string
		orderValue string
		args       []interface{}
	)
	messages = make([]*po_mysql.Message, 0)
	if db, err = mysql.GormDB(); err != nil {
		return
	}
	query = "session_type=? AND session_id=?"
	args = []interface{}{req.SessionType, req.SessionId}

	if req.Seq > 0 {
		if req.Back == true {
			query += " AND seq_id<?"

		} else {
			query += " AND seq_id>?"
		}
		args = append(args, req.Seq)
		orderValue = "seq_id ASC"
	} else {
		orderValue = "seq_id DESC"
	}
	err = db.Where(query, args...).
		Order(orderValue).
		Limit(int(req.PageSize)).
		Find(&messages).Error
	return
}
