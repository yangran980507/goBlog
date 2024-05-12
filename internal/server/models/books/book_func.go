// Package books 存放 book 相关函数
package books

import "blog/pkg/mysql"

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
