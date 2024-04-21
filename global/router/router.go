// Package router 初始化 gin.router 对象
package router

import (
	"blog/pkg/errcode"
	"blog/pkg/response"
	"blog/routes"
	"github.com/gin-gonic/gin"
	"strings"
)

var Router = gin.New()

func InitializeRouter() {
	SetupRouter(Router)
}

func SetupRouter(router *gin.Engine) {
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
		// 如果请求页面是 text/html 的话
		if strings.Contains(acceptString, "text/html") {
			c.String(404, "页面返回404")
		} else {
			response.NewResponse(c, errcode.ErrNotFound.ParseCode()).WithResponse()
		}
	})
}
