// Package admin 公告管理
package admin

import (
	"blog/internal/server/models"
	"blog/internal/server/models/notice"
	"blog/internal/server/requests"
	"blog/pkg/errcode"
	"blog/pkg/logger"
	"blog/pkg/response"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

// NoticeGet 显示公告
func (ac *AdminController) NoticeGet(c *gin.Context) {
	notices, rows := notice.Get()
	if rows != 0 {
		response.NewResponse(c, errcode.ErrSuccess.ParseCode()).WithResponse(gin.H{
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
		response.NewResponse(c, errcode.ErrBind.ParseCode()).
			WithResponse("发布失败，请稍后再试")
		return
	}

	noticeModel := notice.Notice{

		Title:    request.Title,
		Content:  request.Content,
		ShowTime: time.Now().Unix(),
	}

	noticeModel.Create()

	if noticeModel.ID > 0 {

		response.NewResponse(c, errcode.ErrSuccess.ParseCode()).WithResponse(gin.H{
			"notice": noticeModel,
		})
	} else {
		response.NewResponse(c, errcode.ErrSuccess.ParseCode()).
			WithResponse("发布出错，请稍后再试")
	}
}

// NoticeDelete 删除公告
func (ac *AdminController) NoticeDelete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.LogIf(err)
		response.NewResponse(c, errcode.ErrUnknown.ParseCode()).WithResponse("删除失败")
		return
	}
	noticeModel := notice.Notice{
		BaseMode: models.BaseMode{ID: uint(id)},
	}

	row := noticeModel.Delete()

	if row == 1 {
		response.NewResponse(c, errcode.ErrSuccess.ParseCode()).
			WithResponse("删除成功")
	} else {
		response.NewResponse(c, errcode.ErrUnknown.ParseCode()).
			WithResponse("删除失败，请稍后重试")
	}
}
