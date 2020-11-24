package app

import (
	"github.com/DemoHn/Catilina/app/controller"
	"github.com/DemoHn/Catilina/app/model"
	"github.com/DemoHn/Catilina/app/service"
)

// StartServer - 启动服务
func StartServer(conf string) {
	db, err := model.Init(":")
	if err != nil {
		panic(err)
	}

	service.Init(db)
	r := controller.Init()
	r.Run()
}
