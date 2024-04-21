// Package routes 注册路由
package routes

import (
	authContr "blog/internal/server/controllers/v1/auth"
	"github.com/gin-gonic/gin"
)

// RegisterAPIRoutes 注册网页相关路由
func RegisterAPIRoutes(router *gin.Engine) {

	// v1 版路由组
	v1 := router.Group("/v1")
	{
		// 授权相关路由组
		auth := v1.Group("/auth")
		{
			// 注册相关控制器
			sc := new(authContr.SignupController)

			// route: 验证手机号是否存在
			auth.POST("/signup/phone/exist", sc.IsUserExist)
		}
	}
}
