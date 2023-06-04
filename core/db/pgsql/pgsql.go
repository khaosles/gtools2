package pgsql

import (
	"log"

	gcfg "github.com/khaosles/gtools2/core/cfg"
	"github.com/khaosles/gtools2/core/db/internal"
	glog "github.com/khaosles/gtools2/core/log"
	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

/*
   @File: pgsql.go
   @Author: khaosles
   @Time: 2023/4/22 23:39
   @Desc:
*/

var DB *gorm.DB

func init() {
	var err error
	cfg := gcfg.GCfg.Pgsql
	pgsqlConfig := postgres.Config{
		DSN:                  cfg.Dsn(), // DSN data source name
		PreferSimpleProtocol: false,
	}
	if DB, err = gorm.Open(postgres.New(pgsqlConfig), internal.Gorm.Config(cfg.Prefix, cfg.Singular, cfg.LogMode, cfg.LogZap)); err != nil {
		log.Fatal("Database connection failed=> ", cfg.Dsn())
		return
	} else {
		sqlDB, _ := DB.DB()
		sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
		sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
		glog.Debug("数据库连接成功...")
	}
}
