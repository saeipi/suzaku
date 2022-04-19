package ws_server

import (
	"github.com/gorilla/websocket"
	"time"
)

const (
	// Time allowed to write a message to the peer.
	WsWriteWait = 10 * time.Second
	// Time allowed to read the next pong message from the peer.
	WsPongWait = 60 * time.Second
	// Send pings to peer with this period. Must be less than pongWait.
	WsPingPeriod = (WsPongWait * 9) / 10
	// Maximum message size allowed from peer.
	WsMaxMessageSize  = 512
	WsReadBufferSize  = 1024
	WsWriteBufferSize = 1024
	WsHeaderLength    = 4
)

const (
	WsMsgTypeKey       = "msg_type"
	WsMsgTypeUserInfo  = "msg_user_info"
	WsMsgTypeError     = "msg_error"
	WsMsgTypePing      = "msg_ping"
	WsMsgTypePong      = "msg_pong"
	WsMsgTypeClose     = "msg_close"
	WsMsgTypeBroadcast = "msg_broadcast"
)

const (
	WsMsgCodePing      = 10001
	WsMsgCodePong      = 10002
	WsMsgCodeClose     = 10003
	WsMsgCodeBroadcast = 10004
	WsMsgCodeSelfInfo  = 10005
)

var (
	WsMsgBufNewline = []byte{'\n'}
	WsMsgBufPing    = msgBuffer(WsMsgCodePing, nil)
	WsMsgBufPong    = msgBuffer(WsMsgCodePong, nil)
	WsMsgBufClose   = make([]byte, 0)
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  WsReadBufferSize,
	WriteBufferSize: WsWriteBufferSize,
}
