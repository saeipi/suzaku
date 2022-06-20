package ws_server

import (
	"github.com/gorilla/websocket"
	"log"
	"sync"
	"time"
)

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	rwLock sync.RWMutex
	hub    *Hub
	// The websocket connection.
	conn *websocket.Conn
	// 用户ID
	userID string
	// 平台ID
	platformID int32
	// 上线时间戳（毫秒）
	onlineAt int64
	// Buffered channel of outbound messages.
	send chan []byte
	// 关闭通知
	close  chan []byte
	closed bool
}

func newClient(hub *Hub, conn *websocket.Conn, userID string, platformID int32) *Client {
	return &Client{
		hub:        hub,
		conn:       conn,
		userID:     userID,
		platformID: platformID,
		onlineAt:   time.Now().UnixNano() / 1e6,
		send:       make(chan []byte, hub.cfg.ChanClientSendMessage),
		close:      make(chan []byte),
	}
}

func (c *Client) closeConn() {
	c.rwLock.Lock()
	if c.closed == true {
		c.rwLock.Unlock()
		return
	}
	c.closed = true
	close(c.send)
	close(c.close)
	c.rwLock.Unlock()

	c.conn.Close()
	c.hub.unregister <- c
}

func (c *Client) read() {
	defer func() {
		c.closeConn()
	}()

	var (
		msgType int
		bufMsg  []byte
		//bufHeader []byte
		//msgCode   int
		message *Message
		err     error
	)

	c.conn.SetReadLimit(c.hub.cfg.MaxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(c.hub.cfg.PongWaitTime))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(c.hub.cfg.PongWaitTime)); return nil })

	for {
		if msgType, bufMsg, err = c.conn.ReadMessage(); err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				//TODO:需要添加日志
				log.Printf("error: %v", err)
			}
			break
		}
		if msgType == websocket.PingMessage {
			continue
		}
		if len(bufMsg) == 0 {
			continue
		}
		/*
			bufHeader = bufMsg[0:WsHeaderLength]
			msgCode = bytes2Int(bufHeader)
			if msgCode == WsMsgCodePing {
				c.Send(WsMsgBufPong)
				continue
			}*/
		message = &Message{
			Client: c,
			Body:   bufMsg,
		}
		c.hub.read <- message
	}
}

func (c *Client) write() {
	pingTicker := time.NewTicker(c.hub.cfg.PingPeriodTime)
	defer func() {
		pingTicker.Stop()
		c.closeConn()
	}()

	var (
		err     error
		message []byte
		ok      bool
	)

	for {
		select {
		case message, ok = <-c.send:
			if ok == false {
				// chan 已关闭
				return
			}
			if err = c.conn.SetWriteDeadline(time.Now().Add(c.hub.cfg.WriteWaitTime)); err != nil {
				c.conn.WriteMessage(websocket.CloseMessage, WsMsgBufClose)
				return
			}
			if err = c.conn.WriteMessage(websocket.BinaryMessage, message); err != nil {
				return
			}
		case <-pingTicker.C:
			c.conn.SetWriteDeadline(time.Now().Add(c.hub.cfg.WriteWaitTime))
			if err = c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		case <-c.close:
			return
		}
	}
}

func (c *Client) Send(message []byte) {
	c.rwLock.RLock()
	defer c.rwLock.RUnlock()
	if c.closed == true {
		return
	}
	c.send <- message
}

/*
func (c *Client) SendMessage(message []byte) (err error) {
	if c.closed {
		return
	}
	if err = c.conn.SetWriteDeadline(time.Now().Add(WsWriteWait)); err != nil {
		c.conn.WriteMessage(websocket.CloseMessage, WsMsgBufClose)
		return
	}
	if err = c.conn.WriteMessage(websocket.BinaryMessage, message); err != nil {
		return
	}
	return
}
*/

func (c *Client) Close() {
	c.rwLock.RLock()
	defer c.rwLock.RUnlock()
	if c.closed == true {
		return
	}
	c.close <- WsMsgBufClose
}
