// Package client 用户购物车相关操作
package client

import (
	"blog/internal/server/models/book"
	"blog/internal/server/models/cart"
	"blog/internal/server/models/user"
	"blog/pkg/errcode"
	"blog/pkg/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

// ShowCarts 显示购物车信息
func (uc *UserController) ShowCarts(c *gin.Context) {
	// 获取当前用户
	val := c.MustGet("current_user")
	userModel := val.(user.User)
	// 当前用户 id
	uid := fmt.Sprintf("%v", userModel.ID)

	cartModel := cart.GetCart(uid)
	if len(cartModel.Books) == 0 {
		response.NewResponse(c, errcode.ErrUnknown.ParseCode()).WithResponse([]book.Book{})
	}

	response.NewResponse(c, errcode.ErrSuccess.ParseCode()).WithResponse(gin.H{
		"cart": cartModel.Books,
	})
}
