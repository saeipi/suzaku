package server_api

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"suzaku/internal/router/router_api"
	"suzaku/pkg/common/config"
	"suzaku/pkg/common/middleware"
)

var ApiServer *apiServer

type apiServer struct {
	engine *gin.Engine
}

func init() {
	var (
		engine *gin.Engine
	)
	gin.SetMode(gin.ReleaseMode)
	engine = gin.New()
	// 1、使用 Recovery 中间件
	engine.Use(gin.Recovery())
	// 2、跨域
	engine.Use(middleware.Cors())
	// 3、授权验证
	engine.Use(middleware.JwtAuth())

	router_api.RegisterRouter(engine)
	ApiServer = &apiServer{engine}
}

func (s *apiServer) Run() {
	var (
		addr string
	)
	addr = ":" + strconv.Itoa(config.Config.API.Port[0])
	s.engine.Run(addr)
}
