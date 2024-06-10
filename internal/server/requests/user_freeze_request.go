// Package requests 管理员冻结/解冻用户请求验证
package requests

type Person struct {
	LoginName  string `json:"login_name" form:"login_name" binding:"required"`
	IsFreezing bool   `json:"is_freezing" form:"is_freezing" binding:"required"`
}
