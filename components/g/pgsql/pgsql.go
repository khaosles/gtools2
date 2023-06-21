package pgsql

import (
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/khaosles/gtools2/components/g/internal"
	gcfg "github.com/khaosles/gtools2/core/cfg"
	"github.com/khaosles/gtools2/core/cfg/config"
	glog "github.com/khaosles/gtools2/core/log"
)

/*
   @File: pgsql.go
   @Author: khaosles
   @Time: 2023/4/22 23:39
   @Desc:
*/

var (
	pdb  *gorm.DB
	once sync.Once
)

// func init() {
//	var err error
//	cfg := gcfg.GCfg.Pgsql
//	pgsqlConfig := postgres.Config{
//		DSN:                  cfg.Dsn(), // DSN data source name
//		PreferSimpleProtocol: false,
//	}
//	if DB, err = gorm.Open(postgres.New(pgsqlConfig), internal.Gorm.Config(cfg.Prefix, cfg.Singular, cfg.LogMode, cfg.LogZap)); err != nil {
//		log.Fatal("Database connection failed -> ", cfg.DsnHide())
//		return
//	} else {
//		sqlDB, _ := DB.DB()
//		sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
//		sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
//		glog.Debug("Database connection successful...")
//	}
// }

func NewPgsql(cfg *config.Pgsql) *gorm.DB {
	once.Do(func() {
		if cfg == nil {
			cfg = &gcfg.GCfg.Pgsql
		}
		var err error
		pgsqlConfig := postgres.Config{
			DSN:                  cfg.Dsn(), // DSN data source name
			PreferSimpleProtocol: false,
		}
		if pdb, err = gorm.Open(postgres.New(pgsqlConfig), internal.Gorm.Config(cfg.Prefix, cfg.Singular, cfg.LogMode, cfg.LogZap)); err != nil {
			glog.Error("Database connection failed -> ", cfg.DsnHide())
			return
		} else {
			sqlDB, _ := pdb.DB()
			sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
			sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
			sqlDB.SetConnMaxLifetime(time.Minute * time.Duration(cfg.MaxLifeTime))
			glog.Debug("Database connection successful...")
		}
	})
	return pdb
}
