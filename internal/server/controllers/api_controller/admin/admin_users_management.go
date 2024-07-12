// Package admin 管理员用户管理 handlerFunc
package admin

import (
	"blog/internal/server/models/user"
	"blog/internal/server/requests"
	"blog/pkg/app"
	"blog/pkg/errcode"
	"blog/pkg/response"
	"github.com/gin-gonic/gin"
)

// ShowUsers 显示所有用户
func (ac *AdminController) ShowUsers(c *gin.Context) {
	users := make([]user.User, 8)
	users, page := user.Paginate(c, "8")

	if len(users) == 0 {
		response.NewResponse(c, errcode.ErrEmptyValue).WithResponse("空数据库")
		return
	}
	response.NewResponse(c, errcode.ErrSuccess).WithResponse(gin.H{
		"users": users,
		"page":  page,
	})
}

// ManageFreezeUser 修复用户 freeze 字段信息
func (ac *AdminController) ManageFreezeUser(c *gin.Context) {
	var (
		userModel user.User
		request   requests.Freeze
	)
	// 解析接口数据
	id := app.GetIDFromAPI(c, "id")

	userModel.ID = uint(id)
	if err := c.ShouldBind(&request); err != nil {
		userModel.Freeze = !request.IsFrozen
	}

	u, err := userModel.MemberFreezeUpdate()
	if err != nil {
		// 更新失败
		response.NewResponse(c, errcode.ErrServer).
			WithResponse("修改失败，请稍后再试")
		return
	}
	response.NewResponse(c, errcode.ErrSuccess, "修改成功").
		WithResponse(u)
}
