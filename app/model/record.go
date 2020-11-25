package model

import (
	"time"

	"gorm.io/gorm"
)

// Record 用于表示此账本的一条记录
type Record struct {
	gorm.Model
	Amount    int       `gorm:"comment:记账金额"`
	IssueDate time.Time `gorm:"comment:记账时间"`
	IsDebt    bool      `gorm:"comment:记账类型 0-收入 1-支出"`
	UID       int       `gorm:"comment:用户id"`
}
