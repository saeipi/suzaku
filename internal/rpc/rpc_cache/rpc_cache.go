package rpc_cache

import (
	"google.golang.org/grpc"
	"suzaku/internal/rpc/rpc_category"
	"suzaku/pkg/common/config"
	"suzaku/pkg/factory"
	pb_cache "suzaku/pkg/proto/cache"
)

type cacheRpcServer struct {
	pb_cache.UnimplementedCacheServer
	rpc_category.Rpc
}

func NewCacheRpcServer(port int) *cacheRpcServer {
	var (
		rpc *cacheRpcServer
	)
	rpc = &cacheRpcServer{
		Rpc: rpc_category.NewRpcServer(port, config.Config.RPCRegisterName.CacheName),
	}
	return rpc
}

func (rpc *cacheRpcServer) Run() {
	var (
		server *grpc.Server
	)
	server = factory.NewRPCNewServer()
	pb_cache.RegisterCacheServer(server, rpc)
	rpc.Rpc.RunServer(server)
}

func NewCacheClient() (client pb_cache.CacheClient) {
	var (
		clientConn *grpc.ClientConn
	)
	clientConn = factory.ClientConn(config.Config.RPCRegisterName.CacheName)
	client = pb_cache.NewCacheClient(clientConn)
	return
}

func syncFriendsList() {

}