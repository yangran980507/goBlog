// Package app 应用信息
package app

import (
	"blog/global"
	"github.com/gin-gonic/gin"
	"strconv"
)

// IsLocal 是否是本地环境
func IsLocal() bool {
	return global.AppSetting.Env == "local"
}

// IsProduction 是否是线上环境
func IsProduction() bool {
	return global.AppSetting.Env == "production"
}

// GetIDFromAPI 获取接口中的 :id
func GetIDFromAPI(c *gin.Context, key string) int64 {
	// id 字符串
	idStr := c.Param(key)

	id, _ := strconv.Atoi(idStr)
	return int64(id)
}
