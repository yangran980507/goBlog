// Package app 应用信息
package app

import (
	"blog/global"
)

// IsLocal 是否是本地环境
func IsLocal() bool {
	return global.AppSetting.Env == "local"
}

// IsProduction 是否是线上环境
func IsProduction() bool {
	return global.AppSetting.Env == "production"
}
