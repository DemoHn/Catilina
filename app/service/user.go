package service

import "github.com/DemoHn/Catilina/app/model"

// CreateUser - 创建一个用户
func CreateUser(name string) (model.User, error) {
	user := model.User{
		Name: name,
	}

	result := db.Create(&user)
	if result.Error != nil {
		return model.User{}, result.Error
	}
	return user, nil
}
