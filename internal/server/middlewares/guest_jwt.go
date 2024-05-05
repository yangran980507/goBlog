// Package middlewares 非用户使用的接口验证
package middlewares

import (
	"blog/pkg/errcode"
	"blog/pkg/jwt"
	"blog/pkg/response"
	"github.com/gin-gonic/gin"
)

func GuestAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		if _, err := jwt.NewJWT().ParseToken(c); err == nil {
			response.NewResponse(c, errcode.ErrTokenInvalid.ParseCode()).WithResponse("请先退出登陆")
			return
		}

		c.Next()
	}
}
