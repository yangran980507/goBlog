// Package admin 存放 admin 对 user 的操作
package admin

import (
	"blog/internal/server/models/user"
	"blog/pkg/errcode"
	"blog/pkg/logger"
	"blog/pkg/response"
	"github.com/gin-gonic/gin"
)

// ShowUsers 显示所有用户
func (ac *AdminController) ShowUsers(c *gin.Context) {
	users := make([]user.User, 10)
	users, err := user.GetUsers()
	if err != nil {
		logger.LogIf(err)
		response.NewResponse(c, errcode.ErrUnknown.ParseCode()).WithResponse()
	}
	response.NewResponse(c, errcode.ErrSuccess.ParseCode()).WithResponse(gin.H{
		"users": users,
	})
}
