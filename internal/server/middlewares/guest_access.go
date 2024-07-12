// Package middlewares 非用户使用的接口验证
package middlewares

import (
	"blog/pkg/errcode"
	"blog/pkg/jwt"
	"blog/pkg/response"
	"github.com/gin-gonic/gin"
)

// GuestAuth 对需要鉴权的请求接口使用
func GuestAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		if _, err := jwt.NewJWT().ParseToken(c); err == nil {
			// 已授权，返回 err messages
			response.NewResponse(c, errcode.ErrTokenInvalid).
				WithResponse("请先退出登陆")
			// 后续 handler 不再执行
			c.Abort()
		}

	}
}
