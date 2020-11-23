package model

// ref: https://github.com/cgrant/gin-gorm-api-example/blob/master/main.go
// ref2: https://developer.aliyun.com/article/655425
import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// DB - DB instance
var DB *gorm.DB

// Init - init
func Init() error {

}
