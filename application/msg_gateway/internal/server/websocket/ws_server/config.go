package ws_server

import "time"

/*
const (
	// Time allowed to write a message to the peer.
	WsWriteWait = 10 * time.Second
	// Time allowed to read the next pong message from the peer.
	WsPongWait = 60 * time.Second
	// Send pings to peer with this period. Must be less than pongWait.
	WsPingPeriod = (WsPongWait * 9) / 10
	// Maximum message size allowed from peer.
	WsMaxMessageSize        = 4096
	WsReadBufferSize        = 4096
	WsWriteBufferSize       = 4096
	WsHeaderLength          = 4
	WsChanClientSendMessage = 100
	WsChanServerReadMessage = 1000
	WsChanServerRegister    = 1000
	WsChanServerUnregister  = 1000
	WsMaxConnections        = 20000
	WsMinimumTimeInterval   = -1 //5000 TODO: 测试值
)
*/
type WsConfig struct {
	WriteWait             int           `json:"write_wait"`
	WriteWaitTime         time.Duration `json:"write_wait_time"`
	PongWait              int           `json:"pong_wait"`
	PongWaitTime          time.Duration `json:"pong_wait_time"`
	PingPeriod            int           `json:"ping_period"`
	PingPeriodTime        time.Duration `json:"ping_period_time"`
	MaxMessageSize        int64         `json:"max_message_size"`
	ReadBufferSize        int           `json:"read_buffer_size"`
	WriteBufferSize       int           `json:"write_buffer_size"`
	HeaderLength          int           `json:"header_length"`
	ChanClientSendMessage int           `json:"chan_client_send_message"`
	ChanServerReadMessage int           `json:"chan_server_read_message"`
	ChanServerRegister    int           `json:"chan_server_register"`
	ChanServerUnregister  int           `json:"chan_server_unregister"`
	MaxConnections        int           `json:"max_connections"`
	MinimumTimeInterval   int64         `json:"minimum_time_interval"`
}

const (
	WsMsgCodePing  = 10001
	WsMsgCodePong  = 10002
	WsMsgCodeClose = 10003
)

const (
	WsSendMsgOffline = -2
	WsSendMsgFailed  = -1
	WsSendMsgSuccess = 0
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
