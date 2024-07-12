// Package middlewares 验证是否授权
package middlewares

import (
	"blog/internal/server/models/user"
	"blog/pkg/errcode"
	"blog/pkg/jwt"
	"blog/pkg/response"
	"github.com/gin-gonic/gin"
)

// JWTAuth 对需要鉴权的接口使用
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := jwt.NewJWT().ParseToken(c)

		if err != nil {
			response.NewResponse(c, errcode.ErrTokenInvalid).
				WithResponse("token不存在或无效")
			c.Abort()
			return
		}

		userModel := user.GetUserInfo(claims.UserLoginName)

		c.Set("current_user", userModel)
	}
}
