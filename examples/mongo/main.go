package main

import (
	"suzaku/internal/domain/repo/repo_mongo"
	pb_chat "suzaku/pkg/proto/chart"
)

func main() {
	req := pb_chat.GetHistoryMessagesReq{
		PageSize:    20,
		Seq:         32,
		CreatedTs:   0,
		SessionId:   "a168f11853228b0748f0b4dac4658bb6",
		SessionType: 1,
		Back:        true,
	}
	msgs, err := repo_mongo.MgChatRepo.HistoryMessages(&req)

	if err != nil {
		return
	}
	if len(msgs) == 0 {

	}
}
