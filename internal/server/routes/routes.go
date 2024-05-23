// Package routes 注册路由
package routes

import (
	adminServer "blog/internal/server/controllers/api_controller/admin"
	cliServer "blog/internal/server/controllers/api_controller/client"
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
			cliGroup.GET("/book-categories", hc.BooksCategories)
			// 图书信息
			cliGroup.GET("/book-messages", hc.BooksMessages)
			// 定单查询
			cliGroup.GET("/order-query", hc.OrdersQuery)
			// 投票结果
			cliGroup.GET("/votes", hc.VotingResults)
			// 收银台
			cliGroup.GET("/cashier", hc.Cashier)
			// 新书上架
			cliGroup.GET("/new-book", hc.NewBooks)
			// 查询图书
			cliGroup.GET("/book-query", hc.BooksQuery)

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
		// 用户端路由控制实例
		uc := new(cliServer.UserController)
		// 授权路由组
		client := apiGroup.Group("/client")
		{
			// 用户鉴权路由组
			auth := client.Group("/auth")
			{
				// 获取用户注册页面
				auth.POST("/signup", middlewares.GuestAuth(), uc.SignupUser)
				// 获取用户登陆页面
				auth.POST("/login", middlewares.GuestAuth(), uc.LoginUser)
				// 刷新令牌
				auth.POST("/login/refresh-token",
					middlewares.JWTAuth(), uc.RefreshToken)
			}

			// 获取图书
			collection := client.Group("")
			{
				// 通过分类获取图书
				collection.GET("get-book-by-category", uc.GetBookByCategory)
			}
		}

		// 管理员路由组
		admin := apiGroup.Group("/admin")
		{
			ac := new(adminServer.AdminController)

			// 管理员鉴权路由组
			auth := admin.Group("/auth")
			{
				// 登录 admin
				auth.POST("/login", ac.LoginAdmin)
			}

			// 图书管理路由组
			bookManage := admin.Group("/books")
			// 鉴权中间件
			bookManage.Use(middlewares.JWTAuth(), middlewares.AdminAuth())
			{
				// 添加图书
				bookManage.POST("/book-storage", ac.BookStorage)
				// 图书信息
				bookManage.GET("", ac.GetBooksAllByPaginator)
				// 删除图书
				bookManage.DELETE("/delete/:id", ac.DeleteBook)
				// 修改图书
				bookManage.PUT("update/:id", ac.BookUpdate)
				//
				bookManage.GET("/:id", ac.GetBook)

			}

			// 用户管理路由组
			userManage := admin.Group("/users")
			// 鉴权中间件
			userManage.Use(middlewares.JWTAuth(), middlewares.AdminAuth())
			{
				// 显示用户
				userManage.GET("", ac.ShowUsers)
				// 冻结/解冻用户
				userManage.PUT("/manage-freeze", ac.ManageFreezeUser)
			}

			// 公告管理路由组
			noticeManage := admin.Group("/notices")
			{
				// 显示公告
				noticeManage.GET("", ac.NoticeGet)
				// 添加公告
				noticeManage.POST("/notice-release", ac.NoticeCreate)
				// 删除公告
				noticeManage.DELETE("/delete/:id", ac.NoticeDelete)
			}
		}
	}
}
