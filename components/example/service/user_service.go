package service

import (
	muser "github.com/khaosles/gtools2/components/example/mapper"
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

var userServiceInstance UserService

type UserService interface {
	g.Service[model.User]
}

func NewUserService() UserService {
	if userServiceInstance == nil {
		userServiceInstance = internal.NewUserService(muser.NewUserMapper())
	}
	return userServiceInstance
}
