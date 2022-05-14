package main

import (
	"context"
	"fmt"
	"suzaku/internal/rpc/rpc_cache"
	"suzaku/pkg/proto/pb_user"
)

func main() {
	rpc := rpc_cache.NewCacheRpcServer(10688)
	rpc.GetUserInfo(context.Background(),&pb_user.UserInfoReq{UserId:"1524255191468085248"})
	var input int
	fmt.Scan(&input)
}
