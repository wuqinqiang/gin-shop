package db

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	MysqlUrl = "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&parseTime=True&timeout=5s"
)

var Db *gorm.DB

func init() {
	Db, err := gorm.Open(mysql.Open(fmt.Sprintf(MysqlUrl, "root", "root", "127.0.0.1", "3306", "gin-shop")), &gorm.Config{})
	if err != nil {
		fmt.Printf("open db:%v", err)
	}
	sqlDb, err := Db.DB()
	if err != nil {
		fmt.Printf("sql Db:%v", err)
	}
	sqlDb.SetMaxOpenConns(10)
	sqlDb.SetConnMaxIdleTime(10)
	sqlDb.SetConnMaxLifetime(5 * time.Minute)
}
