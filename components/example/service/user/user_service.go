package suser

import (
	muser "github.com/khaosles/gtools2/components/example/mapper/user"
	"github.com/khaosles/gtools2/components/example/model/user"
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
	g.Service[euser.User]
}

func NewUserService() UserService {
	if userServiceInstance == nil {
		userServiceInstance = internal.NewUserService(muser.NewUserMapper())
	}
	return userServiceInstance
}
