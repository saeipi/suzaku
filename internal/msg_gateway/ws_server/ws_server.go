package ws_server

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type WServer struct {
	address string
	hub     *Hub
	engine  *gin.Engine
}

func NewWServer(port int, callback MsgCallback) *WServer {
	var (
		ws *WServer
	)
	ws = &WServer{
		address: ":" + strconv.Itoa(port),
		hub:     NewHub(callback),
		engine:  gin.Default(),
	}
	ws.engine.GET("/", ws.hub.wsHandler)
	return ws
}

func (ws *WServer) Run() {
	ws.hub.Run()
	ws.engine.Run(ws.address)
}

func (ws *WServer) Send(userID string, msg []byte) (ok bool) {
	return ws.hub.Send(userID, msg)
}

func (ws *WServer) SendMessage(userID string, platformID int32, message []byte) (resultCode int, err error) {
	return ws.hub.SendMessage(userID, platformID, message)
}

func (ws *WServer) IsOnline(userID string) (ok bool) {
	return ws.hub.IsOnline(userID)
}
