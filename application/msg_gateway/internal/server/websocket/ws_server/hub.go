package ws_server

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"sync"
	"time"
)

type Hub struct {
	rwLock   sync.RWMutex
	upgrader websocket.Upgrader
	cfg      *WsConfig
	// Register requests from the clients.
	register chan *Client
	// Unregister requests from clients. 只在Client调用closeConn()函数时触发
	unregister chan *Client
	// 客户端发送的消息
	read chan *Message
	// 消息处理
	handler MessageHandler
	// key1:UserID key2:platformID
	clients map[string]map[int32]*Client
	// 访问间隔
	access map[string]int64
	// 在线连接数
	onlineConnections int
}

func NewHub(cfg *WsConfig, handler MessageHandler) *Hub {
	cfg.WriteWaitTime = time.Duration(cfg.WriteWait) * time.Millisecond
	cfg.PongWaitTime = time.Duration(cfg.PongWait) * time.Millisecond
	cfg.PingPeriodTime = time.Duration(cfg.PingPeriod) * time.Millisecond

	return &Hub{
		upgrader: websocket.Upgrader{
			ReadBufferSize:  cfg.ReadBufferSize,
			WriteBufferSize: cfg.WriteBufferSize,
		},
		cfg:        cfg,
		register:   make(chan *Client, cfg.ChanServerRegister),
		unregister: make(chan *Client, cfg.ChanServerUnregister),
		read:       make(chan *Message, cfg.ChanServerReadMessage),
		handler:    handler,
		clients:    make(map[string]map[int32]*Client),
		access:     make(map[string]int64),
	}
}

func (h *Hub) registerClient(client *Client) {
	var (
		ok        bool
		cl        *Client
		platforms map[int32]*Client
	)
	h.rwLock.Lock()
	if platforms, ok = h.clients[client.userID]; ok == false {
		platforms = make(map[int32]*Client)
		h.clients[client.userID] = platforms
	}

	cl, ok = platforms[client.platformID]
	if ok == false {
		platforms[client.platformID] = client
		//atomic.AddInt64(&h.onlineConnections, 1)
		h.onlineConnections += 1
		h.rwLock.Unlock()
		return
	}

	if client.onlineAt > cl.onlineAt {
		platforms[client.platformID] = client
		h.rwLock.Unlock()
		h.close(cl)
		return
	}

	h.rwLock.Unlock()
	h.close(client)
}

func (h *Hub) close(client *Client) {
	// TODO：提示断开
	client.Send(nil)
	client.Close()
}

func (h *Hub) unregisterClient(client *Client) {
	var (
		ok        bool
		platforms map[int32]*Client
		cl        *Client
	)
	h.rwLock.Lock()
	defer h.rwLock.Unlock()
	if platforms, ok = h.clients[client.userID]; ok == false {
		return
	}
	if cl, ok = platforms[client.platformID]; ok == true {
		if cl == client {
			h.onlineConnections -= 1
			delete(platforms, client.platformID)
		}
	}
}

func (h *Hub) Run() {
	var (
		client *Client
		msg    *Message
	)
	go func() {
		for {
			select {
			case client = <-h.register:
				h.registerClient(client)
			case client = <-h.unregister:
				h.unregisterClient(client)
			case msg = <-h.read:
				h.messageHandler(msg)
			}
		}
	}()
}

func (h *Hub) IsOnline(userID string) (ok bool) {
	var (
		platforms map[int32]*Client
	)
	h.rwLock.RLock()
	defer h.rwLock.RUnlock()
	if platforms, ok = h.clients[userID]; ok == false {
		return
	}
	if len(platforms) > 0 {
		ok = true
	}
	return
}

func (h *Hub) Send(userID string, message []byte) (resultCode int) {
	var (
		platforms map[int32]*Client
		client    *Client
		ok        bool
	)
	h.rwLock.RLock()
	if platforms, ok = h.clients[userID]; ok == false {
		h.rwLock.RUnlock()
		resultCode = WsSendMsgOffline
		return
	}
	h.rwLock.RUnlock()
	if len(platforms) == 0 {
		resultCode = WsSendMsgOffline
		return
	}
	for _, client = range platforms {
		client.Send(message)
	}
	return
}

func (h *Hub) SendMessage(userID string, platformID int32, message []byte) (resultCode int, err error) {
	var (
		platforms map[int32]*Client
		client    *Client
		ok        bool
	)
	h.rwLock.RLock()
	if platforms, ok = h.clients[userID]; ok == false {
		h.rwLock.RUnlock()
		resultCode = WsSendMsgOffline
		return
	}
	client, ok = platforms[platformID]
	h.rwLock.RUnlock()
	if ok == false {
		resultCode = WsSendMsgOffline
		return
	}
	client.Send(message)
	resultCode = WsSendMsgFailed
	return
}

func (h *Hub) messageHandler(msg *Message) {
	h.handler.MsgCallback(msg)
}

// serveWs handles websocket requests from the peer.
func (h *Hub) wsHandler(c *gin.Context) {
	var (
		uidVal     interface{}
		pidVal     interface{}
		exists     bool
		userID     string
		platformID int32
		conn       *websocket.Conn
		client     *Client
		lastTs     int64
		nowTs      int64
		err        error
	)

	if h.onlineConnections >= h.cfg.MaxConnections {
		httpErr(c, ErrorWsExceedMaxConnections, ErrorCodeWsExceedMaxConnections)
		return
	}
	uidVal, exists = c.Get(WsKeyUserID)
	if exists == false {
		httpErr(c, ErrorHttpUserIDDoesNotExist, ErrorCodeHttpUserIDDoesNotExist)
		return
	}
	pidVal, exists = c.Get(WsKeyPlatformID)
	if exists == false {
		httpErr(c, ErrorHttpPlatformIDDoesNotExist, ErrorCodeHttpPlatformIDDoesNotExist)
		return
	}
	userID = uidVal.(string)
	platformID = int32(pidVal.(float64))

	nowTs = time.Now().UnixNano() / 1e6
	h.rwLock.Lock()
	lastTs, _ = h.access[userID]
	h.access[userID] = nowTs
	h.rwLock.Unlock()
	if nowTs-lastTs < h.cfg.MinimumTimeInterval {
		httpErr(c, ErrorCodeRequestTooMundane, ErrorCodeHttpRequestTooMundane)
		return
	}

	if conn, err = h.upgrader.Upgrade(c.Writer, c.Request, nil); err != nil {
		// 协议升级失败
		httpError(c, err, ErrorCodeHttpUpgraderFailed)
		return
	}
	client = newClient(h, conn, userID, platformID)
	h.register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.write()
	go client.read()
}
