// Package errcode 常用错误处理方法
package errcode

import (
	"fmt"
	//"sync"
)

type ErrorCode struct {
	Code       CodeInt
	Message    string
	Details    string
	HttpStatus int
}

// Codes 存放全局错误码
var Codes = map[CodeInt]ErrorCode{}

// WithDetails 添加具体错误细节
func (ec *ErrorCode) WithDetails(detail string) {
	ec.Details = detail
}

// ParseCode 解析错误码
func (c CodeInt) ParseCode() ErrorCode {
	if v, ok := Codes[c]; ok {
		return v
	}
	return Codes[ErrUnknown]
}

// Register 注册错误码
func Register(code CodeInt, httpStatus int, message string) {

	errorCode := ErrorCode{
		Code:       code,
		HttpStatus: httpStatus,
		Message:    message,
	}

	if _, ok := Codes[code]; ok {
		panic(fmt.Sprintf("err_code: %d 已经存在", errorCode.Code))
	}

	Codes[code] = errorCode

}
