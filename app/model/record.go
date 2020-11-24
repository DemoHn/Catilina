package model

import (
	"time"

	"gorm.io/gorm"
)

// Record 用于表示此账本的一条记录
type Record struct {
	gorm.Model
	Amount    int       // 记账数目
	IssueDate time.Time // 记账时间
	IsDebt    bool      // 是否为支出 true=支出 false=收入
	User      User      // 记账人
}
