// Package requests 投票验证
package requests

type Poll struct {
	OptionName string `json:"option_name" form:"option_name" binding:"required"`
}
