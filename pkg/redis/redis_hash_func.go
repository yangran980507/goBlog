// Package redis hash 操作函数
package redis

import (
	"blog/pkg/logger"
	"errors"
	"github.com/go-redis/redis/v8"
)

// HSetNX 存入 key field 对应的 value
func (rds *RedisClient) HSetNX(key string, field string, value any) bool {
	if err := rds.Client.HSetNX(rds.Ctx, key, field, value).Err(); err != nil {
		logger.LogIf(err)
		return false
	}
	return true
}

// HGetAll 读取 key field 对应的 value
func (rds *RedisClient) HGetAll(key string) map[string]string {
	value, err := rds.Client.HGetAll(rds.Ctx, key).Result()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			logger.LogIf(err)
			return nil
		}
	}
	return value
}

// HGetTime 获取 time value
func (rds *RedisClient) HGetTime(key string) (time string) {
	time, err := rds.Client.HGet(rds.Ctx, "Time", key).Result()
	if err != nil {
		logger.LogIf(err)
		return ""
	}
	return time
}

// HDel 删除 key 对应的 field
func (rds *RedisClient) HDel(key string, field string) bool {
	if err := rds.Client.HDel(rds.Ctx, key, field).Err(); err != nil {
		logger.LogIf(err)
		return false
	}
	return true
}

// HIncrBy 增加 key field 对应 value 的值
func (rds *RedisClient) HIncrBy(key string, field string) bool {
	if err := rds.Client.HIncrBy(rds.Ctx, key, field, 1).Err(); err != nil {
		logger.LogIf(err)
		return false
	}
	return true
}

// HExists 判断字段是否存在
func (rds *RedisClient) HExists(key string, field string) bool {

	answer, err := rds.Client.HExists(rds.Ctx, key, field).Result()
	if err != nil {
		logger.LogIf(err)
	}
	return answer
}
