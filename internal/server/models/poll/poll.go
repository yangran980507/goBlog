// Package poll 投票模型
package poll

import (
	"blog/pkg/redis"
	"strconv"
)

type Poll struct {
	// 投票项
	OptionName string `json:"option_name"`
	// 得票数
	Count int `json:"count"`
}

// SetPoll 设置投票项
func (p *Poll) SetPoll() bool {
	if !redis.PollRedis.HSetNX("Poll", p.OptionName, 0) {
		return false
	}
	return true
}

// GetPoll 读取投票数
func GetPoll() []Poll {
	var mPoll map[string]string
	var sPoll []Poll
	mPoll = redis.PollRedis.HGetAll("Poll")
	for f, v := range mPoll {
		vInt, _ := strconv.Atoi(v)
		sPoll = append(sPoll, Poll{
			OptionName: f,
			Count:      vInt,
		})
	}
	return sPoll
}

// GetPollOpts 获取投票项
func GetPollOpts() []string {
	mPoll := redis.PollRedis.GetKeys("Poll")
	return mPoll
}

// IncrPoll 投票
func (p *Poll) IncrPoll() bool {
	if !redis.PollRedis.HIncrBy("Poll", p.OptionName) {
		return false
	}
	return true
}

// DeletePoll 删除投票项
func (p *Poll) DeletePoll() bool {
	if !redis.PollRedis.HDel("Poll", p.OptionName) {
		return false
	}
	return true
}
