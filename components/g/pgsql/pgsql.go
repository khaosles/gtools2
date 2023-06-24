package pgsql

import (
	"time"

	"github.com/khaosles/gtools2/core/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/khaosles/gtools2/components/g/internal"
	glog "github.com/khaosles/gtools2/core/log"
)

/*
   @File: pgsql.go
   @Author: khaosles
   @Time: 2023/4/22 23:39
   @Desc:
*/

func NewPgsql(cfg *config.Pgsql) *gorm.DB {
	var err error
	var pdb *gorm.DB
	pgsqlConfig := postgres.Config{
		DSN:                  cfg.Dsn(), // DSN data source name
		PreferSimpleProtocol: false,
	}
	if pdb, err = gorm.Open(postgres.New(pgsqlConfig), internal.Gorm.Config(cfg.Prefix, cfg.Singular, cfg.LogMode, cfg.LogZap)); err != nil {
		glog.Error("Database connection failed -> ", cfg.DsnHide())
		return nil
	} else {
		sqlDB, _ := pdb.DB()
		sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
		sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
		sqlDB.SetConnMaxLifetime(time.Minute * time.Duration(cfg.MaxLifeTime))
		glog.Debug("Database connection successful...")
	}
	return pdb
}
