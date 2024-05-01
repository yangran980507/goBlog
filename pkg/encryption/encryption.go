// Package encryption 加密相关的操作
package encryption

import (
	"blog/initialize/logger"
	"golang.org/x/crypto/bcrypt"
)

// ByBcrypt 使用 bcrypt 对传入字符串进行加密
func ByBcrypt(psw string) string {
	// cost 值越大，耗费时间越长
	bytes, err := bcrypt.GenerateFromPassword([]byte(psw), 14)
	logger.LogIf(err)
	// 返回加密字符串
	return string(bytes)
}

// IsEncrypt 判断是否被加密
func IsEncrypt(psw string) bool {
	return len(psw) == 60
}

// BcryptCheck 将加密后的数据与原数据进行比较
func BcryptCheck(psw string, encrypted string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encrypted), []byte(psw))
	return err == nil
}
