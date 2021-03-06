package do

type CommonCallbackReq struct {
	SendId           string `json:"send_id"`
	CallbackCommand  string `json:"callback_command"`
	ServerMsgId      string `json:"server_msg_id"`
	ClientMsgId      string `json:"client_msg_id"`
	OperationId      string `json:"operation_id"`
	SenderPlatformId int32  `json:"sender_platform_id"`
	SenderNickname   string `json:"sender_nickname"`
	SessionType      int32  `json:"session_type"`
	MsgFrom          int32  `json:"msg_from"`
	ContentType      int32  `json:"content_type"`
	Status           int32  `json:"status"`
	CreatedTs        int64  `json:"created_ts"`
	Content          string `json:"content"`
}

type CommonCallbackResp struct {
	ActionCode  int    `json:"action_code"`
	Code        int    `json:"code"`
	Msg         string `json:"msg"`
	OperationId string `json:"operation_id"`
}

type CallbackAfterSendGroupMsgReq struct {
	CommonCallbackReq
	GroupID string `json:"groupID"`
}

type CallbackBeforeSendGroupMsgResp struct {
	CommonCallbackResp
}
