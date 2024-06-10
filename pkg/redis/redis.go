// Package redis redis连接初始化
package redis

import (
	"blog/pkg/logger"
	"context"
	"github.com/go-redis/redis/v8"
	"sync"
)

type RedisClient struct {
	Client *redis.Client
	Ctx    context.Context
}

var CartRedis, PollRedis *RedisClient
var onceCart, oncePoll sync.Once

// ConnectCart 初始化数据库连接实例
func ConnectCart(add string, user string, pw string, db int) {

	onceCart.Do(func() {
		CartRedis = NewClient(add, user, pw, db)
	})
}

// ConnectPoll 初始化数据库连接实例
func ConnectPoll(add string, user string, pw string, db int) {

	oncePoll.Do(func() {
		PollRedis = NewClient(add, user, pw, db)
	})
}

func NewClient(add string, user string, pw string, db int) *RedisClient {
	// 初始化实例
	rds := &RedisClient{}
	rds.Ctx = context.Background()

	// 初始化配置
	rds.Client = redis.NewClient(&redis.Options{
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
	_, err = rds.Client.Ping(rds.Ctx).Result()
	return
}
