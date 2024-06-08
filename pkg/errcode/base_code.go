// Package errcode 定义一些公共的错误码
package errcode

type CodeInt int

const (

	// ErrSuccess / http.StatusOK / OK / 200
	ErrSuccess CodeInt = iota + 100000

	// ErrUnknown / http.StatusInternalServerError / Internal Server Error / 500
	ErrUnknown

	// ErrBind / http.StatusBadRequest / Error Occurred While Binding Request / 400
	ErrBind

	// ErrValidation / http.StatusBadRequest / Validation Failed / 400
	ErrValidation

	// ErrTokenInvalid / http.StatusUnauthorized / Authorization Failed / 401
	ErrTokenInvalid

	// ErrNotAdmin / http.StatusUnauthorized / Authorization Failed / 401
	ErrNotAdmin

	// ErrNotFound / http.StatusNotFound / Route Did Not Fund / 404
	ErrNotFound
)

const (
	// ErrEmptyCart / http.StatusOK / OK / 200
	ErrEmptyCart CodeInt = iota + 101000
)

func InitializeErrorCode() {
	Register(ErrSuccess, 200, "OK")
	Register(ErrUnknown, 500, "Internal Server Error")
	Register(ErrBind, 400, "Error Occurred While Binding Request")
	Register(ErrValidation, 400, "Validation Failed")
	Register(ErrTokenInvalid, 401, "Authorization Failed")
	Register(ErrNotAdmin, 401, "Have no Authority to execute")
	Register(ErrNotFound, 404, "Route Did Not Fund")
	Register(ErrEmptyCart, 200, "Shopping Carts is Empty")
}
