// Package response 响应处理
package response

import (
	"blog/pkg/errcode"
	"github.com/gin-gonic/gin"
)

// Response 响应控制器
type Response struct {
	Ctx     *gin.Context
	ErrCode errcode.ErrorCode
	Data    any
}

// NewResponse Response 实例
func NewResponse(c *gin.Context, code errcode.CodeInt, details ...string) *Response {

	// 解析错误码
	errCode := code.ParseCode()

	// 如果 details 不为空
	if len(details) != 0 {
		errCode.WithDetails(details[0])
	}

	// 返回 Response 实例
	return &Response{
		Ctx:     c,
		ErrCode: errCode}
}

// Json 返回 JSON
func (r *Response) json() {
	r.Ctx.JSON(r.ErrCode.HttpStatus, r.Data)
}

// WithResponse 返回响应
func (r *Response) WithResponse(data ...interface{}) {

	// 默认返回错误码 err_code 错误信息 message
	response := gin.H{
		"err_code": r.ErrCode.Code,
		"message":  r.ErrCode.Message,
	}

	// 如果详细信息存在
	if r.ErrCode.Details != "" {
		response["details"] = r.ErrCode.Details
	}

	// 数据库数据不为空
	if len(data) != 0 {
		response["data"] = data[0]
	}

	r.Data = response
	r.json()
}
