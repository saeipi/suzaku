package repository_mongo

import (
	"context"
	"github.com/golang/protobuf/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"strconv"
	"suzaku/internal/domain/entity/entity_mongo"
	"suzaku/pkg/common/mongodb"
	pb_chat "suzaku/pkg/proto/chart"
	"suzaku/pkg/proto/pb_ws"
	"suzaku/pkg/utils"
	"time"
)

const singleGocMsgNum = 5000

type MgChatRepository interface {
	SaveUserChatMongo2(uid string, sendTime int64, msg *pb_chat.MsgDataToDB) (err error)
	GetMsgBySeqListMongo2(uid string, seqList []uint32, operationID string) (seqMsg []*pb_ws.MsgData, err error)
}

var MgChatRepo MgChatRepository

type mgChatRepository struct {
}

func init() {
	MgChatRepo = new(mgChatRepository)
}

func (r *mgChatRepository) SaveUserChatMongo2(uid string, sendTime int64, msg *pb_chat.MsgDataToDB) (err error) {
	var (
		db   *mongo.Database
		coll *mongo.Collection
		ctx  context.Context

		//newTime int64
		seqUid  string
		filter  bson.M
		msgInfo entity_mongo.MessageInfo

		chat entity_mongo.UserChat
	)
	ctx, _ = NewContext()
	db, err = mongodb.MgDB()
	if err != nil {
		//TODO:错误
		return
	}
	coll = db.Collection(entity_mongo.MongoCollectionMsg)

	//newTime = getCurrentTimestampByMill()
	seqUid = getSeqUid(uid, msg.MsgData.Seq)
	filter = bson.M{"uid": seqUid}

	msgInfo = entity_mongo.MessageInfo{}
	msgInfo.SendTime = sendTime
	if msgInfo.Msg, err = proto.Marshal(msg.MsgData); err != nil {
		//TODO:错误
		return
	}

	if err = coll.FindOneAndUpdate(ctx, filter, bson.M{"$push": bson.M{"msg": msgInfo}}).Err(); err != nil {
		chat = entity_mongo.UserChat{}
		chat.UID = seqUid
		chat.Msg = append(chat.Msg, msgInfo)

		if _, err = coll.InsertOne(ctx, &chat); err != nil {
			//TODO:错误
			return
		}
	}
	return
}

func (r *mgChatRepository) SaveUserChat(uid string, sendTime int64, msg pb_chat.MsgDataToDB) (err error) {
	return
}

func (r *mgChatRepository) GetMsgBySeqListMongo2(uid string, seqList []uint32, operationID string) (seqMsg []*pb_ws.MsgData, err error) {
	var (
		db   *mongo.Database
		coll *mongo.Collection
		ctx  context.Context

		seqs        map[string][]uint32
		chat        entity_mongo.UserChat
		reqUid      string
		values      []uint32
		singleCount int
		i           int
		msg         *pb_ws.MsgData
		hasSeqList  []uint32
	)
	db, err = mongodb.MgDB()
	if err != nil {
		return
	}
	coll = db.Collection(entity_mongo.MongoCollectionMsg)
	ctx, _ = NewContext()

	seqMsg = make([]*pb_ws.MsgData, 0)
	hasSeqList = make([]uint32, 0)
	seqs = func(uid string, seqList []uint32) (seqs map[string][]uint32) {
		seqs = make(map[string][]uint32)
		for i := 0; i < len(seqList); i++ {
			seqUid := getSeqUid(uid, seqList[i])
			if value, ok := seqs[seqUid]; !ok {
				var temp []uint32
				seqs[seqUid] = append(temp, seqList[i])
			} else {
				seqs[seqUid] = append(value, seqList[i])
			}
		}
		return
	}(uid, seqList)
	chat = entity_mongo.UserChat{}

	for reqUid, values = range seqs {
		if err = coll.FindOne(ctx, bson.M{"uid": reqUid}).Decode(&chat); err != nil {
			//TODO:错误
			continue
		}
		singleCount = 0
		for i = 0; i < len(chat.Msg); i++ {
			msg = new(pb_ws.MsgData)
			if err = proto.Unmarshal(chat.Msg[i].Msg, msg); err != nil {
				//TODO:错误
				return
			}
			if isContainInt32(msg.Seq, values) {
				seqMsg = append(seqMsg, msg)
				hasSeqList = append(hasSeqList, msg.Seq)
				singleCount++
				if singleCount == len(values) {
					break
				}
			}
		}
	}
	if len(hasSeqList) != len(seqList) {
		var diff []uint32
		diff = utils.Difference(hasSeqList, seqList)
		exceptionMSg := genExceptionMessageBySeqList(diff)
		seqMsg = append(seqMsg, exceptionMSg...)
	}
	return
}

func getCurrentTimestampByMill() int64 {
	return time.Now().UnixNano() / 1e6
}
func getSeqUid(uid string, seq uint32) string {
	seqSuffix := seq / singleGocMsgNum
	return indexGen(uid, seqSuffix)
}
func indexGen(uid string, seqSuffix uint32) string {
	return uid + ":" + strconv.FormatInt(int64(seqSuffix), 10)
}
func isContainInt32(target uint32, List []uint32) bool {
	for _, element := range List {
		if target == element {
			return true
		}
	}
	return false
}
func genExceptionMessageBySeqList(seqList []uint32) (exceptionMsg []*pb_ws.MsgData) {
	for _, v := range seqList {
		msg := new(pb_ws.MsgData)
		msg.Seq = v
		exceptionMsg = append(exceptionMsg, msg)
	}
	return exceptionMsg
}
