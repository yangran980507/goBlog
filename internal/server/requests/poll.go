// Package requests 投票验证
package requests

// Poll 投票验证模型
type Poll struct {
	OptionName string `json:"option_name" form:"option_name" binding:"required"`
}
