package logic

import (
	"suzaku/internal/domain/repo/repo_mongo"
	pb_chat "suzaku/pkg/proto/chart"
)

func saveUserChat(uid string, msg *pb_chat.MsgDataToMQ) (err error) {
	var (
		//seq        uint64
		pbSaveData pb_chat.MsgDataToDB
		//nowMsec int64
	)
	//nowMsec = utils.GetCurrentTimestampByMill()
	//seq, err = redis.IncrSeqID(msg.MsgData.SessionId)
	//if err != nil {
	//	return err
	//}
	//msg.MsgData.Seq = uint32(seq)
	pbSaveData = pb_chat.MsgDataToDB{}
	pbSaveData.MsgData = msg.MsgData
	repo_mongo.MgChatRepo.SaveUserChatMongo2(uid, pbSaveData.MsgData.SendTs, &pbSaveData)
	return
}
