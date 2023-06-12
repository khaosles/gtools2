package service

import (
	"github.com/khaosles/gtools2/components/g"
)

/*
   @File: user.go
   @Author: khaosles
   @Time: 2023/6/12 18:12
   @Desc:
*/

type User interface {
	g.Service[g.Model]
}
