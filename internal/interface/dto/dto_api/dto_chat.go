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
	SessionId   string   `json:"session_id" binding:"required"`
	MsgFrom     int32    `json:"msg_from" binding:"required"`
	ContentType int32    `json:"content_type" binding:"required"`
	RecvID      string   `json:"recv_id"`
	GroupID     string   `json:"group_id"`
	ForceList   []string `json:"force_list"`
	//Content     []byte                 `json:"content" binding:"required"`
	Content     []byte                 `json:"content"`
	Options     map[string]bool        `json:"options" `
	ClientMsgID string                 `json:"client_msg_id" binding:"required"`
	CreatedTs   int64                  `json:"created_ts" binding:"required"`
	OffLineInfo *pb_ws.OfflinePushInfo `json:"off_line_info,omitempty" `
}

type SendMsgResp struct {
	ServerMsgId string `json:"server_msg_id,omitempty"`
	ClientMsgId string `json:"client_msg_id,omitempty"`
	SendTime    int64  `json:"send_time,omitempty"`
}

type HistoryMessagesReq struct {
	PageSize    int    `form:"page_size" json:"page_size" example:"10" binding:"required,min=10,max=100"`
	Seq         int64  `form:"seq" json:"seq"`                                      // 最早/后一条消息的Sequence ID
	CreatedTs   int64  `form:"created_ts" json:"created_ts"`                        // 创建消息的时间
	SessionId   string `form:"session_id" json:"session_id" binding:"required"`     // 回话ID
	SessionType int32  `form:"session_type" json:"session_type" binding:"required"` // 1:单聊 2:群聊
	Back        bool   `form:"back" json:"back" binding:"required"`                 // true:往后查,false:向前查
}

type Message struct {
	ServerMsgId      string `gorm:"column:server_msg_id;primary_key" json:"server_msg_id"`         // 服务端生成
	ClientMsgId      string `gorm:"column:client_msg_id" json:"client_msg_id"`                     // 客户端生成
	SendId           string `gorm:"column:send_id" json:"send_id"`                                 // 发送人ID
	RecvId           string `gorm:"column:recv_id" json:"recv_id"`                                 // 接收人ID 或 群ID
	SenderPlatformId int    `gorm:"column:sender_platform_id;default:0" json:"sender_platform_id"` // 发送人平台ID
	SenderNickname   string `gorm:"column:sender_nickname" json:"sender_nickname"`
	SenderAvatarUrl  string `gorm:"column:sender_avatar_url" json:"sender_avatar_url"`
	SessionId        string `gorm:"column:session_id" json:"session_id"`               // 单例:会话ID,群聊:群ID
	SessionType      int    `gorm:"column:session_type;default:0" json:"session_type"` // 1:单聊 2:群聊
	Seq              int    `gorm:"column:seq;default:0" json:"seq"`                   // 会话消息唯一ID
	MsgFrom          int    `gorm:"column:msg_from;default:0" json:"msg_from"`         // 100:用户消息 200:系统消息
	ContentType      int    `gorm:"column:content_type;default:0" json:"content_type"`
	Content          string `gorm:"column:content" json:"content"`
	Status           int    `gorm:"column:status;default:0" json:"status"`
	SendTs           int64  `gorm:"column:send_ts;default:0" json:"send_ts"`       // 消息发送的具体时间(毫秒)
	CreatedTs        int64  `gorm:"column:created_ts;default:0" json:"created_ts"` // 创建消息的时间
	Ex               string `gorm:"column:ex" json:"ex"`
}

type HistoryMessagesResp struct {
	MsgList []*Message `json:"msg_list"`
}
