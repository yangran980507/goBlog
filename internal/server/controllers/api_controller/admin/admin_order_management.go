// Package admin 管理员订单管理
package admin

import (
	"blog/internal/server/models/book"
	"blog/internal/server/models/order"
	"blog/pkg/errcode"
	"blog/pkg/response"
	"github.com/gin-gonic/gin"
)

// GetOrders 获取订单
func (ac *AdminController) GetOrders(c *gin.Context) {

	orders := order.OrdersGetAll()

	if len(orders) == 0 {
		// 无订单
		response.NewResponse(c, errcode.ErrEmptyValue).
			WithResponse("empty orders")
	} else {
		response.NewResponse(c, errcode.ErrSuccess).WithResponse(gin.H{
			"orders": orders,
		})
	}
}

// ShowOrdersDetail 订单详细显示
func (ac *AdminController) ShowOrdersDetail(c *gin.Context) {
	// 获取 detailID
	orderDetailID := c.Param("detailID")

	// 获取订单
	orderDetail := order.OrdersDetailGet(orderDetailID)
	bookModel, row := book.GetByID(orderDetail.BookID)

	if row == 0 {
		// 无订单
		response.NewResponse(c, errcode.ErrBookHadRemoved).
			WithResponse("图书已下架")
	} else {
		response.NewResponse(c, errcode.ErrSuccess).WithResponse(gin.H{
			"bookAmount": orderDetail,
			"book":       bookModel,
		})
	}
}
