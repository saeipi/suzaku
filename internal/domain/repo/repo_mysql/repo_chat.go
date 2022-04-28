package repo_mysql

import (
	"gorm.io/gorm"
	"suzaku/internal/domain/po_mysql"
	"suzaku/pkg/common/mysql"
)

type ChatRepository interface {
	SaveChatLog(log *po_mysql.ChatLog) (err error)
}

var ChatRepo ChatRepository

type chatRepository struct {
}

func init() {
	ChatRepo = new(chatRepository)
}

func (r *chatRepository) SaveChatLog(log *po_mysql.ChatLog) (err error) {
	var (
		db *gorm.DB
	)
	db, err = mysql.GormDB()
	if err != nil {
		return
	}
	err = db.Save(log).Error
	return
}
