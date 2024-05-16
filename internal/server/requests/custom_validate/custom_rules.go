// Package custom_validate 存放自定义规则
package custom_validate

import (
	"blog/pkg/mysql"
	"errors"
	"fmt"
	"github.com/thedevsaddam/govalidator"
	"strings"
)

func init() {
	customRuleExist()
}

// 验证上传请求数据是否已存在
func customRuleExist() {
	govalidator.AddCustomRule("exists",
		func(field string, rule string, message string, value interface{}) error {
			rng := strings.Split(strings.TrimLeft(rule, "exists:"), "-")

			requestValue := value.(string)

			var count int64
			mysql.DB.Table(rng[0]).Where(rng[1]+" = ?", requestValue).Count(&count)

			if count != 0 {
				//  如果有自定义错误信息的话
				if message != "" {
					return errors.New(message)
				}
				// 默认的错误信息
				return fmt.Errorf("%v 已被占用", requestValue)
			}

			return nil
		})
}
