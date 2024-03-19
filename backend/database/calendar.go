package database

import (
	"test3/model"
)

// Get行事曆 
func GetCalendar() (calendar []model.Calendar) {
	db.Find(&calendar)

	return calendar
}

// 新增行事曆
func AddCalendar(newData model.Calendar) bool {
	return db.Save(&newData).Error == nil
}

// 修改行事曆
func UpdateCalendar(id uint, calendar model.Calendar) bool {
	var myCalendar model.Calendar
	db.First(&myCalendar, id)

	// 若 沒有資料 則 回傳否
	if myCalendar.Model.ID == 0 {
		return false
	}

	myCalendar.Name = calendar.Name
	myCalendar.Date = calendar.Date
	myCalendar.Todo = calendar.Todo

	return db.Save(&myCalendar).Error == nil
}

// 刪除行事曆
func DeleteCalendar(id uint) bool {
	var myCalendar model.Calendar
	db.First(&myCalendar, id)

	return db.Delete(&myCalendar).Error == nil
}
