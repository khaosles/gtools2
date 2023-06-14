package sqlite

import (
	"log"

	"github.com/khaosles/gtools2/components/g/internal"
	gcfg "github.com/khaosles/gtools2/core/cfg"
	"github.com/khaosles/gtools2/core/cfg/config"
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

var sdb *gorm.DB

//
//func init() {
//	var err error
//	cfg := gcfg.GCfg.Sqlite
//	if DB, err = gorm.Open(sqlite.Open(cfg.Dsn()), internal.Gorm.Config(
//		cfg.Prefix, cfg.Singular, cfg.LogMode, cfg.LogZap),
//	); err != nil {
//		log.Fatal("Database connection failed=> ", cfg.Dsn())
//		return
//	} else {
//		sqlDB, _ := DB.DB()
//		sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
//		sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
//		glog.Debug("数据库连接成功...")
//
//	}
//}

func NewSqlite(cfg *config.Sqlite) *gorm.DB {
	if cfg == nil {
		cfg = &gcfg.GCfg.Sqlite
	}
	var err error
	if sdb, err = gorm.Open(sqlite.Open(cfg.Dsn()), internal.Gorm.Config(
		cfg.Prefix, cfg.Singular, cfg.LogMode, cfg.LogZap),
	); err != nil {
		log.Fatal("Database connection failed=> ", cfg.Dsn())
		return nil
	} else {
		sqlDB, _ := sdb.DB()
		sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
		sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
		glog.Debug("Database connection successful...")
	}
	return sdb
}
