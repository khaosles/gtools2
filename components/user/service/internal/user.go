package internal

import (
	"github.com/khaosles/gtools2/components/g"
	"github.com/khaosles/gtools2/components/user/mapper"
	"github.com/khaosles/gtools2/components/user/model"
)

/*
   @File: user.go
   @Author: khaosles
   @Time: 2023/6/12 18:13
   @Desc:
*/

type UserServiceImpl struct {
	g.AbstractService[model.User]
}

func NewUserServiceImpl(userMapper mapper.UserMapper) *UserServiceImpl {
	var userServiceImpl UserServiceImpl
	userServiceImpl.Mpr = userMapper
	return &userServiceImpl
}
