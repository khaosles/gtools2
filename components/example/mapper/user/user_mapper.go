package muser

import (
	"github.com/khaosles/gtools2/components/example/mapper"
	"github.com/khaosles/gtools2/components/example/mapper/internal"
	"github.com/khaosles/gtools2/components/example/model/user"
	"github.com/khaosles/gtools2/components/g"
)

/*
   @File: user_mapper.go
   @Author: khaosles
   @Time: 2023/6/12 10:58
   @Desc:
*/

var userMapperInstance UserMapper

type UserMapper interface {
	g.Mapper[euser.User]
}

func NewUserMapper() UserMapper {
	if userMapperInstance == nil {
		userMapperInstance = internal.NewUserMapper(mapper.NewDB())
	}
	return userMapperInstance
}
