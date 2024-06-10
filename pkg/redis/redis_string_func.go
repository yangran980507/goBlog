// Package redis string 操作函数
package redis

import (
	"blog/pkg/logger"
	"errors"
	"github.com/go-redis/redis/v8"
	"time"
)

// Set 存入 key 对应的 value ，设置过期时间
func (rds *RedisClient) Set(key string, value interface{}, expire time.Duration) bool {
	if err := rds.Client.Set(rds.Ctx, key, value, expire).Err(); err != nil {
		logger.LogIf(err)
		return false
	}
	return true
}

// Get 获取 key 对应 value
func (rds *RedisClient) Get(key string) string {
	value, err := rds.Client.Get(rds.Ctx, key).Result()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			logger.LogIf(err)
		}
		return ""
	}
	return value
}

// Del 删除 key 对应的 value
func (rds *RedisClient) Del(key ...string) bool {
	if err := rds.Client.Del(rds.Ctx, key...).Err(); err != nil {
		logger.LogIf(err)
		return false
	}
	return true
}
