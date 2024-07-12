// Package database 存放初始化数据库操作对象
package database

import (
	"blog/global"
	"blog/internal/server/models/book"
	"blog/internal/server/models/notice"
	"blog/internal/server/models/order"
	"blog/internal/server/models/user"
	blogmysql "blog/pkg/mysql"
	"fmt"
	"gorm.io/driver/mysql"
	"time"
)

func InitializeDB() {
	setupDB()
}

func setupDB() {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?"+
		"charset=utf8mb4&parseTime=True&multiStatements=true&loc=Local",
		global.MysqlSetting.UserName,
		global.MysqlSetting.Password,
		global.MysqlSetting.Host,
		global.MysqlSetting.Port,
		global.MysqlSetting.DBName,
	)

	dbConfig := mysql.New(mysql.Config{
		DSN: dsn,
	})

	blogmysql.Connect(dbConfig)

	// 设置最大连接数
	blogmysql.SqlDB.SetMaxOpenConns(global.MysqlSetting.MaxOpenConns)

	// 设置最大空闲连接数
	blogmysql.SqlDB.SetMaxIdleConns(global.MysqlSetting.MaxIdleConns)

	// 设置连接过期时间
	blogmysql.SqlDB.SetConnMaxLifetime(global.MysqlSetting.ConnMaxLifeTime * time.Second)

	//自动迁移至数据库
	blogmysql.DB.AutoMigrate(&user.User{})          // 用户表迁移
	blogmysql.DB.AutoMigrate(&book.Book{})          // 图书表迁移
	blogmysql.DB.AutoMigrate(&book.Category{})      // 图书类别表迁移
	blogmysql.DB.AutoMigrate(&notice.Notice{})      // 公告表迁移
	blogmysql.DB.AutoMigrate(&order.Order{})        // 定单表迁移
	blogmysql.DB.AutoMigrate(&order.OrdersDetail{}) // 定单明细表迁移
}
