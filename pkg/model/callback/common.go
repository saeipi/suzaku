package callback

type CommonCallbackReq struct {
	SendID           string `json:"send_id"`
	CallbackCommand  string `json:"callback_command"`
	ServerMsgID      string `json:"server_msg_id"`
	ClientMsgID      string `json:"client_msg_id"`
	OperationID      string `json:"operation_id"`
	SenderPlatformID int32  `json:"sender_platform_id"`
	SenderNickname   string `json:"sender_nickname"`
	SessionType      int32  `json:"session_type"`
	MsgFrom          int32  `json:"msg_from"`
	ContentType      int32  `json:"content_type"`
	Status           int32  `json:"status"`
	CreateTime       int64  `json:"create_time"`
	Content          string `json:"content"`
}

type CommonCallbackResp struct {
	ActionCode  int    `json:"action_code"`
	ErrCode     int    `json:"err_code"`
	ErrMsg      string `json:"err_msg"`
	OperationID string `json:"operation_id"`
}
