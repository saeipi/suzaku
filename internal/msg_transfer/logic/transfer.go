package logic

import (
	"suzaku/pkg/common/config"
	"suzaku/pkg/common/kafka"
)

var (
	historyCH *HistoryConsumerHandler
	producer *kafka.Producer
)

func NewTransfer() {
	historyCH = NewHistoryConsumerHandler()
	producer = kafka.NewKafkaProducer(config.Config.Kafka.Ms2Pschat.Addr,config.Config.Kafka.Ms2Pschat.Topic)
}