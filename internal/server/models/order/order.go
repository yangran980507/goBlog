// Package order 定单模型
package order

import "blog/pkg/mysql"

// Order 订单
type Order struct {
	// 订单编号
	ID uint `json:"id,omitempty" gorm:"column:id;primaryKey;autoIncrement"`
	// 用户编号
	Uid uint `json:"uid,omitempty" gorm:"column:uid;not null,index"`
	// 真实姓名
	LoginName string `json:"login_name,omitempty" gorm:"column:login_name;not null"`
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
	Date int64 `json:"date" gorm:"not null"`
	// 备注
	Notes string `json:"notes,omitempty"`
	// 是否执行
	Enforce string `json:"enforce,omitempty" `
	// 退款
	ExecTime int64 `json:"exec_time,omitempty" `
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

// UserOrderChange 修改订单状态
func (order *Order) UserOrderChange() error {
	return mysql.DB.Model(&order).Select("refund_explain", "enforce", "exec_time").
		Updates(Order{RefundExplain: order.RefundExplain,
			Enforce: order.Enforce, ExecTime: order.ExecTime}).Error
}

// AdminOrderChange 修改订单状态
func (order *Order) AdminOrderChange() error {
	return mysql.DB.Model(&order).Select("enforce").
		Updates(Order{Enforce: order.Enforce}).Error
}

// OrderIsExecute 订单是否已执行
func (order *Order) OrderIsExecute() bool {
	return mysql.DB.Model(Order{}).
		Where("id = ? AND enforce = ?", order.ID, order.Enforce).First(&order).
		RowsAffected > 0
}

// ExecutedOrderDetail 执行过的订单明细订单是否已执行
func (order *Order) ExecutedOrderDetail(detailModel *OrdersDetail) {
	mysql.DB.Model(OrdersDetail{}).
		Where("id = ?", order.OrderDetailID).
		First(&detailModel)
}
