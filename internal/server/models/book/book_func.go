// Package book 存放 book 相关函数
package book

import (
	"blog/pkg/mysql"
	"blog/pkg/paginator"
	"github.com/gin-gonic/gin"
)

// GetByIsNewBook 查询为新书项的图书
func GetByIsNewBook() ([]*Book, error) {

	var books []*Book
	if err := mysql.DB.Model(Book{}).
		Where("isNewBook = ?", true).Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

// GetByIsCommended 查询为推荐项的图书
func GetByIsCommended() ([]*Book, error) {

	var books []*Book
	if err := mysql.DB.Model(Book{}).
		Where("isCommended = ?", true).
		Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

// GetBooksAll 获取所有图书
func GetBooksAll(c *gin.Context, count int) (books []Book, page paginator.Page) {

	page = paginator.Paginate(
		c,
		mysql.DB.Model(Book{}),
		"admin",
		"book",
		&books,
		count,
		"in_time",
		"desc")

	return
}

// GetCategories 获取分类
func GetCategories() ([]Category, int64) {
	categories := make([]Category, 10)
	row := mysql.DB.Preload("Books").
		Order("id asc").Find(&categories).RowsAffected
	return categories, row
}
