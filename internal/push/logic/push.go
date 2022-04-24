package logic

import (
	"suzaku/pkg/constant"
)

var (
	rpcServer    *pushRpcServer
	pushCH       *PushConsumerHandler
	pushTerminal []int32
	//producer     *kafka.Producer
	sendCount uint64
)

func Initialize(rpcPort int) {
	//producer = kafka.NewKafkaProducer(config.Config.Kafka.Ws2Mschat.Addr, config.Config.Kafka.Ws2Mschat.Topic)
	rpcServer = NewPushRpcServer(rpcPort)
	pushCH = NewPushConsumerHandler()
	pushTerminal = []int32{constant.IOSPlatformID, constant.AndroidPlatformID}
}

func Run() {
	go rpcServer.Run()
	go pushCH.pushConsumerGroup.RegisterHandleAndConsumer(pushCH)
}
