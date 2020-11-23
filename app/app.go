package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// StartServer - 启动服务
func StartServer(conf string) {
	fmt.Println("Hello World!")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
