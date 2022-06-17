package rpc_server

import (
	"google.golang.org/grpc"
	"net"
	"strconv"
	"strings"
	"suzaku/pkg/common/grpc-etcdv3/getcdv3"
	"suzaku/pkg/utils"
)

type RpcServer struct {
	Port         int
	RegisterName string
	Schema       string
	Address      []string
}

func NewRpcServer(port int, registerName string, schema string, address []string) RpcServer {
	return RpcServer{
		Port:         port,
		RegisterName: registerName,
		Schema:       schema,
		Address:      address,
	}
}

func (rpc *RpcServer) Run(server *grpc.Server) {
	var (
		address  string
		listener net.Listener
		err      error
	)
	defer server.GracefulStop()

	address = utils.ServerIP + ":" + strconv.Itoa(rpc.Port)
	listener, err = net.Listen("tcp", address)
	if err != nil {
		return
	}
	err = getcdv3.RegisterEtcd(rpc.Schema, strings.Join(rpc.Address, ","), utils.ServerIP, rpc.Port, rpc.RegisterName, 10)
	if err != nil {
		return
	}
	err = server.Serve(listener)
	if err != nil {
		return
	}
}
