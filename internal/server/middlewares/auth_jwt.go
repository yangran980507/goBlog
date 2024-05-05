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
			response.NewResponse(c, errcode.ErrTokenInvalid.ParseCode()).
				WithResponse("请登录后操作")
			c.Abort()
		}

		userModel := user.GetUserInfo(claims.UserLoginName)

		if userModel.ID == 0 {
			response.NewResponse(c, errcode.ErrTokenInvalid.ParseCode()).
				WithResponse("账号异常，请重新登录")
			c.Abort()
		}

		c.Set("current_user", userModel)

	}
}
