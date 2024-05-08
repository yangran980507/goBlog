// Package custom_validate 自定义的验证函数
package custom_validate

import (
	"blog/internal/server/models/user"
)

// ConfirmDoublePSW 判断两次输入是否相等
func ConfirmDoublePSW(first, second string, errs map[string][]string) map[string][]string {
	if first != second {
		// 两次输入不相等，向返回错误信息中添加自定义错误信息
		errs["password_confirm"] = append(errs["password_confirm"], "两次输入的密码不一致")
	}
	return errs
}

// ConfirmUserExist 判断用户是否存在
func ConfirmUserExist(username string, errs map[string][]string) map[string][]string {

	des := []string{"用户名已被注册"}
	if user.IsUserExist(username) {
		errs["exist"] = des
	}
	return errs
}
