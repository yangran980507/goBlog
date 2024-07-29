// Package client 用户订单 handleFunc
package client

import (
	"blog/internal/server/models/book"
	"blog/internal/server/models/cart"
	"blog/internal/server/models/order"
	"blog/internal/server/requests"
	"blog/pkg/app"
	"blog/pkg/errcode"
	"blog/pkg/response"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

// OrdersSubmit 订单提交
func (uc *UserController) OrdersSubmit(c *gin.Context) {
	// 创建验证结构体空值实例
	request := requests.OrderValidation{}
	// 绑定结构数据到验证结构体中
	if ok := requests.BindAndValid(c, &request, requests.OrderValidate); !ok {
		// 绑定 && 验证失败
		return
	}

	// 验证请求的 books 是否存在，库存是否满足
	orderDetails := make([]order.OrdersDetail, 0)
	for _, v := range request.Books {
		orderDetails = append(orderDetails, order.OrdersDetail{
			BookID:   v.ID,
			BuyCount: v.Amount,
		})
	}

	// 返回创建的详细订单切片
	orderDetailsID, cartID := order.OrdersDetailCreate(orderDetails)

	// 当前用户 id
	uidStr := app.CurrentUser(c)
	uid, _ := strconv.Atoi(uidStr)

	orders := make([]order.Order, 0)
	for _, v := range orderDetailsID {
		// 订单信息实例
		orders = append(orders, order.Order{
			Uid:           uint(uid),
			OrderDetailID: v,
			LoginName:     request.LoginName,
			PayWay:        request.PayWay,
			Carry:         request.Carry,
			Address:       request.Address,
			Phone:         request.Phone,
			Date:          time.Now().Unix(),
			Notes:         request.Notes,
			Enforce:       "已提交",
		})
	}

	// 创建订单
	order.OrdersCreate(orders)

	response.NewResponse(c, errcode.ErrSuccess).WithResponse("订单提交成功")

	// 重置购物车数据
	cart.ReplaceCart(uidStr, cartID)

}

// ShowOrders 订单显示
func (uc *UserController) ShowOrders(c *gin.Context) {
	// 获取当前用户 id
	uid := app.CurrentUser(c)

	// 获取订单
	orders := order.OrdersGet(uid)

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
func (uc *UserController) ShowOrdersDetail(c *gin.Context) {
	// 获取当前用户 id
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
