package main

import (
	"gorm.io/gorm"
	"suzaku/pkg/common/mysql"
)

func main() {
	var (
		db  *gorm.DB
		err error
	)
	if db, err = mysql.GormDB(); err != nil {
		return
	}

	//for i := 0; i < 10; i++ {
	//	second := Second{Second: int64(i)}
	//	err = db.Save(&second).Error
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//}

	//for i := 0; i < 10; i++ {
	//	err = db.Model(Second{}).Where("id=?",i+1).Update("second",10+1).Error
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//}

	db.Where("id=10").Delete(&Second{})
}

//type Second struct {
//	Id        int   `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
//	Second    int64 `gorm:"column:second;default:0" json:"second"`
//	CreatedTs int64 `gorm:"column:created_ts;autoCreateTime:milli" json:"created_ts"`
//	UpdatedTs int64 `gorm:"column:updated_ts;autoUpdateTime:milli" json:"updated_ts"`
//	DeletedAt int64 `gorm:"column:deleted_at;default:0" json:"deleted_at"`
//}

type Second struct {
	Id        int   `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Second    int64 `gorm:"column:second;default:0" json:"second"`
	CreatedAt int64 `gorm:"column:created_at" json:"created_at"`
	UpdatedAt int64 `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt int64 `gorm:"column:deleted_at" json:"deleted_at"`
}
