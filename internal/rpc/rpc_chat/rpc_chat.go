package rpc_chat

import (
	"google.golang.org/grpc"
	"suzaku/internal/rpc/rpc_category"
	"suzaku/pkg/common/config"
	"suzaku/pkg/common/kafka"
	"suzaku/pkg/factory"
	pb_chat "suzaku/pkg/proto/chart"
)

type chatRpcServer struct {
	pb_chat.UnimplementedChatServer
	rpc_category.Rpc
	producer *kafka.Producer
}

func NewChatRpcServer(port int) *chatRpcServer {
	var (
		rpc *chatRpcServer
	)
	rpc = &chatRpcServer{
		Rpc: rpc_category.NewRpcServer(port, config.Config.RPCRegisterName.OfflineMessageName),
	}
	rpc.producer = kafka.NewKafkaProducer(config.Config.Kafka.Ws2Mschat.Addr, config.Config.Kafka.Ws2Mschat.Topic)
	return rpc
}

func (rpc *chatRpcServer) Run() {
	var (
		server *grpc.Server
	)
	server = factory.NewRPCNewServer()
	pb_chat.RegisterChatServer(server, rpc)
	rpc.Rpc.RunServer(server)
}
