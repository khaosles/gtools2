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

type UserImpl struct {
	g.AbstractMapper[model.User]
}

func NewUserImpl(db *gorm.DB) *UserImpl {
	var user UserImpl
	user.DB = db
	return &user
}

func (srv UserImpl) Get() {

}
