// Package config 设置配置信息/viper初始化/封装
package config

import (
	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

// NewSetting 初始化 *Setting 实例
func NewSetting() (*Setting, error) {

	vp := viper.New()

	vp.AddConfigPath("configs/") //添加配置文件路径

	vp.SetConfigName("config") //指定配置文件名

	vp.SetConfigType("yaml") //指定配置文件类型

	err := vp.ReadInConfig() //读取配置文件
	if err != nil {
		return nil, err
	}
	vp.WatchConfig()
	return &Setting{vp}, nil
}

// ReadSection 读取配置文件至配置存储单元
func (s *Setting) ReadSection(key string, val any) error {
	// 反序列化到结构体中
	err := s.vp.UnmarshalKey(key, val)
	if err != nil {
		return err
	}
	return nil
}
