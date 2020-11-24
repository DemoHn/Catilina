package controller

import (
	"strings"

	"github.com/gin-gonic/gin"
)

// Init -
func Init() *gin.Engine {
	r := gin.Default()
	registerHandlers(r, recordRoutes)

	return r
}

func registerHandlers(r *gin.Engine, handlers map[string]gin.HandlerFunc) {
	for url, handler := range handlers {
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
