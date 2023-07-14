package mysql

import (
	"log"
	"time"

	"github.com/khaosles/gtools2/components/g/internal"
	"github.com/khaosles/gtools2/core/config"
	glog "github.com/khaosles/gtools2/core/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
   @File: mysql.go
   @Author: khaosles
   @Time: 2023/4/15 22:32
   @Desc:
*/

func NewMysql(cfg *config.Mysql) *gorm.DB {
	var err error
	var mydb *gorm.DB
	mysqlConfig := mysql.Config{
		DSN:                       cfg.Dsn(), // DSN data source name
		DefaultStringSize:         256,       // string 类型字段的默认长度
		SkipInitializeWithVersion: true,      // 根据版本自动配置
	}
	if mydb, err = gorm.Open(mysql.New(mysqlConfig), internal.Gorm.Config(cfg.Prefix, cfg.Singular, cfg.LogMode, cfg.LogZap)); err != nil {
		log.Fatal("Database connection failed -> ", cfg.DsnHide())
		return nil
	} else {
		mydb.InstanceSet("gorm:table_options", "ENGINE="+cfg.Engine)
		sqlDB, _ := mydb.DB()
		sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
		sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
		sqlDB.SetConnMaxLifetime(time.Minute * time.Duration(cfg.MaxLifeTime))
		glog.Info("Database connection successful...")
	}
	return mydb
}
