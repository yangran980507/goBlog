// Package client 用户购物车相关操作
package client

import (
	"blog/internal/server/models/book"
	"blog/internal/server/models/cart"
	"blog/pkg/app"
	"blog/pkg/errcode"
	"blog/pkg/response"
	"github.com/gin-gonic/gin"
)

// ShowCarts 显示购物车信息
func (uc *UserController) ShowCarts(c *gin.Context) {

	// 当前用户 id
	uid := CurrentUser(c)

	cartModel := cart.GetCart(uid)
	if len(cartModel.BookID) == 0 {
		response.NewResponse(c, errcode.ErrEmptyCart, "购物车为空").
			WithResponse()
		return
	}

	books, _ := book.GetBooksBySlice(cartModel.BookID)

	response.NewResponse(c, errcode.ErrSuccess).WithResponse(gin.H{
		"cart": books,
	})
}

// AddIntoCarts 加入购物车
func (uc *UserController) AddIntoCarts(c *gin.Context) {
	// 获取当前用户 id
	uid := CurrentUser(c)

	// 获取购物车信息
	cartModel := cart.GetCart(uid)

	//
	if len(cartModel.BookID) > 50 {
		// 失败返回失败信息
		response.NewResponse(c, errcode.ErrOverMaxCount, "购物车已达上限").
			WithResponse()
		return
	}

	bookID := app.GetIDFromAPI(c, "id")

	// 图书切片中添加 book
	cartModel.BookID = append(cartModel.BookID, bookID)

	// 重新存入
	if !cartModel.SetCart(uid) {
		// 失败返回失败信息
		response.NewResponse(c, errcode.ErrServer, "加购失败，请稍后重试").
			WithResponse()
		return
	}

	response.NewResponse(c, errcode.ErrSuccess, "加购成功").
		WithResponse()

}

// RemoveFromCarts 删除购物车中图书
func (uc *UserController) RemoveFromCarts(c *gin.Context) {
	// 获取当前用户 id
	uid := CurrentUser(c)

	// 获取购物车信息
	cartModel := cart.GetCart(uid)

	// 获取要删除的购物车位号
	cartID := app.GetIDFromAPI(c, "cart_id")

	// 购物车切片长度
	cartLength := int64(len(cartModel.BookID))
	newCart := make([]int64, cartLength-1)

	switch cartID - 1 {
	case 0:
		newCart = cartModel.BookID[1:]
	case cartLength - 1:
		newCart = cartModel.BookID[:cartLength-1]
	default:
		newCart = append(cartModel.BookID[:cartID-1], cartModel.BookID[cartID:]...)
	}

	cartModel.BookID = newCart

	// 重新存入
	if !cartModel.SetCart(uid) {
		// 失败返回失败信息
		response.NewResponse(c, errcode.ErrServer, "删除失败，请稍后重试").
			WithResponse()
		return
	}

	response.NewResponse(c, errcode.ErrSuccess, "删除成功").
		WithResponse()
}

// FlushCarts 清空购物车
func (uc *UserController) FlushCarts(c *gin.Context) {
	// 获取当前用户 id
	uid := CurrentUser(c)

	if !cart.DelCart(uid) {
		// 失败返回失败信息
		response.NewResponse(c, errcode.ErrServer, "清空失败，请稍后重试").
			WithResponse()
		return
	}

	response.NewResponse(c, errcode.ErrSuccess, "清空成功").
		WithResponse()
}
