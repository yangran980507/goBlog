// Package router 初始化 gin.router 对象
package router

import (
	"blog/internal/server/routes"
	"blog/pkg/errcode"
	"blog/pkg/response"
	"github.com/gin-gonic/gin"
	"strings"
)

func SetupRouter(router *gin.Engine) {
	// 匹配模板
	router.LoadHTMLGlob("templates/**/*.html")

	// 注册全局中间件
	setupMiddlewares(router)

	// 注册 API 路由
	routes.RegisterAPIRoutes(router)

	// 注册 404 路由
	setupNoRoute(router)
}

// 全局中间件
func setupMiddlewares(router *gin.Engine) {
	router.Use(
		gin.Logger(),
		gin.Recovery(),
	)
}

// 404NotFound
func setupNoRoute(router *gin.Engine) {
	// 自定义 handleFunc
	router.NoRoute(func(c *gin.Context) {
		// 获取请求头
		acceptString := c.Request.Header.Get("Accept")
		// 如果请求页面是 text/user_view 的话
		if strings.Contains(acceptString, "text/html_controller") {
			c.String(404, "页面返回404")
		} else {
			response.NewResponse(c, errcode.ErrNotFound).WithResponse()
		}
	})
}
