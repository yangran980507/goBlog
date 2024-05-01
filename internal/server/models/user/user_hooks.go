// Package user 存放 user 模型钩子函数
package user

import (
	"blog/pkg/encryption"
	"gorm.io/gorm"
)

// BeforeSave 在创建和更新 user 模型前调用
func (userModel *User) BeforeSave(tx *gorm.DB) (err error) {

	if !encryption.IsEncrypt(userModel.PassWord) {
		// 加密密码
		userModel.PassWord = encryption.ByBcrypt(userModel.PassWord)
	}
	return
}
