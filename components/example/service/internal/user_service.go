package internal

import (
	"github.com/khaosles/gtools2/components/example/mapper"
	"github.com/khaosles/gtools2/components/example/model"
	"github.com/khaosles/gtools2/components/g"
)

/*
   @File: user_service.go
   @Author: khaosles
   @Time: 2023/6/12 18:13
   @Desc:
*/

type UserService struct {
	g.BaseService[model.User]
}

func NewUserService(userMapper mapper.UserMapper) *UserService {
	var userImpl UserService
	userImpl.Mapper = userMapper
	return &userImpl
}
