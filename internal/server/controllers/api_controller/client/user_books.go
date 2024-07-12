// Package client 用户图书 handlerFunc
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

// GetBookByIsNewBook 通过是否新书获取图书
func (uc *UserController) GetBookByIsNewBook(c *gin.Context) {

	countStr := c.Param("count")

	books, page := book.GetByIsNewBook(c, countStr)

	response.NewResponse(c, errcode.ErrSuccess).WithResponse(gin.H{
		"books": books,
		"page":  page,
	})
}

// GetBookByRecommended 通过是否推荐获取图书
func (uc *UserController) GetBookByRecommended(c *gin.Context) {

	countStr := c.Param("count")

	books, page := book.GetByRecommended(c, countStr)

	response.NewResponse(c, errcode.ErrSuccess).WithResponse(gin.H{
		"books": books,
		"page":  page,
	})
}

// GetBookBySold 通过销量获取图书
func (uc *UserController) GetBookBySold(c *gin.Context) {
	countStr := c.Param("count")

	books, page := book.GetBySold(c, countStr)

	response.NewResponse(c, errcode.ErrSuccess).WithResponse(gin.H{
		"books": books,
		"page":  page,
	})
}
