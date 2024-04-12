// Package errcode 定义一些公共的错误码
package errcode

const (

	// ErrSuccess / http.StatusOK / OK
	ErrSuccess = iota + 100000

	// ErrUnknown / http.StatusInternalServerError / Internal Server Error
	ErrUnknown

	// ErrBind / http.StatusBadRequest / Error Occurred While Binding Request
	ErrBind

	// ErrValidation / http.StatusBadRequest / Validation Failed
	ErrValidation

	// ErrTokenInvalid / http.StatusUnauthorized / Authorization Failed
	ErrTokenInvalid

	// ErrNotFound /http.StatusNotFound / Route Did Not Fund
	ErrNotFound
)

func init() {
	Register(ErrSuccess, 200, "OK")
	Register(ErrUnknown, 500, "Internal Server Error")
	Register(ErrBind, 400, "Error Occurred While Binding Request")
	Register(ErrValidation, 400, "Validation Failed")
	Register(ErrTokenInvalid, 401, "Authorization Failed")
	Register(ErrNotFound, 404, "Route Did Not Fund")
}
