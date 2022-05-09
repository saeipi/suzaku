package po_mysql

type GormCreatedTs struct {
	CreatedTs int64 `gorm:"column:created_ts;autoCreateTime:milli" json:"created_ts"`
}

type GormUpdatedTs struct {
	UpdatedTs int64 `gorm:"column:updated_ts;autoUpdateTime:milli" json:"updated_ts"`
}

type GormDeletedTs struct {
	DeletedTs int64 `gorm:"column:deleted_ts;default:0" json:"deleted_ts"`
}

type GormTs struct {
	CreatedTs int64 `gorm:"column:created_ts;autoCreateTime:milli" json:"created_ts"`
	UpdatedTs int64 `gorm:"column:updated_ts;autoUpdateTime:milli" json:"updated_ts"`
}

type GormModel struct {
	CreatedTs int64 `gorm:"column:created_ts;autoCreateTime:milli" json:"created_ts"`
	UpdatedTs int64 `gorm:"column:updated_ts;autoUpdateTime:milli" json:"updated_ts"`
	DeletedTs int64 `gorm:"column:deleted_ts;default:0" json:"deleted_ts"`
}
