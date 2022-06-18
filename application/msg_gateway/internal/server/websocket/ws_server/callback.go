package ws_server

type Message struct {
	Client *Client `json:"client"`
	// 消息本体
	Body []byte `json:"body"`
}

type MessageHandler interface {
	MsgCallback(msg *Message)
}
