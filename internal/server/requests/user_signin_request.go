// Package requests 用户登陆请求验证
package requests

import "github.com/thedevsaddam/govalidator"

type SigninUserValidation struct {
	LoginName string `json:"login_name" valid:"login_name"`
	Password  string `json:"password" valid:"password"`
}

// SigninUserValidate 用户登陆验证
func SigninUserValidate(data interface{}) map[string][]string {
	// 验证规则
	rules := govalidator.MapData{
		"login_name": {"required"},
		"password":   {"required"},
	}

	// 返回错误信息
	messages := govalidator.MapData{
		"login_name": {
			"required: 用户名为必填",
		},
		"password": {
			"required: 密码为必填",
		},
	}

	// 传入设置的验证规则，错误消息参数，返回错误信息
	return validate(data, rules, messages)
}
