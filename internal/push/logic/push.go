package logic

import (
	"suzaku/pkg/common/config"
	"suzaku/pkg/common/grpc-etcdv3/getcdv3"
	"suzaku/pkg/constant"
)

var (
	rpcServer    *pushRpcServer
	pushCH       *PushConsumerHandler
	pushTerminal []int32
	//producer     *kafka.Producer
	sendCount uint64
	watcher   *getcdv3.Watcher
)

func Initialize(rpcPort int) {
	//producer = kafka.NewKafkaProducer(config.Config.Kafka.Ws2Mschat.Addr, config.Config.Kafka.Ws2Mschat.Topic)
	var (
		catalog string
		err     error
	)
	catalog = config.Config.Etcd.Schema + ":///" + config.Config.RPCRegisterName.OnlineMessageRelayName
	watcher, err = getcdv3.NewWatcher(catalog, config.Config.Etcd.Schema, config.Config.Etcd.Address)
	if err != nil {
		//TODO:error
		return
	}
	rpcServer = NewPushRpcServer(rpcPort)
	pushCH = NewPushConsumerHandler()
	pushTerminal = []int32{constant.IOSPlatformID, constant.AndroidPlatformID}
}

func Run() {
	go rpcServer.Run()
	go pushCH.pushConsumerGroup.RegisterHandleAndConsumer(pushCH)
	watcher.Run()
}
