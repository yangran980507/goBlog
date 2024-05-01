// Package mysql 数据库对象
package mysql

import (
	"blog/pkg/console"
	"database/sql"
	"fmt"
	gormlogger "gorm.io/gorm/logger"

	"gorm.io/gorm"
)

// DB 数据库对象
var DB *gorm.DB
var SqlDB *sql.DB

// Connect 数据库连接设置
func Connect(dbConfig gorm.Dialector) {

	var err error
	DB, err = gorm.Open(dbConfig, &gorm.Config{
		Logger: gormlogger.Default.LogMode(gormlogger.Info),
	})
	// 错误处理
	if err != nil {
		console.Exit(err.Error())
	}

	// 获取 sql.db
	SqlDB, err = DB.DB()
	if err != nil {
		fmt.Println(err.Error())
	}

}
