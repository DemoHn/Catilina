package app

import (
	"fmt"
	"io/ioutil"

	"github.com/DemoHn/Catilina/app/controller"
	"github.com/DemoHn/Catilina/app/model"
	"github.com/DemoHn/Catilina/app/service"
	"gopkg.in/yaml.v2"
)

type configModel struct {
	Server struct {
		Port int
	}
	Db struct {
		Connections map[string]struct {
			Host      string
			Port      int
			Database  string
			Username  string
			Password  string
			Charset   string
			Collation string
		}
	}
}

// StartServer - 启动服务
func StartServer(confFile string) error {
	data, err := parseConfig(confFile)
	if err != nil {
		return err
	}

	dsn, err := getDefaultDSN(data)
	if err != nil {
		return err
	}

	db, err := model.Init(dsn)
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.Record{})
	db.AutoMigrate(&model.User{})
	service.Init(db)
	r := controller.Init()
	return r.Run()
}

// helpers

// parseConfig - 从配置文件中解析YAML数据，并将 yaml 转换成 map
func parseConfig(confFile string) (configModel, error) {
	data, err := ioutil.ReadFile(confFile)
	if err != nil {
		return configModel{}, err
	}

	var m configModel

	err = yaml.Unmarshal(data, &m)
	if err != nil {
		return configModel{}, err
	}

	return m, nil
}

// getDefaultDSN - 读取配置，并将数据库配置格式化成 gorm 认识的 DSN
func getDefaultDSN(m configModel) (string, error) {

	conns := m.Db.Connections
	defaultConn, ok := conns["default"]
	if !ok {
		return "", fmt.Errorf("数据库配置不存在")
	}

	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true",
		defaultConn.Username,
		defaultConn.Password,
		defaultConn.Host,
		defaultConn.Port,
		defaultConn.Database,
		defaultConn.Charset), nil
}
