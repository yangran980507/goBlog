// Package admin 管理员投票管理 handlerFunc
package admin

import (
	"blog/internal/server/models/book"
	"blog/internal/server/models/poll"
	"blog/internal/server/requests"
	"blog/pkg/app"
	"blog/pkg/errcode"
	"blog/pkg/logger"
	"blog/pkg/response"
	"github.com/gin-gonic/gin"
)

// SetPoll 设置投票项
func (ac *AdminController) SetPoll(c *gin.Context) {

	// 绑定验证数据
	request := requests.Poll{}
	if err := c.ShouldBind(&request); err != nil {
		// 绑定验证失败
		logger.LogIf(err)
		response.NewResponse(c, errcode.ErrBind, "添加失败").
			WithResponse()
		return
	}

	// 投票模型实例
	pollModel := &poll.Poll{
		OptionName: request.OptionName,
	}

	// 调用设置投票函数
	if !pollModel.SetPoll() {
		// 失败，返回失败信息
		response.NewResponse(c, errcode.ErrServer, "添加失败").
			WithResponse()
		return
	}

	// 成功，返回成功信息
	response.NewResponse(c, errcode.ErrSuccess, "添加成功").
		WithResponse()
}

// GetPoll 获取投票数
func (ac *AdminController) GetPoll(c *gin.Context) {
	// 读取票数
	polls := poll.GetPoll()

	// 成功，返回成功信息
	response.NewResponse(c, errcode.ErrSuccess).
		WithResponse(gin.H{
			"polls": polls,
		})
}

// GetCategory 获取分类
func (ac *AdminController) GetCategory(c *gin.Context) {

	categories, err := book.GetCategory()
	if err != nil {
		logger.LogIf(err)
		response.NewResponse(c, errcode.ErrServer).WithResponse("服务器错误")
		return
	}
	response.NewResponse(c, errcode.ErrSuccess).WithResponse(gin.H{
		"categories": categories,
	})
}

// DeletePoll 删除投票项
func (ac *AdminController) DeletePoll(c *gin.Context) {
	// 获取要删除的投票项
	optionName := app.GetStrFromAPI(c, "option_name")

	// 投票模型实例
	pollModel := &poll.Poll{
		OptionName: optionName,
	}

	if !pollModel.DeletePoll() {
		// 失败，返回失败信息
		response.NewResponse(c, errcode.ErrServer, "删除失败").
			WithResponse()
		return
	}
	// 成功，返回成功信息
	response.NewResponse(c, errcode.ErrSuccess, "删除成功").
		WithResponse()

}
