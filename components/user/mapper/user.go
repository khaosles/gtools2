package mapper

import (
	"github.com/khaosles/gtools2/components/g"
	"github.com/khaosles/gtools2/components/user/mapper/internal"
	"github.com/khaosles/gtools2/components/user/model"
)

/*
   @File: user.go
   @Author: khaosles
   @Time: 2023/6/12 10:58
   @Desc:
*/

var userMapperInstance UserMapper

type UserMapper interface {
	g.Mapper[model.User]
}

func NewUserMapper() UserMapper {
	if userMapperInstance == nil {
		userMapperInstance = internal.NewUserMapperImpl(NewDB())
	}
	return userMapperInstance
}
