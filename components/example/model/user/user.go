package euser

import (
	"github.com/khaosles/gtools2/components/g"
)

/*
   @File: user.go
   @Author: khaosles
   @Time: 2023/6/12 10:56
   @Desc:
*/

type User struct {
	g.Model
	UserName string
}

func (User) TableName() string {
	return "user"
}
