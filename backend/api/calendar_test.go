package api

import (
	"testing"

	"test3/database"
	"test3/model"

	. "github.com/agiledragon/gomonkey/v2"
	"github.com/gofiber/fiber/v2"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCalendarApi(t *testing.T) {
	var fakeContext *fiber.Ctx

	Convey("TestCalendarApi", t, func() {
		Convey("GetCalendar 當 呼叫時 應 取得行事曆", func() {
			var result []model.Calendar
			myPatches := NewPatches()
			defer myPatches.Reset()
			myPatches.
				ApplyFuncReturn(database.GetCalendar, []model.Calendar{{Name: "fakeName"}}).
				ApplyMethod(fakeContext, "JSON", func(_ *fiber.Ctx, data interface{}) error {
					result = data.([]model.Calendar)
					return nil
				})

			_ = GetCalendar(fakeContext)

			So(result[0].Name, ShouldEqual, "fakeName")
		})

		Convey("AddCalendar 當 新增行事曆成功時 應 回傳成功", func() {
			var result int
			myPatches := NewPatches()
			defer myPatches.Reset()
			myPatches.
				ApplyMethod(fakeContext, "BodyParser", func(context *fiber.Ctx, out interface{}) error { return nil }).
				ApplyFuncReturn(database.AddCalendar, true).
				ApplyMethod(fakeContext, "SendStatus", func(context *fiber.Ctx, status int) error {
					result = status
					return nil
				})

			_ = AddCalendar(fakeContext)

			So(result, ShouldEqual, fiber.StatusOK)
		})

		Convey("AddCalendar 當 新增行事曆失敗時 應 回傳失敗", func() {
			var result int
			myPatches := NewPatches()
			defer myPatches.Reset()
			myPatches.
				ApplyMethod(fakeContext, "BodyParser", func(context *fiber.Ctx, out interface{}) error { return nil }).
				ApplyFuncReturn(database.AddCalendar, false).
				ApplyMethod(fakeContext, "SendStatus", func(context *fiber.Ctx, status int) error {
					result = status
					return nil
				})

			_ = AddCalendar(fakeContext)

			So(result, ShouldEqual, fiber.StatusInternalServerError)
		})

		Convey("UpdateCalendar 當 更新行事曆成功時 應 回傳成功", func() {
			var result int
			myPatches := NewPatches()
			defer myPatches.Reset()
			myPatches.
				ApplyMethodReturn(fakeContext, "ParamsInt", 0, nil).
				ApplyMethod(fakeContext, "BodyParser", func(context *fiber.Ctx, out interface{}) error { return nil }).
				ApplyFuncReturn(database.UpdateCalendar, true).
				ApplyMethod(fakeContext, "SendStatus", func(context *fiber.Ctx, status int) error {
					result = status
					return nil
				})

			_ = UpdateCalendar(fakeContext)

			So(result, ShouldEqual, fiber.StatusOK)
		})

		Convey("UpdateCalendar 當 更新行事曆失敗時 應 回傳失敗", func() {
			var result int
			myPatches := NewPatches()
			defer myPatches.Reset()
			myPatches.
				ApplyMethodReturn(fakeContext, "ParamsInt", 0, nil).
				ApplyMethod(fakeContext, "BodyParser", func(context *fiber.Ctx, out interface{}) error { return nil }).
				ApplyFuncReturn(database.UpdateCalendar, false).
				ApplyMethod(fakeContext, "SendStatus", func(context *fiber.Ctx, status int) error {
					result = status
					return nil
				})

			_ = UpdateCalendar(fakeContext)

			So(result, ShouldEqual, fiber.StatusInternalServerError)
		})

		Convey("DeleteCalendar 當 刪除行事曆成功時 應 回傳成功", func() {
			var result int
			myPatches := NewPatches()
			defer myPatches.Reset()
			myPatches.
				ApplyMethodReturn(fakeContext, "ParamsInt", 1, nil).
				ApplyFuncReturn(database.DeleteCalendar, true).
				ApplyMethod(fakeContext, "SendStatus", func(context *fiber.Ctx, status int) error {
					result = status
					return nil
				})

			_ = DeleteCalendar(fakeContext)

			So(result, ShouldEqual, fiber.StatusOK)
		})

		Convey("DeleteCalendar 當 刪除行事曆失敗時 應 回傳失敗", func() {
			var result int
			myPatches := NewPatches()
			defer myPatches.Reset()
			myPatches.
				ApplyMethodReturn(fakeContext, "ParamsInt", 1, nil).
				ApplyFuncReturn(database.DeleteCalendar, false).
				ApplyMethod(fakeContext, "SendStatus", func(context *fiber.Ctx, status int) error {
					result = status
					return nil
				})

			_ = DeleteCalendar(fakeContext)

			So(result, ShouldEqual, fiber.StatusInternalServerError)
		})
	})
}
