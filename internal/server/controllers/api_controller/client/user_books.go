// Package client 用户图书获取 handler
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
	books := make([]book.Book, 5)
	books, page := book.GetByIsNewBook(c, 5)

	response.NewResponse(c, errcode.ErrSuccess).WithResponse(gin.H{
		"books": books,
		"page":  page,
	})
}

// GetBookBySold 通过销量获取图书
func (uc *UserController) GetBookBySold(c *gin.Context) {
	books := make([]book.Book, 5)
	books, page := book.GetBySold(c, 5)

	response.NewResponse(c, errcode.ErrSuccess).WithResponse(gin.H{
		"books": books,
		"page":  page,
	})
}
