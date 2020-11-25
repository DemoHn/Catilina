package model

import "gorm.io/gorm"

// User - 表示记账人
type User struct {
	gorm.Model
	Name string `gorm:"type:varchar(64)"`
}
