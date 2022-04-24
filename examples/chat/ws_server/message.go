package ws_server

type Message struct {
	// 身份标识
	Identifier string `json:"identifier"`
	// 消息编号
	MsgCode int `json:"msg_code"`
	// 消息本体
	Body interface{} `json:"body"`
}

type Call struct {
	MsgCode  int
	Holder   string
	CallFunc func(msg *Message) (err error)
}
