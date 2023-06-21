package mysql

import (
	"log"
	"time"

	"github.com/khaosles/gtools2/components/g/internal"
	"github.com/khaosles/gtools2/core/cfg/config"
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

var mydb *gorm.DB

// GormMysql 初始化Mysql数据库
//func init() {
//	var err error
//	cfg := gcfg.GCfg.Mysql
//	mysqlConfig := mysql.Config{
//		DSN:                       cfg.Dsn(), // DSN data source name
//		DefaultStringSize:         256,       // string 类型字段的默认长度
//		SkipInitializeWithVersion: true,      // 根据版本自动配置
//	}
//	if DB, err = gorm.Open(mysql.New(mysqlConfig), internal.Gorm.Config(cfg.Prefix, cfg.Singular, cfg.LogMode, cfg.LogZap)); err != nil {
//		log.Fatal("Database connection failed -> ", cfg.DsnHide())
//		return
//	} else {
//		DB.InstanceSet("gorm:table_options", "ENGINE="+cfg.Engine)
//		sqlDB, _ := DB.DB()
//		sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
//		sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
//		glog.Debug("Database connection successful...")
//		return
//	}
//}

func NewMysql(cfg *config.Mysql) *gorm.DB {
	var err error
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
		glog.Debug("Database connection successful...")
		return mydb
	}
}
