// Package user 存放 user 相关函数
package user

import "blog/pkg/mysql"

// IsUserExist 判断 user 是否存在于数据库中
func IsUserExist(user string) bool {
	var count int64
	mysql.DB.Model(&User{}).Where("login_name = ?", user).
		Count(&count)
	return count > 0
}
