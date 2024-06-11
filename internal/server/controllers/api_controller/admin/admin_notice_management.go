// Package admin 管理员公告管理 handlerFunc
package admin

import (
	"blog/internal/server/models"
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
	}
}

// NoticeCreate 发布公告
func (ac *AdminController) NoticeCreate(c *gin.Context) {
	request := requests.NoticeValidation{}
	if err := c.ShouldBind(&request); err != nil {
		// 绑定验证失败
		logger.LogIf(err)
		response.NewResponse(c, errcode.ErrBind, "发布失败，请稍后再试").
			WithResponse()
		return
	}

	noticeModel := notice.Notice{

		Title:    request.Title,
		Content:  request.Content,
		ShowTime: time.Now().Unix(),
	}

	noticeModel.Create()

	if noticeModel.ID > 0 {

		response.NewResponse(c, errcode.ErrSuccess).WithResponse(gin.H{
			"notice": noticeModel,
		})
	} else {
		response.NewResponse(c, errcode.ErrSuccess).
			WithResponse("发布出错，请稍后再试")
	}
}

// NoticeDelete 删除公告
func (ac *AdminController) NoticeDelete(c *gin.Context) {
	// 解析接口数据
	id := app.GetIDFromAPI(c, "id")

	noticeModel := notice.Notice{
		BaseMode: models.BaseMode{ID: uint(id)},
	}

	row := noticeModel.Delete()

	if row == 1 {
		response.NewResponse(c, errcode.ErrSuccess, "删除成功").
			WithResponse()
	} else {
		response.NewResponse(c, errcode.ErrServer, "删除失败，请稍后重试").
			WithResponse()
	}
}
