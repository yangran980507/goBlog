// Package order 定单模型
package order

import "blog/internal/server/models"

// Order 定单
type Order struct {
	// 定单编号
	models.BaseMode
	// 用户编号
	Uid models.BaseMode
	// 付款方式
	Pay string
	// 邮寄方式
	Carry string
	// 邮寄地址
	Address string
	// 定单生效日期
	Date int64
	// 备注
	Notes string
	// 是否执行
	Enforce bool
}

// OrdersDetail 定单明细
type OrdersDetail struct {
}
