package api_server

import (
	"suzaku/internal/router"
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
	router.Register(ginServer.Engine)
	return &ApiServer{ginServer}
}

func (s *ApiServer) Run() {
	s.Gin.Run(config.Config.API.Port[0])
}
