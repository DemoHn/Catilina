package controller

import "github.com/gin-gonic/gin"

import "github.com/DemoHn/Catilina/app/service"

type userCreateReq struct {
	Name string `json:"name"`
}

func userCreateHandler(c *gin.Context) {
	var req userCreateReq

	if err := c.ShouldBindJSON(&req); err != nil {
		handleRsp(c, nil, err)
		return
	}

	name := req.Name
	user, err := service.CreateUser(name)

	userData := map[string]interface{}{
		"id":   user.ID,
		"name": user.Name,
	}
	handleRsp(c, userData, err)
}

var userRoutes = map[string]gin.HandlerFunc{
	"POST:/api/users/create": userCreateHandler,
}
