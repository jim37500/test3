package database

import (
	"time"

	"test3/configuration"
	"test3/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

// 連線資料庫
func Open() {
	var err error
	db, err = gorm.Open(mysql.Open(configuration.Connectionstring), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})

	if err != nil {
		time.Sleep(time.Second)
		Open()
		return
	}

	sqlDB, _ := db.DB()

	// 連線池中空閒連線的最大數量
	sqlDB.SetMaxIdleConns(10)

	// 資料庫連線的最大數量。
	sqlDB.SetMaxOpenConns(100)

	// 連線最長可複用的時間
	sqlDB.SetConnMaxLifetime(time.Hour)

	// 自動轉移資料庫結構
	model.AutoMigrate(db)
}
