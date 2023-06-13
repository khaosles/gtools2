package internal

import (
	"gorm.io/gorm"

	"github.com/khaosles/gtools2/components/example/model/user"
	"github.com/khaosles/gtools2/components/g"
)

/*
   @File: user_mapper.go
   @Author: khaosles
   @Time: 2023/6/12 10:57
   @Desc:
*/

type UserMapper struct {
	g.AbstractMapper[euser.User]
}

func NewUserMapper(db *gorm.DB) *UserMapper {
	var userMapper UserMapper
	userMapper.DB = db
	return &userMapper
}
