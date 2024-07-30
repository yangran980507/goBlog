// Package user 存放 user 相关函数
package user

import (
	"blog/internal/server/models/book"
	"blog/internal/server/models/order"
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

// Paginate 查询数据并进行分页
func Paginate(c *gin.Context, count string) (users []User, page paginator.Page) {
	page = paginator.Paginate(
		c,
		mysql.DB.Model(&User{}).
			Not("is_manager = ?", true),
		"admin",
		"users",
		&users,
		count,
		"id",
		"asc")

	return
}

// ChangeAmount 修改消费值
func ChangeAmount(orderModel *order.Order, bookModel book.Book, detailModel order.OrdersDetail) {
	var count = bookModel.Price * float64(detailModel.BuyCount)
	mysql.DB.Model(User{}).Where("login_name = ?", orderModel.LoginName).
		Select("amount").
		Update("amount += ?", count)
}
