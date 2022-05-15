package po_mysql

type Register struct {
	GormTs
	UserId   string `gorm:"column:user_id;primary_key" json:"user_id"` // 用户ID 系统生成
	Password string `gorm:"column:password" json:"password"`           // 密码
	Ex       string `gorm:"column:ex" json:"ex"`                       // 扩展字段
}
