// Package routes 注册路由
package routes

import (
	adminServer "blog/internal/server/controllers/api_controller/admin"
	cliServer "blog/internal/server/controllers/api_controller/client"
	"blog/internal/server/middlewares"
	"github.com/gin-gonic/gin"
)

// RegisterAPIRoutes 注册网页相关路由
func RegisterAPIRoutes(router *gin.Engine) {

	// 接口路由组
	apiGroup := router.Group("/api")
	{

		// 授权路由组
		client := apiGroup.Group("/client")
		{
			// 用户端路由控制实例
			uc := new(cliServer.UserController)
			// 用户鉴权路由组
			auth := client.Group("/auth")
			{
				// 用户注册
				auth.POST("/signup", middlewares.GuestAuth(), uc.SignupUser)
				// 用户登陆
				auth.POST("/login", middlewares.GuestAuth(), uc.LoginUser)
				// 刷新令牌
				auth.POST("/login/refresh-token", uc.RefreshToken)
			}

			// 获取图书
			collection := client.Group("/books")
			{
				// 通过分类获取图书
				collection.GET("/by-category", uc.GetBookByCategory)
				// 通过新书获取图书
				collection.GET("/by-is_new_book/:count", uc.GetBookByIsNewBook)
				// 通过销售排行获取图书
				collection.GET("/by-sold/:count", uc.GetBookBySold)
				// 通过推荐获取图书
				collection.GET("/by-recommended/:count", uc.GetBookByRecommended)
				// 通过书名获取图书
				collection.GET("/:book_name", uc.GetBookByName)
				// 通过搜索获取图书
				collection.GET("/search/:book_name", uc.GetBookBySearch)
			}

			// 购物车相关
			cart := client.Group("/carts")
			cart.Use(middlewares.JWTAuth())
			{
				// 显示购物车信息
				cart.GET("", uc.ShowCarts)
				// 加入购物车
				cart.POST("/add/:id", uc.AddIntoCarts)
				// 删除购物车中图书
				cart.POST("/remove", uc.RemoveFromCarts)
				// 清空购物车
				cart.DELETE("/flush", uc.FlushCarts)
			}

			// 订单相关
			order := client.Group("/orders")
			order.Use(middlewares.JWTAuth())
			{
				// 提交订单
				order.POST("/submit", uc.OrdersSubmit)
				// 查看订单
				order.GET("", uc.ShowOrders)
				// 查看订单详细
				order.GET("/detail/:detailID", uc.ShowOrdersDetail)
				// 订单取消
				order.POST("/refund", uc.OrderRefund)

			}

			// 公告相关
			notice := client.Group("/notices")
			{
				// 显示公告信息
				notice.GET("", uc.NoticeGet)
			}

			// 投票相关
			poll := client.Group("/polls")
			{
				// 投票
				poll.PUT("/vote", middlewares.JWTAuth(), middlewares.ExecuteAuth(), uc.IncrByPoll)
				// 显示投票结果
				poll.GET("", uc.GetPoll)
				// 显示投票项
				// poll.GET("keys", uc.GetPollOption)
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
				// 所有图书
				bookManage.GET("", ac.GetBooksAllByPaginator)
				// 删除图书
				bookManage.DELETE("/delete/:id", ac.DeleteBook)
				// 修改图书
				bookManage.PUT("/update/:id", ac.BookUpdate)
				// 单册图书信息
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
				userManage.PUT("/manage-freeze/:id", ac.ManageFreezeUser)
			}

			// 公告管理路由组
			noticeManage := admin.Group("/notices")
			noticeManage.Use(middlewares.JWTAuth(), middlewares.AdminAuth())
			{
				// 显示公告
				noticeManage.GET("", ac.NoticeGet)
				// 添加公告
				noticeManage.POST("/notice-release", ac.NoticeCreate)
				// 删除公告
				noticeManage.DELETE("/delete/:id", ac.NoticeDelete)
			}

			pollManage := admin.Group("/polls")
			pollManage.Use(middlewares.JWTAuth(), middlewares.AdminAuth())
			{
				// 显示投票项，投票数
				pollManage.GET("", ac.GetPoll)
				// 获取分类项
				pollManage.GET("/options", ac.GetCategory)
				// 设置投票项
				pollManage.POST("/create", ac.SetPoll)
				// 删除投票项
				pollManage.DELETE("/:option_name/delete", ac.DeletePoll)

			}
		}
	}
}
