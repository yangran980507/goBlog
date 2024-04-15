// Package global 存放全局 gorm.db 对象
package global

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"time"

	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

// DB 数据库对象
var DB *gorm.DB
var SqlDB *sql.DB

func init() {
	Connect(newDBConfig(), gormlogger.Default.LogMode(gormlogger.Info))
}

// Connect 连接数据库
func Connect(dbConfig gorm.Dialector, _logger gormlogger.Interface) {

	var err error
	DB, err = gorm.Open(dbConfig, &gorm.Config{
		Logger: _logger,
	})
	// 错误处理
	if err != nil {
		fmt.Println(err.Error())
	}

	// 获取 sql.db
	SqlDB, err = DB.DB()
	if err != nil {
		fmt.Println(err.Error())
	}

	// 设置最大连接数
	SqlDB.SetMaxOpenConns(MysqlSetting.MaxOpenConns)

	// 设置最大空闲连接数
	SqlDB.SetMaxIdleConns(MysqlSetting.MaxIdleConns)

	// 设置连接过期时间
	SqlDB.SetConnMaxLifetime(MysqlSetting.ConnMaxLifeTime * time.Second)
}

func newDBConfig() gorm.Dialector {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?"+
		"charset=utf8mb4&parseTime=True&multiStatements=true&loc=Local",
		MysqlSetting.UserName,
		MysqlSetting.Password,
		MysqlSetting.Host,
		MysqlSetting.Port,
		MysqlSetting.DBName)

	return mysql.New(mysql.Config{
		DSN: dsn,
	})
}
