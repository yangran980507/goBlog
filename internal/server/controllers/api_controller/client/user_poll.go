// Package client 用户投票 handlerFunc
package client

import (
	"blog/internal/server/models/poll"
	"blog/internal/server/requests"
	"blog/pkg/app"
	"blog/pkg/errcode"
	"blog/pkg/logger"
	"blog/pkg/redis"
	"blog/pkg/response"
	"cmp"
	"github.com/gin-gonic/gin"
	"slices"
	"time"
)

// IncrByPoll 投票
func (uc *UserController) IncrByPoll(c *gin.Context) {

	// 绑定验证数据
	request := requests.Poll{}
	if err := c.ShouldBind(&request); err != nil {
		// 绑定验证失败
		logger.LogIf(err)
		response.NewResponse(c, errcode.ErrBind).
			WithResponse("投票失败")
		return
	}

	// 投票模型实例
	pollModel := &poll.Poll{
		OptionName: request.OptionName,
	}

	if !pollModel.IncrPoll() {
		response.NewResponse(c, errcode.ErrServer).
			WithResponse("投票失败")
		return
	}

	// 获取当前 uid,action
	uid := app.CurrentUser(c)

	// 存入操作记录
	if redis.EventRedis.HExists("Vote", uid) {
		redis.EventRedis.HIncrBy("Vote", uid)
	} else {
		redis.EventRedis.HSetNX("Vote", uid, 1)
		redis.EventRedis.HSetNX("Time", uid, time.Now().Unix())
	}

	// 返回投票成功消息
	response.NewResponse(c, errcode.ErrSuccess).
		WithResponse("投票成功")
}

// GetPoll 获取投票数
func (uc *UserController) GetPoll(c *gin.Context) {
	// 读取票数
	polls := poll.GetPoll()

	if polls == nil {
		response.NewResponse(c, errcode.ErrEmptyValue).
			WithResponse("empty data")
		return
	}
	slices.SortStableFunc(polls, func(a, b poll.Poll) int {
		return cmp.Compare(a.Time, b.Time)
	})
	// 成功，返回成功信息
	response.NewResponse(c, errcode.ErrSuccess).
		WithResponse(gin.H{
			"polls": polls,
		})
}

// GetPollOption 获取投票项
/*
func (uc *UserController) GetPollOption(c *gin.Context) {
	// 读取投票项
	polls := poll.GetPollOpts()

	// 返回投票项
	response.NewResponse(c, errcode.ErrSuccess).
		WithResponse(gin.H{
			"pollKeys": polls,
		})
}
*/
