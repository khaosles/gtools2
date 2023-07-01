package impl

import (
	"github.com/khaosles/gtools2/components/example/mapper"
	"github.com/khaosles/gtools2/components/example/model"
	"github.com/khaosles/gtools2/components/example/service"
	"github.com/khaosles/gtools2/components/g"
)

/*
   @File: user_service.go
   @Author: khaosles
   @Time: 2023/6/12 18:13
   @Desc:
*/

type UserServiceImpl struct {
	g.BaseService[model.User]
}

func init() {
	service.RegisterUserService(NewUserService(mapper.GetUserMapperInstance()))
}

func NewUserService(userMapper mapper.UserMapper) *UserServiceImpl {
	var userImpl UserServiceImpl
	userImpl.Mapper = userMapper
	return &userImpl
}
