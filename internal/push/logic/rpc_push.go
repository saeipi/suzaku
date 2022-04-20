package logic

import (
	"context"
	"google.golang.org/grpc"
	"suzaku/internal/rpc/rpc_category"
	"suzaku/pkg/common/config"
	pb_push "suzaku/pkg/proto/push"
)

type pushRpc struct {
	pb_push.UnimplementedPushMsgServiceServer
	rpc_category.Rpc
}

func NewPushRpcServer(port int) *pushRpc {
	return &pushRpc{
		Rpc: rpc_category.NewRpcServer(port, config.Config.RPCRegisterName.PushName),
	}
}

func (rpc *pushRpc) Run() {
	var (
		server *grpc.Server
	)
	server = grpc.NewServer()
	pb_push.RegisterPushMsgServiceServer(server, rpc)
	rpc.Rpc.RunServer(server)
}

func (rpc *pushRpc) PushMsg(_ context.Context, req *pb_push.PushMsgReq) (resp *pb_push.PushMsgResp, err error) {
	MsgToUser(req)
	resp = &pb_push.PushMsgResp{}
	return
}
