package factory

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"suzaku/pkg/common/config"
	"time"
)
func NewRPCNewServer() (srv *grpc.Server) {
	var (
		keepParams grpc.ServerOption
	)
	keepParams = grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle:     time.Duration(time.Duration(config.Config.RPCKeepalive.IdleTimeout) * time.Millisecond),
		MaxConnectionAgeGrace: time.Duration(time.Duration(config.Config.RPCKeepalive.ForceCloseWait) * time.Millisecond),
		Time:                  time.Duration(time.Duration(config.Config.RPCKeepalive.KeepAliveInterval) * time.Millisecond),
		Timeout:               time.Duration(time.Duration(config.Config.RPCKeepalive.KeepAliveTimeout) * time.Millisecond),
		MaxConnectionAge:      time.Duration(time.Duration(config.Config.RPCKeepalive.MaxLifeTime) * time.Millisecond),
	})
	srv = grpc.NewServer(keepParams)
	return
}
