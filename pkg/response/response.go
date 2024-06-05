// Package response 响应处理
package response

import (
	"blog/pkg/errcode"
	"github.com/gin-gonic/gin"
)

// Response 响应控制器
type Response struct {
	Ctx   *gin.Context
	Coder errcode.Coder
}

// NewResponse Response 实例
func NewResponse(c *gin.Context, coder errcode.Coder) *Response {
	return &Response{c, coder}
}

// 返回状态码/响应体
func (r *Response) json(c *gin.Context, data interface{}) {

	c.JSON(r.Coder.HTTPStatus(), data)
}

// WithResponse 返回响应内容
func (r *Response) WithResponse(data ...interface{}) {

	response := gin.H{
		"err_code": r.Coder.Code(),
		"message":  r.Coder.Message(),
	}

	if len(r.Coder.Details()) > 0 {
		response["details"] = r.Coder.Details()
	}

	if len(data) != 0 {
		response["data"] = data[0]
	}

	r.json(r.Ctx, response)
}
