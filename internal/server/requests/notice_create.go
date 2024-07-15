// Package requests 公告请求验证模型
package requests

// NoticeValidation 公告验证模型
type NoticeValidation struct {
	Content string `json:"content" form:"content" binding:"required"`
}
