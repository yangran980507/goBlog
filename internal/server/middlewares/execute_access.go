// Package middlewares 验证是否执行某操作
package middlewares

import (
	"blog/pkg/app"
	"blog/pkg/errcode"
	"blog/pkg/redis"
	"blog/pkg/response"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

// ExecuteAuth 是否执行某操作
func ExecuteAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取当前用户 uid
		uid := app.CurrentUser(c)

		// 执行完某操作后，往 redis 中存入 hash 数据 uid(key): action(field): answer(value)
		// redis 库中查询 action 对应 answer 是否存在
		if redis.EventRedis.HExists("Vote", uid) {
			count, _ := strconv.Atoi(redis.EventRedis.HGet("Vote", uid))
			then, _ := strconv.Atoi(redis.EventRedis.HGet("Time", uid))
			if expire := time.Duration((time.Now().Unix()-int64(then))/3600) * time.Hour; expire < 24 {
				if count >= 3 {
					response.NewResponse(c, errcode.ErrSuccess).
						WithResponse("每天仅可投票3次")
					c.Abort()
				}
			} else {
				redis.EventRedis.HDel("Vote", uid)
				redis.EventRedis.HDel("Time", uid)
			}
		}
	}
}
