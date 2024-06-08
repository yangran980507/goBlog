// Package client 用户图书分类处理函数
package client

import (
	"blog/internal/server/models/book"
	"blog/pkg/errcode"
	"blog/pkg/response"
	"github.com/gin-gonic/gin"
)

// GetBookByCategory 通过分类获取图书
func (uc *UserController) GetBookByCategory(c *gin.Context) {
	var categories []book.Category
	categories, row := book.GetCategories()
	if row != 0 {
		response.NewResponse(c, errcode.ErrSuccess).WithResponse(gin.H{
			"categories": categories,
		})
	}
}
