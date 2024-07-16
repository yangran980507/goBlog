// Package order 定单模型
package order

// Order 定单
type Order struct {
	// 定单编号
	ID uint `json:"-" gorm:"column:id;primaryKey;autoIncrement"`
	// 用户编号
	Uid int `json:"-"`
	// 付款方式
	PayWay string `json:"pay_way,omitempty" gorm:"not null" `
	// 邮寄方式
	Carry string `json:"carry,omitempty" gorm:"not null"`
	// 邮寄地址
	Address string `json:"address,omitempty" gorm:"not null"`
	// 定单生效日期
	Date int64 `json:"date,omitempty" gorm:"not null"`
	// 备注
	Notes string `json:"notes,omitempty"`
	// 是否执行
	Enforce bool `json:"enforce,omitempty" `
	// 退款
	Refund bool `json:"refund,omitempty" `
	// 退款说明
	RefundExplain string `json:"refund_explain"`
}

// OrdersDetail 定单明细
type OrdersDetail struct {
	// 定单明细编号
	ID uint `json:"-" gorm:"column:id;primaryKey;autoIncrement"`
	// 对应定单编号
	OrderID uint `json:"order_id"`
	// 图书编号
	BookID int64 `json:"-"`
	// 购买数量
	BuyCount int `json:"buy_count,omitempty"`
}
