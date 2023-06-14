package mapper

import (
	"sync"

	"github.com/khaosles/gtools2/components/example/mapper/internal"
	"github.com/khaosles/gtools2/components/example/model"
	"github.com/khaosles/gtools2/components/g"
)

/*
   @File: user_mapper.go
   @Author: khaosles
   @Time: 2023/6/12 10:58
   @Desc:
*/

var (
	userMapperInstance UserMapper
	once               sync.Once
)

type UserMapper interface {
	g.Mapper[model.User]
}

func NewUserMapper() UserMapper {
	once.Do(func() {
		if userMapperInstance == nil {
			userMapperInstance = internal.NewUserMapper(NewDB())
		}
	})
	return userMapperInstance
}
