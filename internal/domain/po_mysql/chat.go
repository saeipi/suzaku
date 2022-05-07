package po_mysql

/*
CreatedTs:调用 SendMsg(context.Context, *SendMsgReq) (*SendMsgResp, error) 的时间
SendTs:进入消息队列的时间
*/
type Message struct {
	ServerMsgId      string `gorm:"column:server_msg_id;primary_key" json:"server_msg_id"`         // 服务端生成
	ClientMsgId      string `gorm:"column:client_msg_id" json:"client_msg_id"`                     // 客户端生成
	SendId           string `gorm:"column:send_id" json:"send_id"`                                 // 发送人ID
	RecvId           string `gorm:"column:recv_id" json:"recv_id"`                                 // 接收人ID 或 群ID
	SenderPlatformId int    `gorm:"column:sender_platform_id;default:0" json:"sender_platform_id"` // 发送人平台ID
	SenderNickname   string `gorm:"column:sender_nickname" json:"sender_nickname"`
	SenderAvatarUrl  string `gorm:"column:sender_avatar_url" json:"sender_avatar_url"`
	SessionType      int    `gorm:"column:session_type;default:0" json:"session_type"` // 1:单聊 2:群聊
	MsgFrom          int    `gorm:"column:msg_from;default:0" json:"msg_from"`         // 100:用户消息 200:系统消息
	ContentType      int    `gorm:"column:content_type;default:0" json:"content_type"`
	Content          string `gorm:"column:content" json:"content"`
	Status           int    `gorm:"column:status;default:0" json:"status"`
	SendTs           int64  `gorm:"column:send_ts;default:0" json:"send_ts"`       // 进入消息队列的时间
	CreatedTs        int64  `gorm:"column:created_ts;default:0" json:"created_ts"` // 创建消息的时间，在send_ts前面
	Ex               string `gorm:"column:ex" json:"ex"`
}
