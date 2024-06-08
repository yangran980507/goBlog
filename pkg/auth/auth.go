// Package auth 授权登陆逻辑
package auth

import (
	"blog/internal/server/models/user"
	"blog/pkg/errcode"
)

type LoginInfo struct {
	LoginName string
	Password  string
}

type Signiner interface {
	Login() (user.User, error)
}

// Login 使用用户名及密码登陆
func (li *LoginInfo) Login() (user.User, error) {
	// 使用用户名获取数据库中 User 数据模型
	userModel := user.GetUserInfo(li.LoginName)

	// ID == 0 ,用户不存在
	if userModel.ID == 0 {
		return user.User{}, errcode.ErrAccountAbsent
	}

	// 密码验证e
	if !userModel.ComparePSW(li.Password) {
		return user.User{}, errcode.ErrPassWord
	}

	// 返回用户数据
	return userModel, nil
}
