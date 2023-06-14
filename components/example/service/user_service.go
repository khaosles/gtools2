package service

import (
	"sync"

	"github.com/khaosles/gtools2/components/example/mapper"
	"github.com/khaosles/gtools2/components/example/model"
	"github.com/khaosles/gtools2/components/example/service/internal"
	"github.com/khaosles/gtools2/components/g"
)

/*
   @File: user_service.go
   @Author: khaosles
   @Time: 2023/6/12 18:12
   @Desc:
*/

var (
	userServiceInstance UserService
	once                sync.Once
)

type UserService interface {
	g.Service[model.User]
}

func NewUserService() UserService {
	once.Do(func() {
		if userServiceInstance == nil {
			userServiceInstance = internal.NewUserService(mapper.NewUserMapper())
		}
	})
	return userServiceInstance
}
