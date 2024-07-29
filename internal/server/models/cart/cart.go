// Package cart 购物车模型
package cart

import (
	jsonPkg "blog/pkg/json"
	"blog/pkg/logger"
	"blog/pkg/redis"
	"time"
)

// Cart 购物车模型
type Cart struct {
	// 购物车中的图书
	BookID []int64 `json:"book_id"`
}

// SetCart 购物车入库
func (cart *Cart) SetCart(uid string) bool {
	// 序列化键值
	strCart, err := jsonPkg.Marshal(cart)
	if err != nil {
		logger.LogIf(err)
		return false
	}
	if !redis.CartRedis.Set(uid+":cart", strCart, time.Duration(604800)*time.Second) {
		return false
	}
	return true
}

// GetCart 读取购物车
func GetCart(uid string) Cart {
	// 获取键值
	strCart := redis.CartRedis.Get(uid + ":cart")

	cart := Cart{}
	// 反序列化
	if err := jsonPkg.UnMarshal(strCart, &cart); err != nil {
		return Cart{}
	}
	return cart
}

// DelCart 清空购物车
func DelCart(uid string) bool {

	if !redis.CartRedis.Del(uid + ":cart") {
		return false
	}

	return true
}
