// Package admin 管理员公告管理 handlerFunc
package admin

import (
	"blog/internal/server/models/notice"
	"blog/internal/server/requests"
	"blog/pkg/app"
	"blog/pkg/errcode"
	"blog/pkg/logger"
	"blog/pkg/response"
	"github.com/gin-gonic/gin"
	"time"
)

// NoticeGet 显示公告
func (ac *AdminController) NoticeGet(c *gin.Context) {
	notices, rows := notice.Get()

	if rows != 0 {
		response.NewResponse(c, errcode.ErrSuccess).WithResponse(gin.H{
			"notices": notices,
		})
	} else {
		response.NewResponse(c, errcode.ErrEmptyValue).WithResponse("无数据")
	}
}

// NoticeCreate 发布公告
func (ac *AdminController) NoticeCreate(c *gin.Context) {
	request := requests.NoticeValidation{}
	if err := c.ShouldBind(&request); err != nil {
		// 绑定验证失败
		logger.LogIf(err)
		response.NewResponse(c, errcode.ErrBind).
			WithResponse("发布失败，请稍后再试")
		return
	}

	noticeModel := notice.Notice{
		
		Content:  request.Content,
		ShowTime: time.Now().Unix(),
	}

	noticeModel.Create()

	if noticeModel.ID > 0 {
		// 发布成功
		response.NewResponse(c, errcode.ErrSuccess).WithResponse("发布成功")
	} else {
		response.NewResponse(c, errcode.ErrServer).
			WithResponse("发布失败，请稍后再试")
	}
}

// NoticeDelete 删除公告
func (ac *AdminController) NoticeDelete(c *gin.Context) {
	// 解析接口数据
	id := app.GetIDFromAPI(c, "id")

	noticeModel := notice.Notice{
		ID: uint(id),
	}

	row := noticeModel.Delete()

	if row == 1 {
		response.NewResponse(c, errcode.ErrSuccess).
			WithResponse("删除成功")
	} else {
		response.NewResponse(c, errcode.ErrServer).
			WithResponse("删除失败，请稍后重试")
	}
}
