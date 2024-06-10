// Package requests 公告请求验证
package requests

type NoticeValidation struct {
	Title   string `json:"title" form:"title" binding:"required"`
	Content string `json:"content" form:"title" binding:"required"`
}
