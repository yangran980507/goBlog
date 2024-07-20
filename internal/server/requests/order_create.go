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
	BookID uint `json:"book_id"`
	Amount int  `json:"amount"`
}

// OrderValidation 订单验证模型
type OrderValidation struct {
	// 真实名
	TrueName string `json:"true_name,omitempty" valid:"true_name"`
	// 付款方式
	PayWay string `json:"pay_way,omitempty"`
	// 付款方式
	Carry string `json:"carry,omitempty"`
	// 邮寄地址
	Address string `json:"address,omitempty" valid:"address"`
	// 联系电话
	Phone string `json:"phone" valid:"phone" valid:"phone"`
	// 备注信息
	Notes string `json:"notes"`
	// 提交图书信息
	Books []Book `json:"books,omitempty"`
}

// OrderValidate 订单请求验证函数
func OrderValidate(data interface{}) map[string][]string {

	// 验证规则
	rules := govalidator.MapData{
		"true_name": []string{"required", "between:3,20"},
		"address":   []string{"required"},
		"phone":     []string{"required", "digits:11"},
	}

	// 返回错误信息
	messages := govalidator.MapData{

		"true_name": []string{
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
	}

	// 传入设置的验证规则，错误消息参数，返回错误信息
	errs := validate(data, rules, messages)

	return errs
}

// ConfirmBooks 判断图书是否存在，库存是否满足
func ConfirmBooks(books []Book) (errs []string, meetBooks []order.OrdersDetail) {

	// 遍历图书切片
	for _, value := range books {

		if name, b := book.IsBookSufficient(value.BookID, value.Amount); !b {
			// 图书库存不足
			errs = append(errs, fmt.Sprintf("《%s》库存不足", name))
		} else {
			meetBooks = append(meetBooks, order.OrdersDetail{
				BookID:   value.BookID,
				BuyCount: value.Amount,
			})
		}
	}
	return
}
