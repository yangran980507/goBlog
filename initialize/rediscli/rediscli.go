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
	redis.Connect(
		fmt.Sprintf("%v:%v", global.RedisSetting.Host, global.RedisSetting.Port),
		global.RedisSetting.User,
		global.RedisSetting.Pw,
		global.RedisSetting.Database,
	)
}
