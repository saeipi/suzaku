package dto_api

import "suzaku/pkg/proto/pb_ws"

type SendMsgReq struct {
	SenderPlatformID int32       `json:"sender_platform_id" binding:"required"`
	SendID           string      `json:"send_id" binding:"required"`
	SenderNickName   string      `json:"sender_nick_name"`
	SenderAvatarUrl  string      `json:"sender_avatar_url"`
	OperationID      string      `json:"operation_id" binding:"required"`
	Data             SendMsgData `json:"data"`
}

type SendMsgData struct {
	SessionType int32    `json:"session_type" binding:"required"`
	MsgFrom     int32    `json:"msg_from" binding:"required"`
	ContentType int32    `json:"content_type" binding:"required"`
	RecvID      string   `json:"recv_id" `
	GroupID     string   `json:"group_id" `
	ForceList   []string `json:"force_list"`
	//Content     []byte                 `json:"content" binding:"required"`
	Content     []byte                 `json:"content"`
	Options     map[string]bool        `json:"options" `
	ClientMsgID string                 `json:"client_msg_id" binding:"required"`
	CreatedTs  int64                   `json:"created_ts" binding:"required"`
	OffLineInfo *pb_ws.OfflinePushInfo `json:"off_line_info,omitempty" `
}

type SendMsgResp struct {
	ServerMsgId string `json:"server_msg_id,omitempty"`
	ClientMsgId string `json:"client_msg_id,omitempty"`
	SendTime    int64  `json:"send_time,omitempty"`
}
