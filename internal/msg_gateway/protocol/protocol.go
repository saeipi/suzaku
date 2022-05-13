package protocol

type MessageReq struct {
	ReqIdentifier int32  `json:"req_identifier" validate:"required"`
	Token         string `json:"token" `
	SendID        string `json:"send_id" validate:"required"`
	OperationID   string `json:"operation_id" validate:"required"`
	MsgIncr       string `json:"msg_incr" validate:"required"`
	Data          []byte `json:"data"`
}

type MessageResp struct {
	ReqIdentifier int32  `json:"req_identifier"`
	MsgIncr       string `json:"msg_incr"`
	OperationID   string `json:"operation_id"`
	Code          int32  `json:"code"`
	Msg           string `json:"msg"`
	Data          []byte `json:"data"`
}

type SeqData struct {
	SeqBegin int64 `mapstructure:"seq_begin" validate:"required"`
	SeqEnd   int64 `mapstructure:"seq_end" validate:"required"`
}

type MessageData struct {
	PlatformID  int32                  `mapstructure:"platform_id" validate:"required"`
	SessionType int32                  `mapstructure:"session_type" validate:"required"`
	MsgFrom     int32                  `mapstructure:"msg_from" validate:"required"`
	ContentType int32                  `mapstructure:"content_type" validate:"required"`
	RecvID      string                 `mapstructure:"recv_id" validate:"required"`
	ForceList   []string               `mapstructure:"force_list"`
	Content     string                 `mapstructure:"content" validate:"required"`
	Options     map[string]interface{} `mapstructure:"options" validate:"required"`
	ClientMsgID string                 `mapstructure:"client_msg_id" validate:"required"`
	OfflineInfo map[string]interface{} `mapstructure:"offline_info" validate:"required"`
	Ext         map[string]interface{} `mapstructure:"ext"`
}

type MaxSeqResp struct {
	MaxSeq int64 `json:"max_seq"`
}

type PullMessageResp struct {
}

type SeqListData struct {
	SeqList []int64 `mapstructure:"seq_list" validate:"required"`
}
