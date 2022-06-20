package ws_server

import (
	"suzaku/application/msg_gateway/internal/config"
	"suzaku/internal/server/gin_server"
	"suzaku/pkg/common/middleware"
)

type WServer struct {
	cfg *config.WsServer
	hub *Hub
	gin *gin_server.GinServer
}

func NewServer(cfg *config.WsServer, handler MessageHandler) *WServer {
	var (
		ws    *WServer
		wscfg *WsConfig
	)
	ws = &WServer{
		cfg: cfg,
		hub: NewHub(wscfg, handler),
		gin: gin_server.NewGinServer(),
	}
	ws.gin.Engine.Use(middleware.JwtAuth())
	ws.gin.Engine.GET("/", ws.hub.wsHandler)
	return ws
}

func (ws *WServer) Run() {
	ws.hub.Run()
	go ws.gin.Run(ws.cfg.Port)
}

func (ws *WServer) Send(userID string, msg []byte) (resultCode int) {
	return ws.hub.Send(userID, msg)
}

func (ws *WServer) SendMessage(userID string, platformID int32, message []byte) (resultCode int, err error) {
	return ws.hub.SendMessage(userID, platformID, message)
}

func (ws *WServer) IsOnline(userID string) (ok bool) {
	return ws.hub.IsOnline(userID)
}
