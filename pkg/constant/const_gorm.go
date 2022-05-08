package constant

import (
	"time"
)

type GormModel struct {
	//ID        int64     `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at;default:NULL" json:"deleted_at"`
}

type Gorm struct {
	ID        int64     `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	DeletedAt time.Time `gorm:"column:deleted_at;default:NULL" json:"deleted_at"`
}

type GormTs struct {
	CreatedTs int64 `gorm:"column:created_ts;default:0" json:"created_ts"`
	UpdatedTs int64 `gorm:"column:updated_ts;default:0" json:"updated_ts"`
	DeletedTs int64 `gorm:"column:deleted_ts;default:0" json:"deleted_ts"`
}

func Update(ts *GormTs) {
	ts.UpdatedTs = time.Now().Unix()
}

func Create(ts *GormTs) {
	ts.CreatedTs = time.Now().Unix()
	ts.UpdatedTs = ts.CreatedTs
}

func Delete(ts *GormTs) {
	ts.DeletedTs = time.Now().Unix()
}
