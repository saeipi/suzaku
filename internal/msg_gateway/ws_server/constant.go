package ws_server

import (
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
	WsMaxMessageSize        = 512
	WsReadBufferSize        = 1024
	WsWriteBufferSize       = 1024
	WsHeaderLength          = 4
	WsChanClientSendMessage = 100
	WsChanServerReadMessage = 1000
	WsChanServerRegister    = 1000
	WsChanServerUnregister  = 1000
	WsMaxConnections        = 20000
	WsMinimumTimeInterval   = 5000
)

const (
	WsMsgCodePing  = 10001
	WsMsgCodePong  = 10002
	WsMsgCodeClose = 10003
)

var (
	WsMsgBufNewline = []byte{'\n'}
	WsMsgBufPing    = make([]byte, 0)
	WsMsgBufPong    = make([]byte, 0)
	WsMsgBufClose   = make([]byte, 0)
)

const (
	WsKeyUserID     = "user_id"
	WsKeyPlatformID = "platform_id"
)

const (
	ErrorCodeHttpUpgraderFailed         = 15001
	ErrorCodeHttpUserIDDoesNotExist     = 15002
	ErrorCodeHttpPlatformIDDoesNotExist = 15003
	ErrorCodeHttpRequestTooMundane      = 15004
)

const (
	ErrorHttpUserIDDoesNotExist     = "user_id 缺失"
	ErrorHttpPlatformIDDoesNotExist = "platform_id 缺失"
	ErrorCodeRequestTooMundane      = "请求过于平凡"
)

const (
	ErrorCodeWsExceedMaxConnections = 16001
)

const (
	ErrorWsExceedMaxConnections = "超出最大连接数限制"
)
