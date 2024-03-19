package model

import (
	"gorm.io/gorm"
)

// 自動轉移資料庫結構
func AutoMigrate(db *gorm.DB) {
	migrateTable(db, "行事曆", &Calendar{})

	checkData(db) // 檢查資料
}

// 轉移資料表結構
func migrateTable(db *gorm.DB, tableName string, structure interface{}) {
	_ = db.Set("gorm:table_options", " COMMENT='"+tableName+"'").AutoMigrate(structure)
}

// 檢查資料 若 不存在 則 自動加入
func checkData(db *gorm.DB) {
	if db.Where("name = ?", "小明").Find(&Calendar{}).RowsAffected == 0 {
		db.Save(&Calendar{Name: "小明", Date: "2024-01-16", Todo: "出差"})
	}
}
