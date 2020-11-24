package model

// ref: https://github.com/cgrant/gin-gorm-api-example/blob/master/main.go
// ref2: https://developer.aliyun.com/article/655425
import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Init - init
func Init(url string) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(url), &gorm.Config{})
}
