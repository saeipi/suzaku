package po_mongo

type Message struct {
	ServerMsgId      string `bson:"server_msg_id"`      // 服务端生成
	ClientMsgId      string `bson:"client_msg_id"`      // 客户端生成
	SendId           string `bson:"send_id"`            // 发送人ID
	RecvId           string `bson:"recv_id"`            // 接收人ID 或 群ID
	SenderPlatformId int    `bson:"sender_platform_id"` // 发送人平台ID
	SenderNickname   string `bson:"sender_nickname"`
	SenderAvatarUrl  string `bson:"sender_avatar_url"`
	SessionId        string `bson:"session_id"`   // 单例:会话ID,群聊:群ID
	SessionType      int    `bson:"session_type"` // 1:单聊 2:群聊
	Seq              int    `bson:"seq"`          // 会话消息唯一ID
	MsgFrom          int    `bson:"msg_from"`     // 100:用户消息 200:系统消息
	ContentType      int    `bson:"content_type"`
	Content          string `bson:"content"`
	Status           int    `bson:"status"`
	SendTs           int64  `bson:"send_ts"`    // 消息发送的具体时间(毫秒)
	CreatedTs        int64  `bson:"created_ts"` // 创建消息的时间，在send_ts之前
	Ex               string `bson:"ex"`
}
