package logic

import (
	"github.com/Shopify/sarama"
	"github.com/golang/protobuf/proto"
	"suzaku/pkg/common/config"
	"suzaku/pkg/common/kafka"
	pb_chat "suzaku/pkg/proto/chart"
	pb_push "suzaku/pkg/proto/push"
)

type MessageHandler func(msg []byte)

type PushConsumerHandler struct {
	msgHandle         map[string]MessageHandler
	pushConsumerGroup *kafka.MConsumerGroup
}

func NewPushConsumerHandler() (handler *PushConsumerHandler) {
	handler = &PushConsumerHandler{msgHandle: make(map[string]MessageHandler)}
	handler.msgHandle[config.Config.Kafka.Ms2Pschat.Topic] = handler.handleMs2PsChat
	handler.pushConsumerGroup = kafka.NewMConsumerGroup(&kafka.MConsumerGroupConfig{KafkaVersion: sarama.V0_10_2_0,
		OffsetsInitial: sarama.OffsetNewest, IsReturnErr: false},
		[]string{config.Config.Kafka.Ms2Pschat.Topic}, config.Config.Kafka.Ms2Pschat.Addr,
		config.Config.Kafka.ConsumerGroupID.MsgToPush)
	return
}

func (h *PushConsumerHandler) handleMs2PsChat(msg []byte) {
	var (
		msgFromMQ pb_chat.PushMsgDataToMQ
		pushMsg   pb_push.PushMsgReq
	)
	msgFromMQ = pb_chat.PushMsgDataToMQ{}
	if err := proto.Unmarshal(msg, &msgFromMQ); err != nil {
		//TODO:错误
		return
	}
	pushMsg = pb_push.PushMsgReq{
		OperationId:  msgFromMQ.OperationId,
		MsgData:      msgFromMQ.MsgData,
		PushToUserId: msgFromMQ.PushToUserId,
	}
	MsgToUser(&pushMsg)
}

func (PushConsumerHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (PushConsumerHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h *PushConsumerHandler) ConsumeClaim(sess sarama.ConsumerGroupSession,
	claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		//log.InfoByKv("kafka get info to mysql", "", "msgTopic", msg.Topic, "msgPartition", msg.Partition, "msg", string(msg.Value))
		h.msgHandle[msg.Topic](msg.Value)
	}
	return nil
}
