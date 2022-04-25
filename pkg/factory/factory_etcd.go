package factory

import (
	"google.golang.org/grpc"
	"strings"
	"suzaku/pkg/common/grpc-etcdv3/getcdv3"
	"suzaku/pkg/common/config"
)

func ClientConn(serviceName string) (clientConn *grpc.ClientConn) {
	clientConn = getcdv3.GetConn(config.Config.Etcd.Schema, strings.Join(config.Config.Etcd.Address, ","), serviceName)
	return
}