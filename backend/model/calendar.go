package model

import (
	"gorm.io/gorm"
)

// 行事曆
type Calendar struct {
	gorm.Model

	Name string `gorm:"comment:姓名"`
	Date string `gorm:"comment:日期"`
	Todo string `gorm:"comment:代辦事項"`
}
