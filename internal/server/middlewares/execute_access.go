// Package middlewares 验证是否执行某操作
package middlewares

import (
	"blog/pkg/app"
	"blog/pkg/redis"
	"github.com/gin-gonic/gin"
)

// ExecuteAuth 是否执行某操作
func ExecuteAuth(action string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取当前用户 uid
		uid := app.CurrentUser(c)

		// 执行完某操作后，往 redis 中存入 hash 数据 uid(key): action(field): answer(value)
		// redis 库中查询 action 对应 answer 是否存在
		if redis.QuestionRedis.HExists("users:"+uid, action) {
			c.Set("isExecute", true)
		}
		c.Set("action", action)
	}
}
