package internal

import (
	muser "github.com/khaosles/gtools2/components/example/mapper/user"
	"github.com/khaosles/gtools2/components/example/model/user"
	"github.com/khaosles/gtools2/components/g"
)

/*
   @File: user_service.go
   @Author: khaosles
   @Time: 2023/6/12 18:13
   @Desc:
*/

type UserService struct {
	g.AbstractService[euser.User]
}

func NewUserService(userMapper muser.UserMapper) *UserService {
	var userImpl UserService
	userImpl.Mpr = userMapper
	return &userImpl
}
