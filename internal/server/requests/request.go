// Package requests 验证接口请求数据相关逻辑
package requests

import (
	"blog/pkg/errcode"
	"blog/pkg/logger"
	"blog/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

// 验证函数类，返回错误信息
type validatorFunc func(object interface{}) map[string][]string

// 初始化验证规则，返回错误信息
func validate(data interface{}, rules govalidator.MapData,
	messages govalidator.MapData) map[string][]string {

	// 初始化验证配置项
	options := govalidator.Options{
		Data:          data,
		Rules:         rules,
		Messages:      messages,
		TagIdentifier: "valid",
	}

	return govalidator.New(options).ValidateStruct()
}

// BindAndValid 统一接口数据绑定，调用验证函数
func BindAndValid(c *gin.Context, object interface{}, handler validatorFunc) bool {

	if err := c.ShouldBind(object); err != nil {
		// 如果绑定数据失败
		logger.LogIf(err)
		return false
	}

	// 调用验证函数验证接口数据
	errs := handler(object)

	if len(errs) > 0 {
		// 发生错误，返回验证错误信息
		response.NewResponse(c, errcode.ErrValidation).WithResponse(errs)
		return false
	}

	return true
}
