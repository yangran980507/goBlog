// Package user 存放 user 相关函数
package user

import (
	"blog/pkg/mysql"
	"blog/pkg/paginator"
	"github.com/gin-gonic/gin"
)

// IsUserExist 判断 user 是否存在于数据库中
func IsUserExist(user string) bool {
	var count int64
	mysql.DB.Model(&User{}).Where("login_name = ?", user).
		Count(&count)
	return count > 0
}

// GetUserInfo 获取用户信息
func GetUserInfo(loginName string) (userModel User) {
	mysql.DB.Where("login_name = ?", loginName).First(&userModel)
	return
}

// GetUsers 获取所有用户
func GetUsers() ([]User, error) {
	users := make([]User, 6)
	if err := mysql.DB.Model(User{}).Where("is_manager = ?", false).
		Order("id asc").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// Paginate 查询数据并进行分页
func Paginate(c *gin.Context, count int) (users []User, page paginator.Page) {
	page = paginator.Paginate(
		c,
		mysql.DB.Model(&User{}).Not("is_manager = ?", true),
		"admin",
		"users",
		&users,
		count,
		"id",
		"asc")

	return
}
