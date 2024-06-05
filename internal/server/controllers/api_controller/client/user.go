// Package client 用户端路由控制器
package client

import (
	"blog/internal/server/controllers"
	"blog/internal/server/models/user"
	"github.com/gin-gonic/gin"
	"strconv"
)

// UserController 用户控制器
type UserController struct {
	controllers.BaseController
}

// CurrentUser 获取当前用户
func CurrentUser(c *gin.Context) string {
	// 获取当前用户
	val := c.MustGet("current_user")
	userModel := val.(user.User)
	return strconv.Itoa(int(userModel.ID))
}
