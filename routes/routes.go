// Package routes 注册路由
package routes

import (
	authContr "blog/internal/server/controllers/client/auth"
	"blog/internal/server/controllers/client/static"
	"github.com/gin-gonic/gin"
)

// RegisterAPIRoutes 注册网页相关路由
func RegisterAPIRoutes(router *gin.Engine) {

	// client 版路由组
	user := router.Group("/bookstore")

	{
		// 应用主页
		// 静态资源控制器
		sc := new(static.StaticController)
		
		user.GET("/home-page", sc.HomePage)

		// 授权相关路由组
		auth := user.Group("/auth")
		{
			// 注册相关控制器
			sc := new(authContr.SignupController)

			// route: 验证手机号是否存在
			auth.GET("/homepage", sc.IsUserExist)
		}
	}
}
