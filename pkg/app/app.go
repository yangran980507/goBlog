// Package app 应用信息
package app

import (
	"blog/global"
	"blog/internal/server/models/user"
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

// GetStrFromAPI 获取接口中的 :id
func GetStrFromAPI(c *gin.Context, key string) string {
	// id 字符串
	val := c.Param(key)
	return val
}

// CurrentUser 获取当前用户
func CurrentUser(c *gin.Context) string {
	// 获取当前用户
	val := c.MustGet("current_user")
	userModel := val.(user.User)
	return strconv.Itoa(int(userModel.ID))
}
