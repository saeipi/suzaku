package po_mysql

import (
	"time"
)

type FriendRequest struct {
	ReqId          string    `gorm:"column:req_id;primary_key" json:"req_id"`             // 事件ID
	FromUserId     string    `gorm:"column:from_user_id" json:"from_user_id"`             // 发起人ID
	ToUserId       string    `gorm:"column:to_user_id" json:"to_user_id"`                 // 目标人ID
	OperatorUserId string    `gorm:"column:operator_user_id" json:"operator_user_id"`     // 处理人ID
	HandleResult   int       `gorm:"column:handle_result;default:0" json:"handle_result"` // 结果
	ReqMsg         string    `gorm:"column:req_msg" json:"req_msg"`                       // 添加好友消息
	HandleMsg      string    `gorm:"column:handle_msg" json:"handle_msg"`                 // 处理消息
	Ex             string    `gorm:"column:ex" json:"ex"`                                 // 扩展字段
	CreatedAt      time.Time `gorm:"column:created_at" json:"created_at"`
	HandleAt       time.Time `gorm:"column:handle_at;default:NULL" json:"handle_at"` // 处理时间
}