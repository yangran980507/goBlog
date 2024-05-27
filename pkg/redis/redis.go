// Package redis redis连接初始化
package redis

import (
	"blog/pkg/logger"
	"context"
	"github.com/go-redis/redis/v8"
	"sync"
)

type RedisClient struct {
	client *redis.Client
}

var Redis *RedisClient
var once sync.Once

// Connect 初始化数据库连接实例
func Connect(add string, user string, pw string, db int) {

	once.Do(func() {
		Redis = NewClient(add, user, pw, db)
	})
}

func NewClient(add string, user string, pw string, db int) *RedisClient {
	// 初始化实例
	rds := &RedisClient{}

	// 初始化配置
	rds.client = redis.NewClient(&redis.Options{
		Addr:     add,
		Username: user,
		Password: pw,
		DB:       db,
	})

	// ping测试
	err := rds.ping()
	logger.LogIf(err)

	// 返回redis
	return rds
}

func (rds *RedisClient) ping() (err error) {
	var ctx = context.Background()

	_, err = rds.client.Ping(ctx).Result()
	return
}
