// Package client 用户购物车相关操作
package client

import (
	"blog/internal/server/models/book"
	"blog/internal/server/models/cart"
	"blog/pkg/errcode"
	"blog/pkg/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

// ShowCarts 显示购物车信息
func (uc *UserController) ShowCarts(c *gin.Context) {

	// 当前用户 id
	uid := CurrentUser(c)

	cartModel := cart.GetCart(uid)
	if cartModel.Books == nil || len(cartModel.Books) == 0 {
		response.NewResponse(c, errcode.ErrEmptyCart.ParseCode()).WithResponse("空购物车")
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

	//
	if len(cartModel.Books) > 50 {
		// 失败返回失败信息
		response.NewResponse(c, errcode.ErrUnknown.ParseCode()).
			WithResponse("购物车已达上限")
		return
	}

	// 图书切片中添加 book
	cartModel.Books = append(cartModel.Books, bookMes)

	// 重新存入
	if !cartModel.SetCart(uid) {
		// 失败返回失败信息
		response.NewResponse(c, errcode.ErrUnknown.ParseCode()).
			WithResponse("加购失败，请稍后重试")
		return
	}

	response.NewResponse(c, errcode.ErrSuccess.ParseCode()).WithResponse("加购成功")

}

// RemoveFromCarts 删除购物车中图书
func (uc *UserController) RemoveFromCarts(c *gin.Context) {
	// 获取当前用户 id
	uid := CurrentUser(c)

	// 获取购物车信息
	cartModel := cart.GetCart(uid)

	// 获取要删除的购物车位号
	cartID, _ := strconv.Atoi(c.Param("cart_id"))

	// 购物车切片长度
	cartLength := len(cartModel.Books)
	newCart := make([]book.Book, cartLength-1)

	switch cartID - 1 {
	case 0:
		newCart = cartModel.Books[1:]
	case cartLength - 1:
		newCart = cartModel.Books[:cartLength-1]
	default:
		newCart = append(cartModel.Books[:cartID-1], cartModel.Books[cartID:]...)
	}

	cartModel.Books = newCart

	// 重新存入
	if !cartModel.SetCart(uid) {
		// 失败返回失败信息
		response.NewResponse(c, errcode.ErrUnknown.ParseCode()).
			WithResponse("删除失败，请稍后重试")
		return
	}

	response.NewResponse(c, errcode.ErrSuccess.ParseCode()).WithResponse("删除成功")
}

// FlushCarts 清空购物车
func (uc *UserController) FlushCarts(c *gin.Context) {
	// 获取当前用户 id
	uid := CurrentUser(c)

	if !cart.DelCart(uid) {
		// 失败返回失败信息
		response.NewResponse(c, errcode.ErrUnknown.ParseCode()).
			WithResponse("清空失败，请稍后重试")
		return
	}

	response.NewResponse(c, errcode.ErrSuccess.ParseCode()).WithResponse("清空成功")
}
