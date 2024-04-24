// Package router 初始化 gin.router 对象
package router

import (
	"blog/pkg/errcode"
	"blog/pkg/response"
	"blog/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
)

func SetupRouter(router *gin.Engine) {
	// 匹配模板
	router.LoadHTMLGlob("templates/**/*")

	// 配置静态文件加载
	//staticAccess(router)

	// 注册全局中间件
	setupMiddlewares(router)

	// 注册 API 路由
	routes.RegisterAPIRoutes(router)

	// 注册 404 路由
	setupNoRoute(router)
}

// 访问静态资源
func staticAccess(router *gin.Engine) {
	// 获取绝对路径
	path, _ := os.Getwd()
	//定义静态文件路径
	fmt.Printf("this is %s\n", path)
	//router.Static("assets", path+"/assets/css")
	router.Static("/assets", path+"/assets")
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
		if strings.Contains(acceptString, "text/html") {
			c.String(404, "页面返回404")
		} else {
			response.NewResponse(c, errcode.ErrNotFound.ParseCode()).WithResponse()
		}
	})
}
