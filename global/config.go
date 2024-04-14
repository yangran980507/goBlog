package global

import (
	"blog/pkg/config"
	"log"
)

var (
	ServerSetting *config.ServerSection
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting failed,err: %s", err.Error())
	}
}

func setupSetting() error {
	setting, err := config.NewSetting()
	if err != nil {
		return err
	}

	//ServerSetting.Values = make(map[string]interface{})

	err = setting.ReadSection("Server", &ServerSetting)
	if err != nil {
		return err
	}
	return nil
}
