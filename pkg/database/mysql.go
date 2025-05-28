package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(dsn string) error {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	DB = db

	// 测试连接
	sqlDB, _ := DB.DB()
	if err := sqlDB.Ping(); err != nil {
		return err
	}
	println("数据库连接成功！")
	return nil
}
