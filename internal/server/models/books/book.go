// Package books 图书模型
package books

import (
	"blog/internal/server/models"
	"blog/pkg/mysql"
)

type Book struct {
	// 图书编号
	models.BaseMode
	// 书号：
	BookNumber string `json:"book_number,omitempty"`
	// 书名
	BookName string `json:"book_name,omitempty"`
	// 图书类型
	BookType string `json:"book_type,omitempty"`
	// 出版社
	Publisher string `json:"publisher,omitempty"`
	// 作者
	Author string `json:"author,omitempty"`
	// 简介
	Introduce string `json:"introduce,omitempty"`
	// 价格
	Price float64 `json:"price,omitempty"`
	// 出版日期
	Pdate int64 `json:"pdate,omitempty"`
	// 封面路径
	PicURL string `json:"picURL,omitempty"`
	// 入库时间
	InTime int64 `json:"inTime,omitempty"`
	// 是否新书 ？
	IsNewBook bool `json:"isNewBook,omitempty"`
	// 是否推荐 ？
	IsCommended bool `json:"isCommended,omitempty"`
	// 库存
	Quantity int `json:"quantity,omitempty"`
	// 已售
	Selled int `json:"selled,omitempty"`
}

// Create 创建图书
func (book *Book) Create() {
	mysql.DB.Create(&book)
}

// Delete 删除图书
func (book *Book) Delete() {
	mysql.DB.Delete(&book, book.BaseMode.ID)
}
