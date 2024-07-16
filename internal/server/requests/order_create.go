// Package requests 定单验证模型
package requests

import "github.com/thedevsaddam/govalidator"

// OrderValidation 订单验证模型
type OrderValidation struct {
	// 用户名
	LoginName string `json:"login_name,omitempty"`
	// 真实名
	TrueName string `json:"true_name,omitempty"`
	// 付款方式
	PayWay string `json:"pay_way,omitempty" valid:"pay_way"`
	// 邮寄地址
	Address string `json:"address,omitempty" valid:"address"`
	// 联系电话
	Phone string `json:"phone" valid:"phone"`
	// 备注信息
	Notes string `json:"notes" valid:"notes"`
}

// OrderValidate 订单请求验证函数
func OrderValidate(data interface{}) map[string][]string {

	// 验证规则
	rules := govalidator.MapData{
		"login_name": []string{"required", "alpha_num", "between:3,15"},
		"true_name":  []string{"required", "between:3,20"},
		"post_code":  []string{"required", "num", "digits:6"},
		"address":    []string{"required"},
		"phone":      []string{"required", "digits:11", "num"},
	}

	// 返回错误信息
	messages := govalidator.MapData{
		"login_name": []string{
			"required: 用户名为必填",
			"alpha_num: 用户名由字母或者数字组成",
			"between: 用户名在3到15个字符之间",
		},

		"true_name": []string{
			"required: 真实姓名为必填",
			"between: 真实姓名在3到20个字符之间",
		},

		"post_code": []string{
			"required: 邮编为必填",
			"num: 邮政编码为纯数字",
			"digits： 邮编仅有六位数字",
		},

		"address": []string{
			"required: 收货地址为必填",
		},

		"phone": []string{
			"required: 联系电话为必填",
			"digits: 手机号码为11位",
			"num: 联系电话由纯数字组成",
		},
	}

	// 传入设置的验证规则，错误消息参数，返回错误信息
	return validate(data, rules, messages)
}
