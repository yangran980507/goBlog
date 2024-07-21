// Package client 用户购物车 handlerFunc
package client

import (
	"blog/internal/server/models/book"
	"blog/internal/server/models/cart"
	"blog/pkg/app"
	"blog/pkg/errcode"
	"blog/pkg/helps"
	"blog/pkg/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// ShowCarts 显示购物车信息
func (uc *UserController) ShowCarts(c *gin.Context) {

	// 当前用户 id
	uid := app.CurrentUser(c)

	cartModel := cart.GetCart(uid)

	bookLength := len(cartModel.BookID)
	if bookLength == 0 {
		response.NewResponse(c, errcode.ErrEmptyValue).
			WithResponse("购物车为空")
		return
	}

	books, _ := book.GetBooksBySlice(cartModel.BookID)

	response.NewResponse(c, errcode.ErrSuccess).WithResponse(gin.H{
		"cart": books})
}

// AddIntoCarts 加入购物车
func (uc *UserController) AddIntoCarts(c *gin.Context) {
	// 获取当前用户 id
	uid := app.CurrentUser(c)

	// 获取购物车信息
	cartModel := cart.GetCart(uid)

	//
	if len(cartModel.BookID) > 50 {
		// 失败返回失败信息
		response.NewResponse(c, errcode.ErrOverMaxCount).
			WithResponse("购物车已达上限")
		return
	}

	bookID := app.GetIDFromAPI(c, "id")

	// 判断图书是否被删除
	// 图书编号实例
	bookModel := book.Book{
		ID: uint(bookID),
	}

	_, row := bookModel.Get()
	if row != 1 {
		response.NewResponse(c, errcode.ErrBookHadRemoved).
			WithResponse("图书已下架")
		return
	}

	// 判断商品是否存在购物车中
	if helps.JudgeElementInSliceExist(bookID, cartModel.BookID) {
		// 商品已在购物车中
		response.NewResponse(c, errcode.ErrBookHadExisted).
			WithResponse("该书已在购物车中")
		return
	} else {
		// 购物车中无此商品
		// 图书切片中添加 book
		cartModel.BookID = append(cartModel.BookID, bookID)
	}
	// 添加更新时间
	cartModel.UpdateTime = time.Now().Unix()

	// 重新存入
	if !cartModel.SetCart(uid) {
		// 失败返回失败信息
		response.NewResponse(c, errcode.ErrServer).
			WithResponse("加购失败，请稍后重试")
		return
	}

	response.NewResponse(c, errcode.ErrSuccess).
		WithResponse("加购成功")

}

// RemoveFromCarts 删除购物车中图书
func (uc *UserController) RemoveFromCarts(c *gin.Context) {
	// 获取当前用户 id
	uid := app.CurrentUser(c)

	// 获取购物车信息
	cartModel := cart.GetCart(uid)

	// 获取请求删除的图书切片
	var books book.Carts
	err := c.ShouldBind(&books)
	if err != nil {
		response.NewResponse(c, errcode.ErrBind).WithResponse("删除失败")
		return
	}

	// 获取请求删除的 ID 切片
	condition := make([]int64, 0)
	for _, v := range books.Books {
		condition = append(condition, int64(v.ID))
	}

	fmt.Println(condition)
	// 换取新的 购物车数据
	cartModel.BookID = helps.GenerateNewSliceByDeleteOldSlice(cartModel.BookID, condition)

	fmt.Println(cartModel.BookID)

	// 删除对应的购物车位号
	cartModel.UpdateTime = time.Now().Unix()

	// 重新存入
	if !cartModel.SetCart(uid) {
		// 失败返回失败信息
		response.NewResponse(c, errcode.ErrServer).
			WithResponse("删除失败，请稍后重试")
		return
	}

	// 删除成功
	response.NewResponse(c, errcode.ErrSuccess, "删除成功").
		WithResponse(gin.H{
			"bookLength": len(cartModel.BookID),
		})
}

// FlushCarts 清空购物车
func (uc *UserController) FlushCarts(c *gin.Context) {
	// 获取当前用户 id
	uid := app.CurrentUser(c)

	if !cart.DelCart(uid) {
		// 失败返回失败信息
		response.NewResponse(c, errcode.ErrServer).
			WithResponse("清空失败，请稍后重试")
		return
	}

	response.NewResponse(c, errcode.ErrSuccess).
		WithResponse("清空成功")
}
