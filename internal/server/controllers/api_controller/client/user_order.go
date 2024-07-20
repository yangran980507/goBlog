// Package client 用户订单 handleFunc
package client

import (
	"blog/internal/server/models/order"
	"blog/internal/server/requests"
	"blog/pkg/app"
	"blog/pkg/errcode"
	"blog/pkg/response"
	"fmt"
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
	errs, orderDetails := requests.ConfirmBooks(request.Books)

	if len(orderDetails) == 0 {
		response.NewResponse(c, errcode.ErrBooksQuantityDeficit).
			WithResponse(gin.H{
				"errs": errs,
			})
		return
	}

	orderDetailsID := order.OrdersDetailCreate(orderDetails)

	// 当前用户 id
	uid, _ := strconv.Atoi(app.CurrentUser(c))

	orders := make([]order.Order, 0)
	for _, v := range orderDetailsID {
		// 订单信息实例
		orders = append(orders, order.Order{
			Uid:           uint(uid),
			OrderDetailID: v,
			PayWay:        request.PayWay,
			Carry:         request.Carry,
			Address:       request.Address,
			Phone:         request.Phone,
			Date:          time.Now().Unix(),
			Notes:         request.Notes,
		})
	}

	rows := order.OrdersCreate(orders)

	if int(rows) == len(orders) {
		fmt.Println("success")
	}

	if len(errs) != 0 {
		response.NewResponse(c, errcode.ErrBooksQuantityDeficit).WithResponse(
			gin.H{
				"errs": errs,
			})
	} else {
		response.NewResponse(c, errcode.ErrSuccess).WithResponse("订单提交成功")
	}
}
