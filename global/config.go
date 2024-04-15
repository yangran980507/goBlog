// Package global 存放全局配置参数对象
package global

import (
	"blog/pkg/config"
	"log"
)

// ServerSetting 全局变量
var (
	ServerSetting *config.ServerSection
	MysqlSetting  *config.MysqlSection
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting failed,err: %s", err.Error())
	}
}

func setupSetting() error {

	// 新建 viper 实例
	setting, err := config.NewSetting()
	if err != nil {
		return err
	}

	// 调用 ReadSection 方法将配置文件写入全局变量中
	err = setting.ReadSection("Server", &ServerSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("MySQL", &MysqlSetting)
	if err != nil {
		return err
	}
	return nil
}
