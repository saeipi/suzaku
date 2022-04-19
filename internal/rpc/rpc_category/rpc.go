package rpc_category

import (
	"google.golang.org/grpc"
	"net"
	"strconv"
	"strings"
	"suzaku/pkg/common/config"
	"suzaku/pkg/common/grpc-etcdv3/getcdv3"
	"suzaku/pkg/utils"
)

type Rpc struct {
	Port         int
	RegisterName string
	Schema       string
	Address      []string
}

func NewRpcServer(port int, registerName string) Rpc {
	return Rpc{
		Port:         port,
		RegisterName: registerName,
		Schema:       config.Config.Etcd.Schema,
		Address:      config.Config.Etcd.Address,
	}
}

func (rpc *Rpc) RunServer(server *grpc.Server) {
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
