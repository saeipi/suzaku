package ws_server

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"strconv"
)

type WServer struct {
	address  string
	hub      *Hub
	engine   *gin.Engine
	validate *validator.Validate
}

func NewWServer(port int, callback MsgCallback) *WServer {
	var (
		ws *WServer
	)
	ws = &WServer{
		address:  ":" + strconv.Itoa(port),
		hub:      NewHub(callback),
		engine:   gin.Default(),
		validate: validator.New(),
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

func (ws *WServer) IsOnline(userID string) (ok bool) {
	return ws.hub.IsOnline(userID)
}

func (ws *WServer) SendMessage(userID string, platformID int32, msg []byte) (ok bool) {
	return ws.hub.Send(userID, msg)
}
