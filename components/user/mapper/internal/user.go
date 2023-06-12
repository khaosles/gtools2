package internal

import (
	"gorm.io/gorm"

	"github.com/khaosles/gtools2/components/g"
	"github.com/khaosles/gtools2/components/user/model"
)

/*
   @File: user.go
   @Author: khaosles
   @Time: 2023/6/12 10:57
   @Desc:
*/

type UserMapperImpl struct {
	g.AbstractMapper[model.User]
}

func NewUserMapperImpl(db *gorm.DB) *UserMapperImpl {
	var userMapper UserMapperImpl
	userMapper.DB = db
	return &userMapper
}
