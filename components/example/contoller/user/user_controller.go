package cuser

import (
	suser "github.com/khaosles/gtools2/components/example/service/user"
)

/*
   @File: user_controller.go
   @Author: khaosles
   @Time: 2023/6/13 10:07
   @Desc:
*/

var UserController = new(userController)

type userController struct {
	userService suser.UserService
}

func init() {
	if UserController == nil {
		UserController = &userController{userService: suser.NewUserService()}
	}
}
