package ws_server

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"strconv"
	"sync"
)

type Hub struct {
	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients. 只在Client调用closeConn()函数时触发
	unregister chan *Client

	// 客户端发送的消息
	read chan *Message

	// Registered clients.
	clients map[string]*Client

	// key:MsgCode value:Call
	calls map[int][]*Call

	sync.Mutex
}

func newHub() *Hub {
	return &Hub{
		register:   make(chan *Client, 1000),
		unregister: make(chan *Client, 1000),
		read:       make(chan *Message, 1000),
		clients:    make(map[string]*Client),
		calls:      make(map[int][]*Call),
	}
}

func (h *Hub) registerClient(client *Client) {
	h.Lock()
	defer h.Unlock()
	h.clients[client.identifier] = client
}

func (h *Hub) unregisterClient(client *Client) {
	h.Lock()
	if _, ok := h.clients[client.identifier]; ok {
		delete(h.clients, client.identifier)
	}
	h.Unlock()
}

func (h *Hub) Run() {
	go func() {
		for {
			select {
			case client := <-h.register:
				h.registerClient(client)
			case client := <-h.unregister:
				h.unregisterClient(client)
			case read := <-h.read:
				h.messageHandler(read)
			}
		}
	}()
}

func (h *Hub) Send(identifier string, message []byte) {
	var (
		client *Client
		ok     bool
	)
	h.Lock()
	if client, ok = h.clients[identifier]; ok == false {
		h.Unlock()
		return
	}
	h.Unlock()
	client.sendMsg(message)
}

func (h *Hub) messageHandler(msg *Message) {
	var (
		calls []*Call
		call  *Call
		ok    bool
	)
	h.Lock()
	if calls, ok = h.calls[msg.MsgCode]; ok == false {
		h.Unlock()
		return
	}
	h.Unlock()

	for _, call = range calls {
		call.CallFunc(msg)
	}
}

var uid = 0

// serveWs handles websocket requests from the peer.
func (h *Hub) wsHandler(c *gin.Context) {
	var (
		conn   *websocket.Conn
		client *Client
		err    error
	)
	uid++
	if conn, err = upgrader.Upgrade(c.Writer, c.Request, nil); err != nil {
		log.Println(err)
		return
	}

	client = &Client{hub: h, conn: conn, identifier: strconv.Itoa(uid), send: make(chan []byte, WsMaxMessageSize), close: make(chan []byte, 0)}
	h.register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.write()
	go client.read()
}

func (h *Hub) RegisterCall(msgCode int, holder string, callFunc func(*Message) error) (err error) {
	var (
		cl *Call
		ok bool
	)
	if msgCode == 0 || holder == "" || callFunc == nil {
		//err = errors.New("参数错误")
		return
	}
	h.Lock()
	if _, ok = h.calls[msgCode]; ok == false {
		h.calls[msgCode] = make([]*Call, 0)
	}
	h.Unlock()
	for _, call := range h.calls[msgCode] {
		if call.Holder == holder {
			//重复添加
			//err = errors.New("重复注册回调")
			return
		}
	}
	cl = &Call{
		MsgCode:  msgCode,
		Holder:   holder,
		CallFunc: callFunc,
	}
	h.Lock()
	h.calls[msgCode] = append(h.calls[msgCode], cl)
	h.Unlock()
	return
}

func (h *Hub) Close(identifier string) {
	var (
		client *Client
		ok     bool
	)
	if identifier == "" {
		return
	}
	h.Lock()
	if client, ok = h.clients[identifier]; ok == false {
		h.Unlock()
		return
	}
	h.Unlock()
	client.closeClient()
}
