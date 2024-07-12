// Package order order 相关函数
package order

import "blog/pkg/mysql"

// Create 批量创建 order
func Create(orders []*Order) error {
	return mysql.DB.Model(Order{}).Create(&orders).Error
}
