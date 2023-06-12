package service

import (
	"github.com/khaosles/gtools2/components/g"
	"github.com/khaosles/gtools2/components/user/mapper"
	"github.com/khaosles/gtools2/components/user/model"
	"github.com/khaosles/gtools2/components/user/service/internal"
)

/*
   @File: user.go
   @Author: khaosles
   @Time: 2023/6/12 18:12
   @Desc:
*/

var userServiceInstance UserService

type UserService interface {
	g.Service[model.User]
}

func NewUserService() UserService {
	if userServiceInstance == nil {
		userServiceInstance = internal.NewUserServiceImpl(mapper.NewUserMapper())
	}
	return userServiceInstance
}
