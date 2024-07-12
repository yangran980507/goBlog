// Package requests 管理员冻结/解冻用户请求验证
package requests

// Freeze 用户控制验证模型
type Freeze struct {
	IsFrozen bool `json:"is_frozen" form:"is_frozen" binding:"required"`
}
