package controller

import (
	"strings"

	"github.com/gin-gonic/gin"
)

// Init -
func Init() *gin.Engine {
	r := gin.Default()
	registerHandlers(r, recordRoutes, userRoutes)

	return r
}

func registerHandlers(r *gin.Engine, handlers ...map[string]gin.HandlerFunc) {
	for _, h := range handlers {
		for url, handler := range h {
			urlData := strings.Split(url, ":")
			if len(urlData) == 2 {
				method := urlData[0]
				route := urlData[1]

				switch method {
				case "GET":
					r.GET(route, handler)
				case "POST":
					r.POST(route, handler)
				}
			}
		}
	}
}

//// helpers
func handleRsp(c *gin.Context, data interface{}, err error) {
	if err != nil {
		errorData := map[string]interface{}{
			"data":    "",
			"message": err.Error(),
			"code":    500,
		}
		c.JSON(200, errorData)
	} else {
		successData := map[string]interface{}{
			"data":    data,
			"message": "",
			"code":    0,
		}
		c.JSON(200, successData)
	}
}
