package repo_mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
	"suzaku/internal/domain/po_mongo"
	"suzaku/pkg/common/mongodb"
	pb_chat "suzaku/pkg/proto/chart"
	"suzaku/pkg/proto/pb_ws"
	"time"
)

const singleGocMsgNum = 5000

type MgChatRepository interface {
	SaveMessage(msg *po_mongo.Message) (err error)
	HistoryMessages(req *pb_chat.GetHistoryMessagesReq) (messages []*po_mongo.Message, err error)
}

var MgChatRepo MgChatRepository

type mgChatRepository struct {
}

func init() {
	MgChatRepo = new(mgChatRepository)
}

func (r *mgChatRepository) SaveMessage(msg *po_mongo.Message) (err error) {
	var (
		db         *mongo.Database
		coll       *mongo.Collection
		ctx        context.Context
		cancelFunc context.CancelFunc
	)
	ctx, cancelFunc = NewContext()
	defer cancelFunc()

	if db, err = mongodb.MgDB(); err != nil {
		//TODO:错误
		return
	}
	coll = db.Collection(po_mongo.MongoCollectionMsg)
	if _, err = coll.InsertOne(ctx, msg); err != nil {
		//TODO:错误
		return
	}
	return
}

func (r *mgChatRepository) HistoryMessages(req *pb_chat.GetHistoryMessagesReq) (messages []*po_mongo.Message, err error) {
	var (
		db          *mongo.Database
		coll        *mongo.Collection
		ctx         context.Context
		cannel      context.CancelFunc
		findoptions *options.FindOptions
		cur         *mongo.Cursor
		filter      map[string]interface{}
		sort        bson.D
	)
	messages = make([]*po_mongo.Message, 0)
	filter = make(map[string]interface{})
	if db, err = mongodb.MgDB(); err != nil {
		return
	}
	coll = db.Collection(po_mongo.MongoCollectionMsg)
	ctx, cannel = NewContext()
	defer cannel()

	findoptions = &options.FindOptions{}
	findoptions.SetLimit(int64(req.PageSize))
	sort = bson.D{
		bson.E{"seq", -1},
	}
	findoptions.SetSort(sort)
	// findoptions.SetSkip(0)

	filter["session_id"] = req.SessionId
	filter["session_type"] = req.SessionType
	if req.Seq > 0 {
		if req.Back == true {
			//小于($lt)
			filter["seq"] = bson.M{"$lt": req.Seq}

		} else {
			//大于($gt)
			filter["seq"] = bson.M{"$gt": req.Seq}
		}
	}

	cur, err = coll.Find(ctx, filter, findoptions)
	if err != nil {
		return
	}
	defer cur.Close(context.Background())
	err = cur.All(context.Background(), &messages)
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
