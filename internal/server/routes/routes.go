// Package routes 注册路由
package routes

import (
	authServer "blog/internal/server/controllers/api_controller/auth"
	"blog/internal/server/controllers/html_controller"
	"blog/internal/server/middlewares"
	"github.com/gin-gonic/gin"
)

// RegisterAPIRoutes 注册网页相关路由
func RegisterAPIRoutes(router *gin.Engine) {

	// 渲染路由组
	htmlGroup := router.Group("")
	{
		cliGroup := htmlGroup.Group("/bookstore")
		{
			// 静态资源控制器
			hc := new(html_controller.HtmlController)

			// 应用主页
			cliGroup.GET("/home", hc.HomePage)

			// 用户注册
			cliGroup.GET("/signup", hc.SignupPage)

			// 图书分类
			cliGroup.GET("/books-categories", hc.BooksCategories)

			// 图书信息
			cliGroup.GET("/books-messages", hc.BooksMessages)

			// 定单查询
			cliGroup.GET("/order-query", hc.OrdersQuery)

			// 投票结果
			cliGroup.GET("/votes", hc.VotingResults)

			// 收银台
			cliGroup.GET("/cashier", hc.Cashier)

			// 新书上架
			cliGroup.GET("/new-books", hc.NewBooks)

			// 查询图书
			cliGroup.GET("/books-query", hc.BooksQuery)

			// 空购物车
			cliGroup.GET("/empty-shopping", hc.EmptyShoppingCart)

			// 购物车
			cliGroup.GET("/shopping-cart", hc.ShoppingCart)

			// 销售排行
			cliGroup.GET("/sales-rank", hc.SalesRank)
		}

	}

	// 接口路由组
	apiGroup := router.Group("/api")
	{
		// 授权相关路由组
		auth := apiGroup.Group("/auth")

		{
			// 用户注册控制实例
			sc := new(authServer.SignupController)
			// 获取用户注册页面
			auth.POST("/signup", middlewares.GuestAuth(), sc.SignupUser)

			// 用户登陆控制实例
			lc := new(authServer.LoginController)
			// 获取用户注册页面
			auth.POST("/login", middlewares.GuestAuth(), lc.LoginUser)
			// 刷新令牌
			auth.POST("/login/refresh-token", middlewares.JWTAuth(), lc.RefreshToken)
		}
	}
}
