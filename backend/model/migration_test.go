package model

import (
//  "encoding/json"
 "reflect"
 "testing"

 . "github.com/agiledragon/gomonkey/v2"
 . "github.com/smartystreets/goconvey/convey"
 "gorm.io/driver/sqlite"
 "gorm.io/gorm"
)

func TestMigration(t *testing.T) {
 Convey("migration", t, func() {
  db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})

  Convey("AutoMigrate_呼叫時展示資料不存在_應自動加入", func() {
   var calendar Calendar

   myPatches := NewPatches()
   defer myPatches.Reset()
   myPatches.ApplyMethod(reflect.TypeOf(db), "Set", func(*gorm.DB, string, interface{}) *gorm.DB { return db })

   AutoMigrate(db)

   So(db.Where("name = ? AND date = ? AND todo = ?", "小明", "2024-01-16", "出差").First(&calendar).RowsAffected, ShouldNotBeZeroValue)
  })
 })
}