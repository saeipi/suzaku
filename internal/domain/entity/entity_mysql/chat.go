package entity_mysql

import (
	"time"
)

/*
serverMsgID生成:
func GetMsgID(sendID string) string {
	t := time.Now().Format("2006-01-02 15:04:05")
	return utils.Md5(t + "-" + sendID + "-" + strconv.Itoa(rand.Int()))
}
*/

type ChatLog struct {
	ServerMsgID      string    `gorm:"column:server_msg_id;primary_key;type:char(64)" json:"serverMsgID"` // 服务端生成
	ClientMsgID      string    `gorm:"column:client_msg_id;type:char(64)" json:"clientMsgID"`             // 客户端生成
	SendID           string    `gorm:"column:send_id;type:char(64)" json:"sendID"`                        // 发送人ID
	RecvID           string    `gorm:"column:recv_id;type:char(64)" json:"recvID"`                        // 发送人ID 或 群ID
	SenderPlatformID int32     `gorm:"column:sender_platform_id" json:"senderPlatformID"`                 // 发送人平台ID
	SenderNickname   string    `gorm:"column:sender_nick_name;type:varchar(255)" json:"senderNickname"`
	SenderFaceURL    string    `gorm:"column:sender_face_url;type:varchar(255)" json:"senderFaceURL"`
	SessionType      int32     `gorm:"column:session_type" json:"sessionType"` // 1:单聊 2:群聊
	MsgFrom          int32     `gorm:"column:msg_from" json:"msgFrom"`         // 100:用户消息 200:系统消息
	ContentType      int32     `gorm:"column:content_type" json:"contentType"`
	Content          string    `gorm:"column:content;type:varchar(3000)" json:"content"`
	Status           int32     `gorm:"column:status" json:"status"`
	SendTime         time.Time `gorm:"column:send_time" json:"sendTime"`
	CreateTime       time.Time `gorm:"column:create_time" json:"createTime"`
	Ex               string    `gorm:"column:ex;type:varchar(1024)" json:"ex"`
}
