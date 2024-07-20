// Package order 定单模型
package order

import "blog/pkg/mysql"

// Order 订单
type Order struct {
	// 订单编号
	ID uint `json:"-" gorm:"column:id;primaryKey;autoIncrement"`
	// 用户编号
	Uid uint `json:"uid"`
	// 对应订单明细编号
	OrderDetailID uint `json:"order_detail_id"`
	// 付款方式
	PayWay string `json:"pay_way,omitempty" gorm:"not null" `
	// 邮寄方式
	Carry string `json:"carry,omitempty" gorm:"not null"`
	// 邮寄地址
	Address string `json:"address,omitempty" gorm:"not null"`
	// 联系电话
	Phone string `json:"phone,omitempty" gorm:"not null"`
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

// OrdersDetail 订单明细
type OrdersDetail struct {
	// 订单明细编号
	ID uint `json:"-" gorm:"column:id;primaryKey;autoIncrement"`
	// 图书编号
	BookID uint `json:"-"`
	// 购买数量
	BuyCount int `json:"buy_count,omitempty"`
}

// OrdersDetailCreate 订单明细创建
func OrdersDetailCreate(bookMap []OrdersDetail) (orderDetailsID []uint) {
	mysql.DB.Create(&bookMap)

	for _, v := range bookMap {
		orderDetailsID = append(orderDetailsID, v.ID)
	}

	return
}

// OrdersCreate 订单明细创建
func OrdersCreate(orders []Order) (rows int64) {

	return mysql.DB.Create(&orders).RowsAffected

}
