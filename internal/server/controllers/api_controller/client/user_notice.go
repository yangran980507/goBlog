// Package client 用户公告 handlerFunc
package client

import (
	"blog/internal/server/models/notice"
	"blog/pkg/errcode"
	"blog/pkg/response"
	"github.com/gin-gonic/gin"
)

// NoticeGet 获取公告
func (uc *UserController) NoticeGet(c *gin.Context) {
	notices, rows := notice.Get()
	if rows != 0 {
		response.NewResponse(c, errcode.ErrSuccess).WithResponse(gin.H{
			"notices": notices,
		})
	}
}
