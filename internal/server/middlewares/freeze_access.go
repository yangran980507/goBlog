// Package middlewares 用户权限冻结验证
package middlewares

import (
	"blog/internal/server/models/user"
	"blog/pkg/errcode"
	"blog/pkg/response"
	"github.com/gin-gonic/gin"
)

// FreezeAuth client 冻结验证
func FreezeAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		userModel := c.MustGet("current_user").(user.User)

		if userModel.Freeze {
			// 用户被冻结
			response.NewResponse(c, errcode.ErrFrozen).
				WithResponse("权限被冻结")
			c.Abort()
		}
	}
}
