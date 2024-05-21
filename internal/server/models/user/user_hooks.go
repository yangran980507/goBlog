// Package user 存放 user 模型钩子函数
package user

import (
	"blog/pkg/encryption"
	"gorm.io/gorm"
)

// BeforeCreate 在创建和更新 user 模型前调用
func (userModel *User) BeforeCreate(tx *gorm.DB) (err error) {

	if !encryption.IsEncrypt(userModel.PassWord) {
		// 加密密码
		userModel.PassWord = encryption.ByBcrypt(userModel.PassWord)
	}
	return
}
