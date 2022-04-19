package constant

import (
	"time"
)

type GormModel struct {
	ID        int64     `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at;default:NULL" json:"deleted_at"`
}

type Gorm struct {
	ID        int64     `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	DeletedAt time.Time `gorm:"column:deleted_at;default:NULL" json:"deleted_at"`
}
