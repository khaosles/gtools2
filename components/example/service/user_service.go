package service

import (
	"sync"

	"github.com/khaosles/gtools2/components/example/model"
	"github.com/khaosles/gtools2/components/g"
	"github.com/khaosles/gtools2/utils/assert"
)

/*
   @File: user_service.go
   @Author: khaosles
   @Time: 2023/6/12 18:12
   @Desc:
*/

var (
	userServiceInstance UserService
	userServiceOnce     sync.Once
)

type UserService interface {
	g.Service[model.User]
}

func GetUserServiceInstance() UserService {
	assert.IsNotImplemented(userServiceInstance, "UserService not implement.")
	return userServiceInstance
}

func RegisterUserService(userService UserService) {
	userServiceOnce.Do(func() {
		userServiceInstance = userService
	})
}
