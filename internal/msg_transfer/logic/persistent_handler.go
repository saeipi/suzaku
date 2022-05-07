package logic

import (
	"github.com/Shopify/sarama"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/jinzhu/copier"
	"suzaku/internal/domain/po_mysql"
	"suzaku/internal/domain/repo/repo_mysql"
	"suzaku/pkg/common/config"
	"suzaku/pkg/common/kafka"
	"suzaku/pkg/constant"
	pb_chat "suzaku/pkg/proto/chart"
	"suzaku/pkg/proto/pb_ws"
	"suzaku/pkg/utils"
)

type PersistentConsumerHandler struct {
	msgHandle               map[string]MessageHandler
	persistentConsumerGroup *kafka.MConsumerGroup
}

func NewPersistentConsumerHandler() (handler *PersistentConsumerHandler) {
	handler = &PersistentConsumerHandler{msgHandle: make(map[string]MessageHandler)}
	handler.msgHandle[config.Config.Kafka.Ws2Mschat.Topic] = handler.MessageHandler
	handler.persistentConsumerGroup = kafka.NewMConsumerGroup(&kafka.MConsumerGroupConfig{KafkaVersion: sarama.V0_10_2_0,
		OffsetsInitial: sarama.OffsetNewest, IsReturnErr: false}, []string{config.Config.Kafka.Ws2Mschat.Topic},
		config.Config.Kafka.Ws2Mschat.Addr, config.Config.Kafka.ConsumerGroupID.MsgToMySQL)
	return
}

func (h *PersistentConsumerHandler) MessageHandler(msg []byte, msgKey string) {
	var (
		msgFromMQ pb_chat.MsgDataToMQ
		err       error
		isPersist bool
	)
	msgFromMQ = pb_chat.MsgDataToMQ{}
	err = proto.Unmarshal(msg, &msgFromMQ)
	if err != nil {
		//TODO:错误
		return
	}
	//Control whether to store history messages (mysql)
	isPersist = utils.GetSwitchFromOptions(msgFromMQ.MsgData.Options, constant.IsPersistent)
	if isPersist == true {
		if msgFromMQ.MsgData.SessionType == constant.SingleChatType && msgKey == msgFromMQ.MsgData.RecvId {
			if err = h.InsertMessageToChatLog(msgFromMQ); err != nil {
				//TODO:错误
				return
			}
		} else if msgFromMQ.MsgData.SessionType == constant.GroupChatType && msgKey == msgFromMQ.MsgData.SendId {
			if err = h.InsertMessageToChatLog(msgFromMQ); err != nil {
				//TODO:错误
				return
			}
		}
	}
}

func (h *PersistentConsumerHandler) InsertMessageToChatLog(msg pb_chat.MsgDataToMQ) (err error) {
	var (
		message *po_mysql.Message
		tips    pb_ws.TipsComm
	)
	message = new(po_mysql.Message)
	copier.Copy(message, msg.MsgData)
	switch msg.MsgData.SessionType {
	case constant.GroupChatType:
		message.RecvId = msg.MsgData.GroupId
	case constant.SingleChatType:
		message.RecvId = msg.MsgData.RecvId
	}
	if msg.MsgData.ContentType >= constant.NotificationBegin && msg.MsgData.ContentType <= constant.NotificationEnd {
		_ = proto.Unmarshal(msg.MsgData.Content, &tips)
		marshaler := jsonpb.Marshaler{
			OrigName:     true,
			EnumsAsInts:  false,
			EmitDefaults: false,
		}
		message.Content, _ = marshaler.MarshalToString(&tips)
	} else {
		message.Content = string(msg.MsgData.Content)
	}
	message.SendTs = msg.MsgData.SendTs
	err = repo_mysql.ChatRepo.SaveMessage(message)
	if err != nil {
		//TODO:错误
	}
	return
}

func (PersistentConsumerHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (PersistentConsumerHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h *PersistentConsumerHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		//log.InfoByKv("kafka get info to mysql", "", "msgTopic", msg.Topic, "msgPartition", msg.Partition, "msg", string(msg.Value))
		h.msgHandle[msg.Topic](msg.Value, string(msg.Key))
		sess.MarkMessage(msg, "")
	}
	return nil
}
