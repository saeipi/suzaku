package logic

import (
	"github.com/jinzhu/copier"
	"suzaku/internal/domain/po_mongo"
	"suzaku/internal/domain/repo/repo_mongo"
	pb_chat "suzaku/pkg/proto/chart"
)

func saveMessage(msg *pb_chat.MsgDataToMQ) (err error) {
	var (
		message *po_mongo.Message
	)
	message = new(po_mongo.Message)
	copier.Copy(message, msg.MsgData)
	repo_mongo.MgChatRepo.SaveMessage(message)
	return
}
