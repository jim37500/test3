package api

import (
	"test3/database"
	"test3/model"

	"github.com/gofiber/fiber/v2"
)

// Get行事曆
func GetCalendar(context *fiber.Ctx) error {
	return context.JSON(database.GetCalendar())
}

// 新增行事曆
func AddCalendar(context *fiber.Ctx) error {
	var myCalendar model.Calendar
	_ = context.BodyParser(&myCalendar)

	// 當 新增行事曆成功時 則 回傳成功
	if database.AddCalendar(myCalendar) {
		return context.SendStatus(fiber.StatusOK)
	}

	// 否則 回傳錯誤
	return context.SendStatus(fiber.StatusInternalServerError)
}

// 修改行事曆
func UpdateCalendar(context *fiber.Ctx) error {
	id, _ := context.ParamsInt("id")

	var myCalendar model.Calendar
	_ = context.BodyParser(&myCalendar)

	// 當 更新行事曆成功時 則 回傳成功
	if database.UpdateCalendar(uint(id), myCalendar) {
		return context.SendStatus(fiber.StatusOK)
	}

	// 否則 回傳錯誤
	return context.SendStatus(fiber.StatusInternalServerError)
}

// 刪除行事曆
func DeleteCalendar(context *fiber.Ctx) error {
	id, _ := context.ParamsInt("id")

	// 當 刪除行事曆成功時 則 回傳成功
	if database.DeleteCalendar(uint(id)) {
		return context.SendStatus(fiber.StatusOK)
	}

	// 否則 回傳錯誤
	return context.SendStatus(fiber.StatusInternalServerError)
}
