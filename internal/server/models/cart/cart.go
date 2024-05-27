// Package cart 购物车模型
package cart

import "blog/internal/server/models/book"

// Cart 购物车模型
type Cart struct {
	// 购物车中的图书
	Books []book.Book `json:"books"`
	// 购物车更新时间
	UpdateTime int64
}
