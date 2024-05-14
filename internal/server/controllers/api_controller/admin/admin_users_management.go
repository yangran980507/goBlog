// Package admin 存放 admin 对 user 的操作
package admin

import (
	"blog/internal/server/models/user"
	"blog/pkg/errcode"
	"blog/pkg/response"
	"github.com/gin-gonic/gin"
)

// ShowUsers 显示所有用户
func (ac *AdminController) ShowUsers(c *gin.Context) {
	users := make([]user.User, 10)
	users, page := user.Paginate(c, 6)

	response.NewResponse(c, errcode.ErrSuccess.ParseCode()).WithResponse(gin.H{
		"users": users,
		"page":  page,
	})
}
