package sqlite

import (
	"github.com/khaosles/gtools2/components/g/internal"
	"github.com/khaosles/gtools2/core/config"
	glog "github.com/khaosles/gtools2/core/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

/*
   @File: sqlite.go
   @Author: khaosles
   @Time: 2023/4/23 01:02
   @Desc:
*/

func NewSqlite(cfg *config.Sqlite) *gorm.DB {
	var err error
	var db *gorm.DB
	if db, err = gorm.Open(sqlite.Open(cfg.Dsn()), internal.Gorm.Config(
		cfg.Prefix, cfg.Singular, cfg.LogMode, cfg.LogZap),
	); err != nil {
		glog.Error("Database connection failed=> ", cfg.Dsn())
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
		sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
		glog.Info("Database connection successful...")
	}
	return db
}
