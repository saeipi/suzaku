package callback

type CallbackBeforeSendSingleMsgReq struct {
	CommonCallbackReq
	RecvID string `json:"recv_id"`
}

type CallbackBeforeSendSingleMsgResp struct {
	CommonCallbackResp
}

type CallbackAfterSendSingleMsgReq struct {
	CommonCallbackReq
	RecvID string `json:"recv_id"`
}

type CallbackAfterSendSingleMsgResp struct {
	CommonCallbackResp
}

type CallbackBeforeSendGroupMsgReq struct {
	CommonCallbackReq
	GroupID string `json:"group_id"`
}

type CallbackBeforeSendGroupMsgResp struct {
	CommonCallbackResp
}

type CallbackAfterSendGroupMsgReq struct {
	CommonCallbackReq
	GroupID string `json:"group_id"`
}

type CallbackAfterSendGroupMsgResp struct {
	CommonCallbackResp
}

type CallbackWordFilterReq struct {
	CommonCallbackReq
}

type CallbackWordFilterResp struct {
	CommonCallbackResp
	Content string `json:"content"`
}
