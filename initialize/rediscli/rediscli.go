package rediscli

import (
	"blog/global"
	"blog/pkg/redis"
	"fmt"
)

func InitializeRedis() {
	setupRedis()
}

func setupRedis() {

	// 连接购物车数据库
	redis.ConnectCart(
		fmt.Sprintf("%v:%v", global.RedisSetting.Host, global.RedisSetting.Port),
		global.RedisSetting.User,
		global.RedisSetting.Pw,
		global.RedisSetting.CartDatabase,
	)

	// 连接投票数据库
	redis.ConnectPoll(
		fmt.Sprintf("%v:%v", global.RedisSetting.Host, global.RedisSetting.Port),
		global.RedisSetting.User,
		global.RedisSetting.Pw,
		global.RedisSetting.PollDatabase,
	)
}
