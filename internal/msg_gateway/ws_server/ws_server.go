package ws_server

import (
	"suzaku/internal/server/gin_server"
	"suzaku/pkg/common/middleware"
)

type WServer struct {
	port int
	hub  *Hub
	gin  *gin_server.GinServer
}

func NewWServer(port int, callback MsgCallback) *WServer {
	var (
		ws *WServer
	)
	ws = &WServer{
		port: port,
		hub:  NewHub(callback),
		gin:  gin_server.NewGinServer(),
	}
	ws.gin.Engine.Use(middleware.JwtAuth())
	ws.gin.Engine.GET("/", ws.hub.wsHandler)
	return ws
}

func (ws *WServer) Run() {
	ws.hub.Run()
	ws.gin.Run(ws.port)
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
