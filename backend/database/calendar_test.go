package database

import (
	"testing"

	"test3/model"

	. "github.com/smartystreets/goconvey/convey"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCalendarDatabase(t *testing.T) {
	Convey("TestCalendarDatabase", t, func() {
		db, _ = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})

		_ = db.AutoMigrate(&model.Calendar{})

		db.Create(&model.Calendar{Name: "fakeName"})

		Convey("GetCalendar 當 從資料庫找全部行事曆時 應 回傳全部行事曆", func() {
			result := GetCalendar()

			So(result[0].Name, ShouldEqual, "fakeName")
		})

		Convey("AddCalendar 當 新增行事曆時 應 新增行事曆", func() {
			result := AddCalendar(model.Calendar{Name: "fakeName2"})

			var myCalendar model.Calendar
			db.First(&myCalendar, 2)
			So(myCalendar.Name, ShouldEqual, "fakeName2")
			So(result, ShouldEqual, true)
		})

		Convey("UpdateCalendar 當 更新行事曆時 應 更新行事曆", func() {
			calendar := model.Calendar{Name: "fakeName3"}
			db.Save(&calendar)

			result := UpdateCalendar(calendar.ID, model.Calendar{Name: "fakeName5"})

			var myCalendar model.Calendar
			db.First(&myCalendar, calendar.ID)
			So(myCalendar.Name, ShouldEqual, "fakeName5")
			So(result, ShouldEqual, true)
		})

		Convey("UpdateCalendar 當 更新行事曆時找不到行事曆 應 回傳否", func() {
			result := UpdateCalendar(0, model.Calendar{})

			So(result, ShouldEqual, false)
		})

		Convey("DeleteCalendar 當 刪除行事曆時 應 刪除行事曆", func() {
			calendar := model.Calendar{Name: "fakeName4"}
			db.Save(&calendar)

			result := DeleteCalendar(calendar.ID)

			var myCalendar model.Calendar
			db.First(&myCalendar, calendar.ID)
			So(myCalendar.ID, ShouldEqual, 0)
			So(result, ShouldEqual, true)
		})
	})
}
