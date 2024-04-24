// Package html_controller html渲染控制器及其方法
package html_controller

import (
	"blog/internal/server/controllers"
	"github.com/gin-gonic/gin"
)

// HtmlController 页面渲染控制器
type HtmlController struct {
	controllers.BaseController
}

// HomePage 渲染主页
func (hc *HtmlController) HomePage(c *gin.Context) {
	c.HTML(200, "主页.html", nil)
}

// SignupPage 注册页面
func (hc *HtmlController) SignupPage(c *gin.Context) {
	c.HTML(200, "用户注册.html", nil)
}

// BooksMessages 图书信息
func (hc *HtmlController) BooksMessages(c *gin.Context) {
	c.HTML(200, "图书信息.html", nil)
}

// BooksCategories 图书类别
func (hc *HtmlController) BooksCategories(c *gin.Context) {
	c.HTML(200, "图书分类.html", nil)
}

// OrdersQuery 定单查询
func (hc *HtmlController) OrdersQuery(c *gin.Context) {
	c.HTML(200, "定单查询.html", nil)
}

// VotingResults 投票结果
func (hc *HtmlController) VotingResults(c *gin.Context) {
	c.HTML(200, "投票结果.html", nil)
}

// Cashier 收银
func (hc *HtmlController) Cashier(c *gin.Context) {
	c.HTML(200, "收银台.html", nil)
}

// NewBooks 新书上架
func (hc *HtmlController) NewBooks(c *gin.Context) {
	c.HTML(200, "新书上架.html", nil)
}

// BooksQuery 图书查询
func (hc *HtmlController) BooksQuery(c *gin.Context) {
	c.HTML(200, "查询图书.html", nil)
}

// EmptyShoppingCart 空购物车
func (hc *HtmlController) EmptyShoppingCart(c *gin.Context) {
	c.HTML(200, "空购物车.html", nil)
}

// ShoppingCart 购物车
func (hc *HtmlController) ShoppingCart(c *gin.Context) {
	c.HTML(200, "购物车.html", nil)
}

// SalesRank 销售排行
func (hc *HtmlController) SalesRank(c *gin.Context) {
	c.HTML(200, "销售排行.html", nil)
}
