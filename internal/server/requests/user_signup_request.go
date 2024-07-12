// Package requests 用户注册请求验证
package requests

import (
	"blog/internal/server/requests/custom_validate"
	"github.com/thedevsaddam/govalidator"
)

// SignupUserValidation 用户注册验证模型
type SignupUserValidation struct {
	// 登陆名
	LoginName string `json:"login_name" valid:"login_name"`
	// 真实名
	TrueName string `json:"true_name" valid:"true_name"`
	// 密码
	PassWord string `json:"pass_word" valid:"pass_word"`
	// 密码验证
	PassWordConfirm string `json:"password_confirm" valid:"password_confirm"`
	// 电话号码
	Phone string `json:"phone" valid:"phone"`
	// 地址
	Address string `json:"address" valid:"address"`
}

// SignupUserValidate 用户注册验证
func SignupUserValidate(data interface{}) map[string][]string {

	rules := govalidator.MapData{
		"login_name":       []string{"required", "alpha_num", "between:3,15"},
		"true_name":        []string{"required", "between:3,20"},
		"pass_word":        []string{"required", "alpha_num", "min:6"},
		"password_confirm": []string{"required"},
		"phone":            []string{"required", "digits:11"},
	}

	messages := govalidator.MapData{

		"login_name": []string{
			"required: 用户名为必填",
			"alpha_num: 用户名由字母或者数字组成",
			"between: 用户名在3到15个字符之间",
		},

		"true_name": []string{
			"required: 用户姓名为必填",
			"between: 用户姓名在3到20个字符之间",
		},

		"pass_word": []string{
			"required: 密码为必填",
			"alpha_num: 密码字母或数字组成",
			"min: 密码最少为6位",
		},

		"password_confirm": []string{
			"required: 再次输入密码",
		},

		"phone": []string{
			"required: 电话号码为必填",
			"digits: 手机号码为11位",
		},
	}

	// 传入设置的验证规则，错误消息参数，返回错误信息
	errs := validate(data, rules, messages)

	_data := data.(*SignupUserValidation)
	// 自定义验证两次密码是否相等
	custom_validate.ConfirmDoublePSW(_data.PassWord, _data.PassWordConfirm, errs)
	// 自定义验证用户名是否被注册
	custom_validate.ConfirmUserExist(_data.LoginName, errs)

	// 返回错误信息
	return errs
}
