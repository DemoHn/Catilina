package service

import (
	"gorm.io/gorm"
)

var db *gorm.DB

// Init services
func Init(gDB *gorm.DB) error {
	db = gDB

	return nil
}
