package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB

func ConnectMysql(connectStr string) error {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: connectStr,
	}))
	if err != nil {
		return fmt.Errorf("数据库连接失败：%s", err)
	}

	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	DB = db
	db.AutoMigrate(&User{}, &Topic{}, &Reply{})
	return nil
}
