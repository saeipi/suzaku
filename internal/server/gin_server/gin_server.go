package gin_server

import (
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"strconv"
	"suzaku/pkg/common/middleware"
)

type GinServer struct {
	Engine *gin.Engine
}

func NewGinServer() *GinServer {
	var (
		engine *gin.Engine
	)
	gin.SetMode(gin.ReleaseMode)
	engine = gin.New()
	// http://127.0.0.1:10000/debug/pprof/
	pprof.Register(engine)
	// 1、使用 Recovery 中间件
	engine.Use(gin.Recovery())
	// 2、跨域
	engine.Use(middleware.Cors())
	return &GinServer{engine}
}

func (s *GinServer) Run(port int) {
	var (
		addr string
		err  error
	)
	addr = ":" + strconv.Itoa(port)
	err = s.Engine.Run(addr)
	if err != nil {
		fmt.Println("GinServer Start Failed.", err.Error())
	}
}
