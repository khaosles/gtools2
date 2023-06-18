package controller

import (
	"sync"

	"github.com/khaosles/gtools2/components/example/model"
	"github.com/khaosles/gtools2/components/example/service"
	"github.com/khaosles/gtools2/components/g"
)

/*
   @File: user_controller.go
   @Author: khaosles
   @Time: 2023/6/13 10:07
   @Desc:
*/

var (
	UserController = new(userController)
	once           sync.Once
)

type userController struct {
	g.BaseController[model.User]
	userService service.UserService
}

func init() {
	once.Do(func() {
		UserController = &userController{userService: service.NewUserService()}
	})
}
