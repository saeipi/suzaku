package po_mysql

type FriendRequest struct {
	FromUserId   string `gorm:"column:from_user_id;primary_key" json:"from_user_id"` // 发起人ID
	ToUserId     string `gorm:"column:to_user_id" json:"to_user_id"`                 // 目标人ID
	HandleUserId string `gorm:"column:handle_user_id" json:"handle_user_id"`         // 处理人ID
	HandleResult int    `gorm:"column:handle_result;default:0" json:"handle_result"` // 结果
	ReqMsg       string `gorm:"column:req_msg" json:"req_msg"`                       // 添加好友消息
	HandleMsg    string `gorm:"column:handle_msg" json:"handle_msg"`                 // 处理消息
	HandledTs    int64  `gorm:"column:handled_ts;default:0" json:"handled_ts"`
	ReqTs        int64  `gorm:"column:req_ts;default:0" json:"req_ts"` // 请求时间
	Ex           string `gorm:"column:ex" json:"ex"`                   // 扩展字段
}