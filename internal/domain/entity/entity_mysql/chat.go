package entity_mysql

import (
	"time"
)

type ChatLog struct {
	ServerMsgID      string    `gorm:"column:server_msg_id;primary_key;type:char(64)" json:"serverMsgID"`
	ClientMsgID      string    `gorm:"column:client_msg_id;type:char(64)" json:"clientMsgID"`
	SendID           string    `gorm:"column:send_id;type:char(64)" json:"sendID"`
	RecvID           string    `gorm:"column:recv_id;type:char(64)" json:"recvID"`
	SenderPlatformID int32     `gorm:"column:sender_platform_id" json:"senderPlatformID"`
	SenderNickname   string    `gorm:"column:sender_nick_name;type:varchar(255)" json:"senderNickname"`
	SenderFaceURL    string    `gorm:"column:sender_face_url;type:varchar(255)" json:"senderFaceURL"`
	SessionType      int32     `gorm:"column:session_type" json:"sessionType"`
	MsgFrom          int32     `gorm:"column:msg_from" json:"msgFrom"`
	ContentType      int32     `gorm:"column:content_type" json:"contentType"`
	Content          string    `gorm:"column:content;type:varchar(3000)" json:"content"`
	Status           int32     `gorm:"column:status" json:"status"`
	SendTime         time.Time `gorm:"column:send_time" json:"sendTime"`
	CreateTime       time.Time `gorm:"column:create_time" json:"createTime"`
	Ex               string    `gorm:"column:ex;type:varchar(1024)" json:"ex"`
}
