// Package errcode 定义一些公共的错误码
package errcode

type CodeInt int

// 实现 Error 接口
func (c CodeInt) Error() string {
	return c.ParseCode().Message
}

// 基础错误
const (

	// ErrSuccess / http.StatusOK / OK / 200
	ErrSuccess CodeInt = iota + 100001

	// ErrServer / http.StatusInternalServerError / Internal Server Error / 500
	ErrServer

	// ErrBind / http.StatusBadRequest / Error Occurred While Binding Request / 400
	ErrBind

	// ErrNotFound / http.StatusNotFound / Route Did Not Fund / 404
	ErrNotFound

	// ErrUnknown / http.StatusInternalServerError / Internal Server Error / 500
	ErrUnknown
)

// 用户授权模块错误
const (
	// ErrValidation / http.StatusBadRequest / Validation Failed / 400
	ErrValidation CodeInt = iota + 100101

	// ErrTokenInvalid / http.StatusUnauthorized / Authorization Failed / 401
	ErrTokenInvalid

	// ErrTokenTimeOut / http.StatusUnauthorized / Authorization Failed / 401
	ErrTokenTimeOut

	// ErrNotAdmin (非管理员账户) / http.StatusUnauthorized / Authorization Failed / 401
	ErrNotAdmin

	// ErrLoginNameUsed (登录名已被使用) / http.StatusUnauthorized / Authorization Failed / 401
	ErrLoginNameUsed

	// ErrAccountAbsent (账户不存在)/ http.StatusNotFound / Authorization Failed / 404
	ErrAccountAbsent

	// ErrPassWord (密码错误) / http.StatusNotFound / Authorization Failed / 401
	ErrPassWord

	// ErrFrozen (用户冻结) / http.StatusUnauthorized / Authorization Failed / 401
	ErrFrozen
)

// 用户业务模块错误
const (
	// ErrEmptyValue (空数据库) / http.StatusOK / OK / 200
	ErrEmptyValue CodeInt = iota + 100201

	// ErrOverMaxCount (超过购物车最大存量) / http.StatusOK / OK / 200
	ErrOverMaxCount

	// ErrBookHadExisted (图书已在购物车中) / http.StatusOK / OK / 200
	ErrBookHadExisted

	// ErrBookHadRemoved (图书已下架) / http.StatusOK / OK / 200
	ErrBookHadRemoved

	// ErrBooksQuantityDeficit (图书库存不足) / http.StatusOK / OK / 200
	ErrBooksQuantityDeficit

	// ErrPollHadExisted (投票项存在) / http.StatusOK / OK / 200
	ErrPollHadExisted

	// ErrOrderHadExecuted (订单已执行) / http.StatusOK / OK / 200
	ErrOrderHadExecuted
)

func InitializeErrorCode() {
	Register(ErrSuccess, 200, "OK")
	Register(ErrServer, 500, "Internal Server Error")
	Register(ErrUnknown, 500, "Internal Server Error")
	Register(ErrBind, 400, "Error Occurred While Binding Request")
	Register(ErrValidation, 400, "Validation Failed")
	Register(ErrTokenTimeOut, 401, "Token Had Over MaxFreshTime")
	Register(ErrTokenInvalid, 401, "Authorization Failed")
	Register(ErrNotAdmin, 401, "Have no Authority to execute")
	Register(ErrNotFound, 404, "Route Did Not Found")
	Register(ErrLoginNameUsed, 401, "LoginName is Used")
	Register(ErrAccountAbsent, 404, "Account is Not Found")
	Register(ErrPassWord, 401, "PassWord Error")
	Register(ErrEmptyValue, 200, "DataBase is Empty")
	Register(ErrOverMaxCount, 200, "Shopping Carts is Fulled")
	Register(ErrBookHadExisted, 200, "Book Had Existed")
	Register(ErrBookHadRemoved, 200, "Book Had Removed")
	Register(ErrBooksQuantityDeficit, 200, "Books Quantity Deficit")
	Register(ErrPollHadExisted, 200, "Poll Had Existed")
	Register(ErrOrderHadExecuted, 200, "Order Had Executed")
	Register(ErrFrozen, 401, "User Had Frozen")
}
