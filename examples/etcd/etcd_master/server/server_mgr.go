package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"suzaku/micro/etcd/etcd_master/ctrl"
	"suzaku/pkg/common/config"
)

type ServerMgr struct {
	server *gin.Engine
	cfg    config.ServerConfig
}

var (
	SG_SERMGR *ServerMgr
)

func InitRouter(cfg config.ServerConfig) (err error) {
	var (
		addr = fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	)
	SG_SERMGR = &ServerMgr{
		server: gin.Default(),
		cfg:    cfg,
	}
	SG_SERMGR.SetApi()
	err = SG_SERMGR.server.Run(addr)
	return
}

func (s *ServerMgr) SetApi() {
	var ctrl = ctrl.NewJobCtrl()
	routes := s.server.Group("job")
	{
		routes.POST("save", ctrl.Save)
		routes.POST("delete", ctrl.Delete)
	}
}
