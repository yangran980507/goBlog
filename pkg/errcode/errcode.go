// Package errcode 常用错误处理方法
package errcode

import (
	"fmt"
	//"sync"
)

type Coder interface {
	Code() int
	Message() string
	HTTPStatus() int
	Details() []string
}

type ErrorCode struct {
	code       int
	msg        string
	details    []string
	httpStatus int
}

// 存放全局错误码
var codes map[int]Coder

//var codeMux sync.Mutex

// Code 返回自定义错误码
func (e *ErrorCode) Code() int {
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
func (e *ErrorCode) WithDetails(details ...string) {
	//e.details = make([]string, 0)
	for _, d := range details {
		e.details = append(e.details, d)
	}
}

func (e *ErrorCode) ParseCode() Coder {
	if v, ok := codes[e.Code()]; ok {
		return v
	}
	return codes[ErrUnknown]
}

func (e *ErrorCode) Error() string {
	return fmt.Sprintf("error_code: %v,error_message: %s", e.Code(), e.Message())
}

// Register 添加错误码
func Register(code int, httpStatus int, message string, details ...string) {

	coder := &ErrorCode{
		code:       code,
		httpStatus: httpStatus,
		msg:        message,
	}

	if len(details[0]) > 0 {
		coder.WithDetails(details...)
	}

	if _, ok := codes[coder.Code()]; ok {
		panic(fmt.Sprintf("err_code: %d 已经存在", coder.Code()))
	}

	//codeMux.Lock()
	//defer codeMux.Unlock()
	codes[coder.Code()] = coder
}
