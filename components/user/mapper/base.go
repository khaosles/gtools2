package mapper

import (
	"gorm.io/gorm"

	"github.com/khaosles/gtools2/components/g/pgsql"
)

/*
   @File: base.go
   @Author: khaosles
   @Time: 2023/6/12 11:20
   @Desc:
*/

func NewDB() *gorm.DB {
	return pgsql.DB
}
