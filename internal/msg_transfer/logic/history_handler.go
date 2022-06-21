package logic

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/golang/protobuf/proto"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc"
	"suzaku/internal/domain/po_mongo"
	"suzaku/internal/domain/repo/repo_mongo"
	"suzaku/pkg/common/config"
	"suzaku/pkg/common/kafka"
	"suzaku/pkg/constant"
	"suzaku/pkg/factory"
	pb_chat "suzaku/pkg/proto/chart"
	pb_push "suzaku/pkg/proto/push"
	"suzaku/pkg/utils"
)

type MessageHandler func(msg []byte, msgKey string)

type HistoryConsumerHandler struct {
	msgHandle            map[string]MessageHandler
	historyConsumerGroup *kafka.MConsumerGroup
	singleMsgCount       uint64
	groupMsgCount        uint64
}

func NewHistoryConsumerHandler() (handler *HistoryConsumerHandler) {
	handler = &HistoryConsumerHandler{msgHandle: make(map[string]MessageHandler)}
	handler.msgHandle[config.Config.Kafka.Ws2Mschat.Topic] = handler.MessageHandler
	handler.historyConsumerGroup = kafka.NewMConsumerGroup(&kafka.MConsumerGroupConfig{KafkaVersion: sarama.V3_1_0_0,
		OffsetsInitial: sarama.OffsetNewest, IsReturnErr: false}, []string{config.Config.Kafka.Ws2Mschat.Topic},
		config.Config.Kafka.Ws2Mschat.Addr, config.Config.Kafka.ConsumerGroupID.MsgToMongo)
	return
}

func (h *HistoryConsumerHandler) MessageHandler(msg []byte, msgKey string) {
	var (
		//nowNano      int64
		msgFromMQ pb_chat.MsgDataToMQ
		err       error
		isHistory bool
		//isPersist    bool
		isSenderSync bool
	)
	//nowNano = utils.GetCurrentTimestampByNano()
	msgFromMQ = pb_chat.MsgDataToMQ{}
	err = proto.Unmarshal(msg, &msgFromMQ)
	if err != nil {
		//TODO:错误
		return
	}
	//Control whether to store offline messages (mongo)
	isHistory = utils.GetSwitchFromOptions(msgFromMQ.MsgData.Options, constant.IsHistory)
	//Control whether to store history messages (mysql)
	//isPersist = utils.GetSwitchFromOptions(msgFromMQ.MsgData.Options, constant.IsPersistent)
	isSenderSync = utils.GetSwitchFromOptions(msgFromMQ.MsgData.Options, constant.IsSenderSync)
	switch msgFromMQ.MsgData.SessionType {
	case constant.SingleChatType:
		if isHistory {
			err = saveMessage(&msgFromMQ)
			if err != nil {
				//TODO:错误
				return
			}
			h.singleMsgCount++
		}
		if !isSenderSync && msgKey == msgFromMQ.MsgData.SendId {
		} else {
			go sendMessageToPush(&msgFromMQ, msgKey)
		}
	case constant.GroupChatType:
		if isHistory {
			err = saveMessage(&msgFromMQ)
			if err != nil {
				//TODO:错误
				return
			}
			h.groupMsgCount++
		}
		go sendMessageToPush(&msgFromMQ, msgFromMQ.MsgData.RecvId)
	}
	return
}

func saveMessage(msg *pb_chat.MsgDataToMQ) (err error) {
	var (
		message *po_mongo.Message
	)
	message = new(po_mongo.Message)
	copier.Copy(message, msg.MsgData)
	repo_mongo.MgChatRepo.SaveMessage(message)
	return
}

func sendMessageToPush(message *pb_chat.MsgDataToMQ, pushToUserID string) {
	var (
		rpcPushMsg pb_push.PushMsgReq
		mqPushMsg  pb_chat.PushMsgDataToMQ
		clientConn *grpc.ClientConn

		pid    int32
		offset int64
		err    error

		client pb_push.PushMsgServiceClient
	)
	rpcPushMsg = pb_push.PushMsgReq{
		OperationId:  message.OperationId,
		MsgData:      message.MsgData,
		PushToUserId: pushToUserID,
	}
	mqPushMsg = pb_chat.PushMsgDataToMQ{
		OperationId:  message.OperationId,
		MsgData:      message.MsgData,
		PushToUserId: pushToUserID,
	}

	clientConn = factory.ClientConn(config.Config.RPCRegisterName.PushName)
	if clientConn == nil {
		// 消息再次放入kafka消息队列
		pid, offset, err = producer.SendMessage(&mqPushMsg)
		if err != nil {
			//TODO:错误
			fmt.Println("kafka send failed", mqPushMsg.OperationId, "send data", message.String(), "pid", pid, "offset", offset, "err", err.Error())
			return
		}
		return
	}
	client = pb_push.NewPushMsgServiceClient(clientConn)
	_, err = client.PushMsg(context.Background(), &rpcPushMsg)
	if err != nil {
		// 消息再次放入kafka消息队列
		pid, offset, err = producer.SendMessage(&mqPushMsg)
		if err != nil {
			//TODO:错误
			fmt.Println("kafka send failed", mqPushMsg.OperationId, "send data", mqPushMsg.String(), "pid", pid, "offset", offset, "err", err.Error())
		}
		return
	}
}

func (HistoryConsumerHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (HistoryConsumerHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h *HistoryConsumerHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		//log.InfoByKv("kafka get info to mongo", "", "msgTopic", msg.Topic, "msgPartition", msg.Partition, "msg", string(msg.Value))
		h.msgHandle[msg.Topic](msg.Value, string(msg.Key))
		sess.MarkMessage(msg, "")
	}
	return nil
}
