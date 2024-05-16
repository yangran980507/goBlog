// Package admin 存放 admin 对 user 的操作
package admin

import (
	"blog/internal/server/models/user"
	"blog/internal/server/requests"
	"blog/pkg/errcode"
	"blog/pkg/logger"
	"blog/pkg/response"
	"github.com/gin-gonic/gin"
)

// ShowUsers 显示所有用户
func (ac *AdminController) ShowUsers(c *gin.Context) {
	users := make([]user.User, 6)
	users, page := user.Paginate(c, 6)

	response.NewResponse(c, errcode.ErrSuccess.ParseCode()).WithResponse(gin.H{
		"users": users,
		"page":  page,
	})
}

// ManageFreezeUser 修复用户 freeze 字段信息
func (ac *AdminController) ManageFreezeUser(c *gin.Context) {
	// 创建验证信息变量
	person := requests.Person{}
	if err := c.ShouldBind(&person); err != nil {
		// 绑定验证失败
		logger.LogIf(err)
		response.NewResponse(c, errcode.ErrBind.ParseCode()).
			WithResponse("请求失败，请稍后再试")
		return
	}

	userModel := user.User{
		LoginName: person.LoginName,
		Freeze:    person.IsFreezing,
	}

	err := userModel.MemberFreezeUpdate()
	if err != nil {
		// 更新失败
		logger.LogIf(err)
		response.NewResponse(c, errcode.ErrUnknown.ParseCode()).
			WithResponse("更新失败，请稍后再试")
		return
	}

	ac.ShowUsers(c)
}
