package po_mysql

type GormTs struct {
	CreatedTs int64 `gorm:"column:created_ts;autoCreateTime:milli" json:"created_ts"`
	UpdatedTs int64 `gorm:"column:updated_ts;autoUpdateTime:milli" json:"updated_ts"`
	DeletedAt int64 `gorm:"column:deleted_at;default:0" json:"deleted_at"`
}
