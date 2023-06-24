package model

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
	g.BaseModel
	UserName string `json:"userName,omitempty"`
}

func (User) TableName() string {
	return "user"
}
