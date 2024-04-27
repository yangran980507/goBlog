package database

import (
	"blog/global"
	blogmysql "blog/pkg/mysql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

func InitializeDB() {
	blogmysql.Connect(newDBConfig(), gormlogger.Default.LogMode(gormlogger.Info))
}

func newDBConfig() gorm.Dialector {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?"+
		"charset=utf8mb4&parseTime=True&multiStatements=true&loc=Local",
		global.MysqlSetting.UserName,
		global.MysqlSetting.Password,
		global.MysqlSetting.Host,
		global.MysqlSetting.Port,
		global.MysqlSetting.DBName)

	return mysql.New(mysql.Config{
		DSN: dsn,
	})
}
