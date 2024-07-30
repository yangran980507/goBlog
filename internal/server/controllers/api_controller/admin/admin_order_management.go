// Package admin 管理员订单管理
package admin

import (
	"blog/internal/server/models/book"
	"blog/internal/server/models/order"
	"blog/internal/server/requests"
	"blog/pkg/errcode"
	"blog/pkg/logger"
	"blog/pkg/response"
	"errors"
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

// OrderExecute 订单执行
func (ac *AdminController) OrderExecute(c *gin.Context) {
	// 创建验证结构体空值实例
	request := requests.OrderExecute{}
	// 绑定结构数据到验证结构体中
	err := c.ShouldBind(&request)
	if err != nil {
		logger.LogIf(err)
		response.NewResponse(c, errcode.ErrBind).WithResponse(err)
		return
	}

	orderModel := &order.Order{
		ID:      request.OrderID,
		Enforce: request.Enforce,
	}

	err = orderModel.AdminOrderChange()
	if err != nil {
		if errors.Is(err, errcode.ErrOrderHadExecuted) {
			response.NewResponse(c, errcode.ErrOrderHadExecuted).
				WithResponse(errcode.ErrOrderHadExecuted.Error())
			return
		}
		logger.LogIf(err)
		response.NewResponse(c, errcode.ErrServer).
			WithResponse("服务器错误，请稍后再试")
		return
	}

	response.NewResponse(c, errcode.ErrSuccess).
		WithResponse("状态已修改")

	var (
		detailModel order.OrdersDetail
	)

	if orderModel.OrderIsExecute() {
		// 订单为已执行状态
		// 查询此刻的图书 ID
		orderModel.ExecutedOrderDetail(&detailModel)

		// 修改图书销量和库存
		book.ChangeQuantityAndSold(detailModel)
	}

}
