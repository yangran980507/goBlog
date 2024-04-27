// Package mysql 存放 gorm.db 对象
package mysql

import (
	"blog/global"
	"blog/internal/server/models/user"
	"database/sql"
	"fmt"
	"time"

	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

// DB 数据库对象
var DB *gorm.DB
var SqlDB *sql.DB

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
	SqlDB.SetMaxOpenConns(global.MysqlSetting.MaxOpenConns)

	// 设置最大空闲连接数
	SqlDB.SetMaxIdleConns(global.MysqlSetting.MaxIdleConns)

	// 设置连接过期时间
	SqlDB.SetConnMaxLifetime(global.MysqlSetting.ConnMaxLifeTime * time.Second)

	DB.AutoMigrate(&user.User{})
}
