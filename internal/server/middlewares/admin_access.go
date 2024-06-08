// Package middlewares admin 接口验证
package middlewares

import (
	"blog/internal/server/models/user"
	"blog/pkg/errcode"
	"blog/pkg/response"
	"github.com/gin-gonic/gin"
)

// AdminAuth admin 身份验证
func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		userModel := c.MustGet("current_user").(user.User)

		if !userModel.IsManager {
			// 不是 管理员份身份
			response.NewResponse(c, errcode.ErrNotAdmin, "非管理员身份没有该权限").
				WithResponse()
			c.Abort()
		}
	}
}
