package mysql

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"suzaku/pkg/common/config"
	"sync"
	"time"
)

var MysqlDB *mysqlDB

type mysqlDB struct {
	sync.RWMutex
	dbMap map[string]*gorm.DB
}

func init() {
	MysqlDB = &mysqlDB{dbMap: make(map[string]*gorm.DB)}
	MysqlDB.open(config.Config.Mysql.Address[0], config.Config.Mysql.Db)
}

func dbKey(address string, dbName string) string {
	return address + "_" + dbName
}

func (m *mysqlDB) open(address string, dbName string) (err error) {
	var (
		args  string
		opts  *gorm.Config
		sqlDB *sql.DB
		db    *gorm.DB
		key   string
	)
	args = fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=true&loc=Local",
		config.Config.Mysql.Username,
		config.Config.Mysql.Password,
		address,
		dbName)

	opts = &gorm.Config{
		SkipDefaultTransaction: false, // 禁用默认事务(true: Error 1295: This command is not supported in the prepared statement protocol yet)
		PrepareStmt:            false, // 创建并缓存预编译语句(true: Error 1295)
	}

	db, err = gorm.Open(mysql.Open(args), opts)
	if err != nil {
		fmt.Println("Failed to connect db:", err)
		return
	}

	sqlDB, err = db.DB()
	if err != nil {
		fmt.Println("Failed to sql DB:", err)
		return
	}
	//设置最大空闲连接
	sqlDB.SetMaxIdleConns(config.Config.Mysql.MaxIdleConn)
	//设置最大连接数
	sqlDB.SetMaxOpenConns(config.Config.Mysql.MaxOpenConn)
	//设置连接超时时间
	sqlDB.SetConnMaxLifetime(time.Duration(config.Config.Mysql.ConnLifetime) * time.Millisecond)

	key = dbKey(address, dbName)
	m.dbMap[key] = db
	return
}

func GetDB(address string, dbName string) (db *gorm.DB, err error) {
	db, err = MysqlDB.getDB(address, dbName)
	return
}

func (m *mysqlDB) getDB(address string, dbName string) (db *gorm.DB, err error) {
	var (
		key string
		ok  bool
	)
	key = dbKey(address, dbName)
	if _, ok = m.dbMap[key]; ok == false {
		if err = m.open(address, dbName); err != nil {
			return
		}
	}
	m.Lock()
	defer m.Unlock()
	db = m.dbMap[key]
	return
}

func GormDB() (db *gorm.DB, err error) {
	return GetDB(config.Config.Mysql.Address[0], config.Config.Mysql.Db)
}

//事务处理
func Transaction(handle func(tx *gorm.DB) (err error)) (err error) {
	var (
		db *gorm.DB
	)
	db, err = GormDB()
	if err != nil {
		return
	}
	tx := db.Begin(&sql.TxOptions{Isolation: sql.LevelRepeatableRead})
	err = handle(tx)
	if err != nil {
		return
	}
	err = tx.Commit().Error
	return
}
