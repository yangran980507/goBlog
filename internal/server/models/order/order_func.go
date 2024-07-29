// Package order order 相关函数
package order

import "blog/pkg/mysql"

// OrdersDetailCreate 订单明细创建
func OrdersDetailCreate(bookMap []OrdersDetail) (orderDetailsID []uint, cart []int64) {
	mysql.DB.Create(&bookMap)

	for _, v := range bookMap {
		orderDetailsID = append(orderDetailsID, v.ID)
		cart = append(cart, int64(v.BookID))
	}

	return
}

// OrdersCreate 订单明细创建
func OrdersCreate(orders []Order) (rows int64) {

	return mysql.DB.Create(&orders).RowsAffected

}

// OrdersGet 获取订单
func OrdersGet(uid string) (orders []Order) {
	mysql.DB.Model(Order{}).
		Select([]string{"order_detail_id", "pay_way", "carry", "enforce",
			"date", "login_name", "refund", "refund_explain"}).
		Where("uid = ?", uid).Find(&orders)

	return
}

// OrdersDetailGet 获取订单详细
func OrdersDetailGet(id string) (orderDetail OrdersDetail) {
	mysql.DB.Model(OrdersDetail{}).
		Where("id = ?", id).First(&orderDetail)

	return
}
