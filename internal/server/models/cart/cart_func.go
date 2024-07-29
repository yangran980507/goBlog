package cart

import "blog/pkg/helps"

// ReplaceCart 重置购物车数据
func ReplaceCart(uid string, condition []int64) (bool, int) {
	// 获取购物车信息
	cartModel := GetCart(uid)

	// 换取新的 购物车数据
	cartModel.BookID = helps.GenerateNewSliceByDeleteOldSlice(cartModel.BookID, condition)

	return cartModel.SetCart(uid), len(cartModel.BookID)
}
