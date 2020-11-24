package controller

import "github.com/gin-gonic/gin"

func listRecordsHandler(c *gin.Context) {

}

func addRecordsHandler(c *gin.Context) {

}

var recordRoutes = map[string]gin.HandlerFunc{
	"GET:/api/records/list": listRecordsHandler,
	"POST:/api/records/add": addRecordsHandler,
}
