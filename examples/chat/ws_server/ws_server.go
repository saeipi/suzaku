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

func Init(port int) *WServer {
	var (
		ws *WServer
	)
	ws = &WServer{
		address: ":" + strconv.Itoa(port),
		hub:     newHub(),
		engine:  gin.Default(),
	}
	ws.engine.GET("/", ws.hub.wsHandler)
	return ws
}

func (ws *WServer) Run() {
	ws.hub.Run()
	ws.engine.Run(ws.address)
}
