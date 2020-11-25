package service

import (
	"fmt"
	"time"

	"github.com/DemoHn/Catilina/app/model"
)

// ListSearch 用于构造搜索条件
type ListSearch struct {
}

// AddRecord 添加一条记账记录
func AddRecord(name string, amount int, debt bool, issueTime time.Time) (model.Record, error) {
	var user model.User
	var record model.Record
	result := db.Where("name = ?", name).First(&user)

	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return record, fmt.Errorf("错误：用户 (%s) 不存在", name)
		}
		return record, result.Error
	}

	record = model.Record{
		UID:       int(user.ID),
		IssueDate: issueTime,
		IsDebt:    debt,
		Amount:    amount,
	}

	result2 := db.Create(&record)
	if result2.Error != nil {
		return record, result2.Error
	}
	return record, nil
}

// ListRecords 用于给定条件下列出一定的记账记录
func ListRecords(search ListSearch) []model.Record {
	return []model.Record{}
}
