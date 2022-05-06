package po_mysql
import "time"

type Group struct {
	GroupId       string `gorm:"column:group_id;primary_key" json:"group_id"`
	GroupName     string `gorm:"column:group_name" json:"group_name"`           // 名称
	Notification  string `gorm:"column:notification" json:"notification"`       // 通知
	Introduction  string `gorm:"column:introduction" json:"introduction"`       // 介绍
	AvatarUrl     string `gorm:"column:avatar_url" json:"avatar_url"`           // 头像
	CreatorUserId string `gorm:"column:creator_user_id" json:"creator_user_id"` // 创建者ID
	GroupType     int    `gorm:"column:group_type;default:0" json:"group_type"`
	Status        int    `gorm:"column:status;default:0" json:"status"`
	CreateTs      int64  `gorm:"column:create_ts;default:0" json:"create_ts"`
	Ex            string `gorm:"column:ex" json:"ex"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

type GroupMember struct {
	GroupId        string `gorm:"column:group_id;primary_key" json:"group_id"`         // 群ID
	UserId         string `gorm:"column:user_id;NOT NULL" json:"user_id"`              // 用户ID
	Nickname       string `gorm:"column:nickname" json:"nickname"`                     // 在群中的昵称
	UserAvatarUrl  string `gorm:"column:user_avatar_url" json:"user_avatar_url"`       // 在群中的头像
	RoleLevel      int    `gorm:"column:role_level;default:0" json:"role_level"`       // 角色等级
	JoinTime       int64  `gorm:"column:join_time;default:0" json:"join_time"`         // 加入时间
	JoinSource     int    `gorm:"column:join_source;default:0" json:"join_source"`     // 来源
	OperatorUserId string `gorm:"column:operator_user_id" json:"operator_user_id"`     // 操作员
	MuteEndTime    int64  `gorm:"column:mute_end_time;default:0" json:"mute_end_time"` // 禁言结束时间
	Ex             string `gorm:"column:ex" json:"ex"`                                 // 扩展字段
}

type GroupRequest struct {
	UserId       string `gorm:"column:user_id;primary_key" json:"user_id"`           // 事件ID
	GroupId      string `gorm:"column:group_id" json:"group_id"`                     // 发起人ID
	HandleUserId string `gorm:"column:handle_user_id" json:"handle_user_id"`         // 处理人ID
	HandleResult int    `gorm:"column:handle_result;default:0" json:"handle_result"` // 结果
	HandleMsg    string `gorm:"column:handle_msg" json:"handle_msg"`                 // 处理消息
	ReqMsg       string `gorm:"column:req_msg" json:"req_msg"`                       // 添加好友消息
	ReqTs        int64  `gorm:"column:req_ts;default:0" json:"req_ts"`               // 请求时间
	HandleTs     int64  `gorm:"column:handle_ts;default:0" json:"handle_ts"`
	Ex           string `gorm:"column:ex" json:"ex"` // 扩展字段
}