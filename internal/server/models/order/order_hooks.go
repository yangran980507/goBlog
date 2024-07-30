// Package order 相关钩子函数
package order

import (
	"blog/pkg/errcode"
	"blog/pkg/mysql"
	"gorm.io/gorm"
)

// BeforeUpdate 判断订单是否已执行
func (order *Order) BeforeUpdate(tx *gorm.DB) (err error) {
	var count int64
	mysql.DB.Model(&order).
		Where("id = ? AND enforce = ?", order.ID, "已执行").Count(&count)
	if count > 0 {
		err = errcode.ErrOrderHadExecuted
		return
	}
	return nil
}
