// Package middlewares 验证是否授权
package middlewares

import (
	"blog/internal/server/models/user"
	"blog/pkg/errcode"
	"blog/pkg/jwt"
	"blog/pkg/response"
	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := jwt.NewJWT().ParseToken(c)

		if err != nil {
			response.NewResponse(c, errcode.ErrTokenInvalid.ParseCode()).
				WithResponse("请登录后操作")
			return
		}

		userModel := user.GetUserInfo(claims.UserLoginName)

		if userModel.ID == 0 {
			response.NewResponse(c, errcode.ErrTokenInvalid.ParseCode()).
				WithResponse("账号异常，请重新登录")
			return
		}

		c.Set("current_user", userModel)

		c.Next()
	}
}
