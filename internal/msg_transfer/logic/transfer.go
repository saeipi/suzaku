package logic

import (
	"suzaku/pkg/common/config"
	"suzaku/pkg/common/kafka"
)

var (
	persistentCH *PersistentConsumerHandler
	historyCH    *HistoryConsumerHandler
	producer     *kafka.Producer
)

func Initialize() {
	persistentCH = NewPersistentConsumerHandler()
	historyCH = NewHistoryConsumerHandler()
	producer = kafka.NewKafkaProducer(config.Config.Kafka.Ms2Pschat.Addr, config.Config.Kafka.Ms2Pschat.Topic)
}

func Run() {
	go persistentCH.persistentConsumerGroup.RegisterHandleAndConsumer(persistentCH)
	go historyCH.historyConsumerGroup.RegisterHandleAndConsumer(historyCH)
}
