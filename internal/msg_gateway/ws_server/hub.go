package ws_server

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"strconv"
	"sync"
	"time"
)

type Hub struct {
	sync.Mutex
	rwLock   sync.RWMutex
	upgrader websocket.Upgrader
	// Register requests from the clients.
	register chan *Client
	// Unregister requests from clients. 只在Client调用closeConn()函数时触发
	unregister chan *Client
	// 客户端发送的消息
	read chan *Message
	// 回调
	callback MsgCallback
	// key1:UserID key2:platformID
	clients map[string]map[int32]*Client
	// 访问间隔
	access map[string]int64
	// 在线连接数
	onlineConnections int64
}

func NewHub(callback MsgCallback) *Hub {
	return &Hub{
		upgrader: websocket.Upgrader{
			ReadBufferSize:  WsReadBufferSize,
			WriteBufferSize: WsWriteBufferSize,
		},
		register:   make(chan *Client, WsChanServerRegister),
		unregister: make(chan *Client, WsChanServerUnregister),
		read:       make(chan *Message, WsChanServerReadMessage),
		callback:   callback,
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
	h.Lock()
	if platforms, ok = h.clients[client.userID]; ok == false {
		platforms = make(map[int32]*Client)
		h.clients[client.userID] = platforms
	}
	cl, ok = platforms[client.platformID]
	if ok == true {
		if client.onlineAt > cl.onlineAt {
			platforms[client.platformID] = client
			h.Unlock()
			h.close(cl)
		} else {
			h.Unlock()
			h.close(client)
		}
		return
	}
	platforms[client.platformID] = client
	//atomic.AddInt64(&h.onlineConnections, 1)
	h.onlineConnections += 1
	h.Unlock()
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
	h.Lock()
	defer h.Unlock()
	if platforms, ok = h.clients[client.userID]; ok == false {
		return
	}
	if cl, ok = platforms[client.platformID]; ok == true {
		if cl == client {
			delete(platforms, client.platformID)
		}
	}
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

func (h *Hub) Send(userID string, message []byte) (ok bool) {
	var (
		platforms map[int32]*Client
		client    *Client
	)
	h.rwLock.RLock()
	if platforms, ok = h.clients[userID]; ok == false {
		h.rwLock.RUnlock()
		return
	}
	h.rwLock.RUnlock()
	for _, client = range platforms {
		client.Send(message)
	}
	ok = true
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
		resultCode = WsSendMsgOffline
		h.rwLock.RUnlock()
		return
	}
	client, ok = platforms[platformID]
	h.rwLock.RUnlock()
	if ok == false {
		resultCode = WsSendMsgOffline
		return
	}
	err = client.SendMessage(message)
	if err != nil {
		resultCode = WsSendMsgFailed
	}
	return
}

func (h *Hub) messageHandler(msg *Message) {
	h.callback(msg)
}

// serveWs handles websocket requests from the peer.
func (h *Hub) wsHandler(c *gin.Context) {
	var (
		userID     string
		platform   string
		val        int64
		platformID int32
		conn       *websocket.Conn
		client     *Client
		lastTs     int64
		nowTs      int64
		err        error
	)

	if h.onlineConnections >= WsMaxConnections {
		httpErr(c, ErrorWsExceedMaxConnections, ErrorCodeWsExceedMaxConnections)
		return
	}
	userID = c.GetHeader(WsKeyUserID)
	if userID == "" {
		httpErr(c, ErrorHttpUserIDDoesNotExist, ErrorCodeHttpUserIDDoesNotExist)
		return
	}
	platform = c.GetHeader(WsKeyPlatformID)
	if platform == "" {
		httpErr(c, ErrorHttpPlatformIDDoesNotExist, ErrorCodeHttpPlatformIDDoesNotExist)
		return
	}
	val, err = strconv.ParseInt(platform, 10, 32)
	if err != nil {
		httpError(c, err, ErrorCodeHttpPlatformIDDoesNotExist)
		return
	}
	platformID = int32(val)

	nowTs = time.Now().UnixNano() / 1e6
	lastTs, _ = h.access[userID]
	h.access[userID] = nowTs
	if nowTs-lastTs < WsMinimumTimeInterval {
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
