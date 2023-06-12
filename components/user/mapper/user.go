package mapper

import (
	"github.com/khaosles/gtools2/components/g"
	"github.com/khaosles/gtools2/components/user/mapper/internal"
	"github.com/khaosles/gtools2/components/user/model"
)

/*
   @File: user.go
   @Author: khaosles
   @Time: 2023/6/12 10:58
   @Desc:
*/

// var wireSet = wire.NewSet(
// 	NewDB,
// 	wire.Bind(new(User), new(*internal.UserImpl)),
// 	internal.NewUserImpl,
// )

type User interface {
	g.Mapper[model.User]
}

//	func NewUser() User {
//		wire.Build(wireSet)
//		return &internal.UserImpl{}
//	}
func NewUser() User {
	db := NewDB()
	userImpl := internal.NewUserImpl(db)
	return userImpl
}
