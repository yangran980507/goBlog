// Package global 存放全局配置参数对象
package global

import (
	"blog/pkg/config"
	"blog/pkg/console"
)

// ServerSetting 全局变量
var (
	ServerSetting *config.ServerSection
	MysqlSetting  *config.MysqlSection
	LogSetting    *config.LogSection
	AppSetting    *config.AppSection
)

func InitializeConf() {
	err := setupSetting()
	if err != nil {
		console.Exit("init.setupSetting failed,err:" + err.Error())
	}
}

func setupSetting() error {

	// 新建 viper 实例
	setting, err := config.NewSetting()
	if err != nil {
		return err
	}

	// 调用 ReadSection 方法将配置文件写入全局变量中
	// 服务端配置
	err = setting.ReadSection("Server", &ServerSetting)
	if err != nil {
		return err
	}

	// mysql配置
	err = setting.ReadSection("MySQL", &MysqlSetting)
	if err != nil {
		return err
	}

	// 日志配置
	err = setting.ReadSection("Logger", &LogSetting)
	if err != nil {
		return err
	}

	// 应用配置
	err = setting.ReadSection("App", &AppSetting)
	if err != nil {
		return err
	}

	return nil
}
