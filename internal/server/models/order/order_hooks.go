// Package order 钩子函数
package order

import "gorm.io/gorm"

func (order *Order) AfterCreate(tx *gorm.DB) (err error) {
	return nil
}
