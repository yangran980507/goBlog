// Package requests 定单验证模型
package requests

import (
	"blog/internal/server/models/book"
	"blog/internal/server/models/order"
	"fmt"
	"github.com/thedevsaddam/govalidator"
)

// Book 图书信息
type Book struct {
	ID     uint `json:"id"`
	Amount int  `json:"amount"`
}

// OrderValidation 订单验证模型
type OrderValidation struct {
	// 真实名
	LoginName string `json:"login_name,omitempty" valid:"login_name"`
	// 付款方式
	PayWay string `json:"pay_way,omitempty" valid:"pay_way"`
	// 邮寄方式
	Carry string `json:"carry,omitempty" valid:"carry"`
	// 邮寄地址
	Address string `json:"address,omitempty" valid:"address"`
	// 联系电话
	Phone string `json:"phone" valid:"phone" valid:"phone"`
	// 备注信息
	Notes string `json:"notes"`
	// 提交图书信息
	Books []Book `json:"books,omitempty"`
}

// OrderRefund 订单取消验证
type OrderRefund struct {
	OrderID       uint   `json:"order_id,omitempty"`
	RefundExplain string `json:"refund_explain,omitempty"`
}

// OrderExecute 订单执行验证
type OrderExecute struct {
	OrderID   uint   `json:"order_id,omitempty"`
	LoginName string `json:"login_name,omitempty" form:"login_name" binding:"required"`
	Enforce   string `json:"enforce,omitempty" form:"enforce" binding:"required"`
}

// OrderValidate 订单请求验证函数
func OrderValidate(data interface{}) map[string][]string {

	// 验证规则
	rules := govalidator.MapData{
		"login_name": []string{"required", "between:3,20"},
		"address":    []string{"required"},
		"phone":      []string{"required", "digits:11"},
		"pay_way":    []string{"required"},
		"carry":      []string{"required"},
	}

	// 返回错误信息
	messages := govalidator.MapData{

		"login_name": []string{
			"required: 真实姓名为必填",
			"between: 真实姓名在3到20个字符之间",
		},

		"address": []string{
			"required: 收货地址为必填",
		},

		"phone": []string{
			"required: 联系电话为必填",
			"digits: 手机号码为11位",
		},

		"pay_way": []string{
			"required: 支付方式为必选",
		},

		"carry": []string{
			"required: 邮寄方式为必选",
		},
	}

	// 传入设置的验证规则，错误消息参数，返回错误信息
	errs := validate(data, rules, messages)

	return errs
}

// ConfirmBooks 判断图书是否存在，库存是否满足
func ConfirmBooks(books []Book) (errs []string, meetBooks []order.OrdersDetail) {

	// 遍历图书切片
	for _, value := range books {

		if name, b := book.IsBookSufficient(value.ID, value.Amount); !b {
			// 图书库存不足
			errs = append(errs, fmt.Sprintf("《%s》库存不足", name))
		} else {
			meetBooks = append(meetBooks, order.OrdersDetail{
				BookID:   value.ID,
				BuyCount: value.Amount,
			})
		}
	}
	return
}
