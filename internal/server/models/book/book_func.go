// Package book 存放 book 相关函数
package book

import (
	"blog/pkg/mysql"
	"blog/pkg/paginator"
	"github.com/gin-gonic/gin"
)

// GetByIsNewBook 查询为新书项的图书
func GetByIsNewBook(c *gin.Context, count int) (books []Book, page paginator.Page) {

	page = paginator.Paginate(
		c,
		mysql.DB.Model(Book{}).Where("is_new_book = ?", true),
		"client",
		"books/by-is_new_book",
		&books,
		count,
		"in_time",
		"desc")

	return
}

// GetBySold 查询销售排行图书
func GetBySold(c *gin.Context, count int) (books []Book, page paginator.Page) {

	page = paginator.Paginate(
		c,
		mysql.DB.Model(Book{}),
		"client",
		"books/by-sold",
		&books,
		count,
		"sold",
		"desc")

	return
}

// GetBooksBySlice 通过 id 切片获取图书
func GetBooksBySlice(ids []int64) ([]Book, int64) {
	var books []Book
	row := mysql.DB.Where(ids).Find(&books).RowsAffected
	return books, row
}

// GetBooksAll 获取所有图书
func GetBooksAll(c *gin.Context, count int) (books []Book, page paginator.Page) {

	page = paginator.Paginate(
		c,
		mysql.DB.Model(Book{}),
		"admin",
		"books",
		&books,
		count,
		"in_time",
		"desc")

	return
}

// GetCategories 获取分类
func GetCategories() ([]Category, int64) {
	categories := make([]Category, 0)
	row := mysql.DB.Preload("Books").
		Order("category_id asc").Find(&categories).RowsAffected
	return categories, row
}
