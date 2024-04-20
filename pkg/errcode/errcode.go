// Package errcode 常用错误处理方法
package errcode

import (
	"fmt"
	//"sync"
)

type Coder interface {
	Code() CodeInt
	Message() string
	HTTPStatus() int
	Details() []string
}

type ErrorCode struct {
	code       CodeInt
	msg        string
	details    []string
	httpStatus int
}

// Codes 存放全局错误码
var Codes = map[CodeInt]Coder{}

//var codeMux sync.Mutex

// Code 返回自定义错误码
func (e *ErrorCode) Code() CodeInt {
	return e.code
}

// Message 返回自定义错误消息
func (e *ErrorCode) Message() string {
	return e.msg
}

// Details 返回错误细节
func (e *ErrorCode) Details() []string {
	return e.details
}

// HTTPStatus 返回网络状态码
func (e *ErrorCode) HTTPStatus() int {
	return e.httpStatus
}

// WithDetails 添加错误细节
func (e *ErrorCode) withDetails(detail string) {
	e.details = append(e.details, detail)
}

func (c CodeInt) ParseCode() Coder {
	if v, ok := Codes[c]; ok {
		return v
	}
	return Codes[ErrUnknown]
}

func (e *ErrorCode) Error() string {
	return fmt.Sprintf("error_code: %v,error_message: %s", e.Code(), e.Message())
}

// Register 添加错误码
func Register(code CodeInt, httpStatus int, message string, details ...string) {

	coder := &ErrorCode{
		code:       code,
		httpStatus: httpStatus,
		msg:        message,
	}

	for _, detail := range details {
		coder.withDetails(detail)
	}

	if _, ok := Codes[coder.Code()]; ok {
		panic(fmt.Sprintf("err_code: %d 已经存在", coder.Code()))
	}

	//codeMux.Lock()
	//defer codeMux.Unlock()
	Codes[coder.Code()] = coder

}
