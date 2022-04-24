package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"suzaku/examples/etcd/etcd_master/cfg"
	"suzaku/examples/etcd/etcd_master/ctrl"
)

type ServerMgr struct {
	server *gin.Engine
	cfg    *cfg.Server
}

var (
	SG_SERMGR *ServerMgr
)

func InitRouter(cfg *cfg.Server) (err error) {
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
