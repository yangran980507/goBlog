// Package cart 购物车模型
package cart

import (
	"blog/internal/server/models/book"
	jsonPkg "blog/pkg/json"
	"blog/pkg/logger"
	"blog/pkg/redis"
)

// Cart 购物车模型
type Cart struct {
	// 购物车中的图书
	Books []book.Book `json:"books"`
	// 购物车更新时间
	UpdateTime int64 `json:"update_time"`
}

// SetCart 购物车入库
func (c *Cart) SetCart(uid string) bool {
	// 序列化键值
	strCart, err := jsonPkg.Marshal(c)
	if err != nil {
		logger.LogIf(err)
		return false
	}
	if !redis.Redis.Set(uid+":cart", strCart, -1) {
		return false
	}
	return true
}

// GetCart 读取购物车
func GetCart(uid string) Cart {
	// 获取键值
	strCart := redis.Redis.Get(uid + ":cart")

	cart := Cart{}
	// 反序列化
	if err := jsonPkg.UnMarshal(strCart, &cart); err != nil {
		logger.LogIf(err)
		return Cart{}
	}
	return cart
}

// DelCart 清空购物车
func DelCart(uid string) bool {

	if !redis.Redis.Del(uid + ":cart") {
		return false
	}

	return true
}
