package main

import (
	"suzaku/pkg/common/config"
	"suzaku/pkg/common/mysql"
)

func main() {
	config.Config.Mysql.Address = []string{"192.168.184.134:8066"}
	config.Config.Mysql.Username = "root"
	config.Config.Mysql.Password = "123456"
	config.Config.Mysql.Db = "suzaku"

	db, err := mysql.GetDB(config.Config.Mysql.Address[0], config.Config.Mysql.Db)
	if err != nil {
		return
	}
	var list []Order
	err = db.Raw("SELECT * FROM orders").Find(&list).Error
	if err != nil {
		return
	}
}

type Order struct {
	Id         int64  `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	OrderType  int    `gorm:"column:order_type" json:"order_type"`
	CustomerId int    `gorm:"column:customer_id" json:"customer_id"`
	Amount     string `gorm:"column:amount" json:"amount"`
}
