package mapper

import (
	"github.com/khaosles/gtools2/components/g/pgsql"
	"gorm.io/gorm"
)

/*
   @File: base.go
   @Author: khaosles
   @Time: 2023/6/12 11:20
   @Desc:
*/

func NewDB() *gorm.DB {
	return pgsql.NewPgsql(nil)
}
