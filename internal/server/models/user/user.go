// Package user 用户模型
package user

import (
	"blog/internal/server/models"
	"blog/pkg/mysql"
)

type User struct {
	// 用户编号
	models.BaseMode
	// 登陆名
	LoginName string `json:"login_name"`
	// 真实名
	TrueName string `json:"true_name,omitempty"`
	// 密码
	PassWord string `json:"-"`
	// 地址
	Address string `json:"address,omitempty"`
	// 邮编
	PostCode string `json:"post_code,omitempty"`
	// 电话号码
	Phone string `json:"-"`
	// 账号是否可用
	Freeze bool `json:"freeze"`
	// 用户身份 1：用户；2：管理
	Status int `json:"status,omitempty"`
	// 折扣等
	Grade int `json:"grade"`
	// 消费金额
	Amount float64 `json:"amount"`
}

// Create 增加数据
func (userModel *User) Create() {
	mysql.DB.Create(&userModel)
}