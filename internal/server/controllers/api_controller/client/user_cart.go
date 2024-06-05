// Package client 用户购物车相关操作
package client

import (
	"blog/internal/server/models/book"
	"blog/internal/server/models/cart"
	"blog/pkg/errcode"
	"blog/pkg/response"
	"github.com/gin-gonic/gin"
)

// ShowCarts 显示购物车信息
func (uc *UserController) ShowCarts(c *gin.Context) {

	// 当前用户 id
	uid := CurrentUser(c)

	cartModel := cart.GetCart(uid)
	if cartModel.Books == nil {
		response.NewResponse(c, errcode.ErrUnknown.ParseCode()).WithResponse()
		return
	}

	if len(cartModel.Books) == 0 {
		response.NewResponse(c, errcode.ErrSuccess.ParseCode()).WithResponse(gin.H{
			"cart": cartModel.Books,
		})
		return
	}
	response.NewResponse(c, errcode.ErrSuccess.ParseCode()).WithResponse(gin.H{
		"cart": cartModel.Books,
	})
}

// AddIntoCarts 加入购物车
func (uc *UserController) AddIntoCarts(c *gin.Context) {
	// 获取当前用户 id
	uid := CurrentUser(c)

	// 获取购物车信息
	cartModel := cart.GetCart(uid)

	// 解析接口数据中 id
	bookModel := book.Book{
		BaseMode: book.GetIDFromAPI(c),
	}

	// 获取图书数据
	bookMes, _ := bookModel.Get()

	// 图书切片中添加 book
	cartModel.Books = append(cartModel.Books, bookMes)

	// 重新存入
	if !cartModel.SetCart(uid) {
		// 失败返回失败信息
		response.NewResponse(c, errcode.ErrUnknown.ParseCode()).
			WithResponse("加入购物车失败，请稍后重试")
	}

	response.NewResponse(c, errcode.ErrSuccess.ParseCode()).WithResponse("添加成功")
	
}
