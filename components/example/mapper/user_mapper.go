package mapper

import (
	"sync"

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
	userMapperOnce     sync.Once // 生成单例对象
)

type UserMapper interface {
	g.Mapper[model.User]
}

func GetUserMapperInstance() UserMapper {
	userMapperOnce.Do(func() {
		mapper := userMapper{}
		//mapper.DB = configures.GetDBInstance()
		userMapperInstance = mapper
	})
	return userMapperInstance
}

type userMapper struct {
	g.BaseMapper[model.User]
}
