package internal

import (
	"github.com/khaosles/gtools2/components/g"
)

/*
   @File: user.go
   @Author: khaosles
   @Time: 2023/6/12 18:13
   @Desc:
*/

type UserImpl struct {
	g.AbstractService[g.Model]
}

func NewUserImpl() *UserImpl {
	var userImpl UserImpl
	// userImpl.Mpr = mapper.UserMapper
	return &userImpl
}
