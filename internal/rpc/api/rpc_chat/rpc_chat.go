package rpc_chat

import (
	"google.golang.org/grpc"
	"suzaku/internal/rpc/rpc_category"
	"suzaku/pkg/common/config"
	"suzaku/pkg/common/kafka"
	pb_chat "suzaku/pkg/proto/chart"
)

type chatRpc struct {
	pb_chat.UnimplementedChatServer
	rpc_category.Rpc
	producer *kafka.Producer
}

func NewRpcChatServer(port int) *chatRpc {
	var (
		rpc *chatRpc
	)
	rpc = &chatRpc{
		Rpc: rpc_category.NewRpcServer(port, config.Config.RPCRegisterName.OfflineMessageName),
	}
	rpc.producer = kafka.NewKafkaProducer(config.Config.Kafka.Ws2Mschat.Addr, config.Config.Kafka.Ws2Mschat.Topic)
	return rpc
}

func (rpc *chatRpc) Run() {
	var (
		server *grpc.Server
	)
	server = grpc.NewServer()
	pb_chat.RegisterChatServer(server, rpc)
	rpc.Rpc.RunServer(server)
}
