package internal

import (
	"github.com/khaosles/gtools2/components/example/model"
	"github.com/khaosles/gtools2/components/g"
	"gorm.io/gorm"
)

/*
   @File: user_mapper.go
   @Author: khaosles
   @Time: 2023/6/12 10:57
   @Desc:
*/

type UserMapper struct {
	g.BaseMapper[model.User]
}

func NewUserMapper(db *gorm.DB) *UserMapper {
	var userMapper UserMapper
	userMapper.DB = db
	return &userMapper
}
