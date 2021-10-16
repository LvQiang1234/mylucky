package dbs

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"mylucky/mylog"
	"time"
)

var (
	db *gorm.DB
	sqlDB *sql.DB
)

func OpenMysqlDB(dbUrl string, config *gorm.Config, maxIdleConns, maxOpenConns int, models ...interface{}) (instance *gorm.DB, err error) {
	if config == nil {
		config = &gorm.Config {}
	}
	if config.NamingStrategy == nil {
		config.NamingStrategy = schema.NamingStrategy{
			TablePrefix: "t_",
			SingularTable: true,
		}
	}
	if db, err = gorm.Open(mysql.Open(dbUrl), config); err != nil {
		mylog.Error("opens mysql database failed: %s", err.Error())
		return
	}
	instance = db
	if sqlDB, err = db.DB(); err == nil {
		sqlDB.SetMaxOpenConns(maxIdleConns)
		sqlDB.SetMaxOpenConns(maxOpenConns)
		sqlDB.SetConnMaxLifetime(time.Hour)
	} else {
		mylog.Error("%v", err)
	}
	if err = db.AutoMigrate(models...); nil != err {
		mylog.Error("auto migate tables failed: %s", err.Error())
	}
	return
}

// DB 获取数据库连接
func DB() *gorm.DB {
	return db
}


