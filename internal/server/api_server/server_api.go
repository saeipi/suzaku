package api_server

import (
	"suzaku/internal/router/router_api"
	"suzaku/internal/server/gin_server"
	"suzaku/pkg/common/config"
)

type ApiServer struct {
	Gin *gin_server.GinServer
}

func NewApiServer() *ApiServer {
	var (
		ginServer *gin_server.GinServer
	)
	ginServer = gin_server.NewGinServer()
	router_api.RegisterRouter(ginServer.Engine)
	return &ApiServer{ginServer}
}

func (s *ApiServer) Run() {
	s.Gin.Run(config.Config.API.Port[0])
}
