package service

import (
	"time"

	"github.com/DemoHn/Catilina/app/model"
)

// ListSearch 用于构造搜索条件
type ListSearch struct {
}

// AddRecord 添加一条记账记录
func AddRecord(name string, amount int, debt bool, issueTime time.Time) {

}

// ListRecords 用于给定条件下列出一定的记账记录
func ListRecords(search ListSearch) []model.Record {

}
