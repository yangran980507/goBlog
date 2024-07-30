// Package book 存放 book 相关函数
package book

import (
	"blog/internal/server/models/order"
	"blog/pkg/mysql"
	"blog/pkg/paginator"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// IsBookSufficient 验证图书库存是否满足
func IsBookSufficient(id uint, amount int) (string, bool) {
	var book Book
	mysql.DB.Model(Book{}).Where("id = ?", id).First(&book)

	if book.Quantity < amount {
		return book.BookName, false
	}
	return book.BookName, true
}

// GetByID ID获取
func GetByID(id uint) (bookModel Book, row int64) {
	result := mysql.DB.Model(Book{}).
		Select([]string{"book_name", "book_number", "category_name",
			"publisher", "author", "price", "pdate", "pic_url"}).
		Where("id = ?", id).First(&bookModel)
	return bookModel, result.RowsAffected
}

// GetByName 书名获取
func GetByName(name string) (bookModel Book, row int64) {
	result := mysql.DB.Model(Book{}).
		Select([]string{"id", "book_name", "category_name",
			"publisher", "author", "introduce", "price", "pdate", "pic_url", "quantity"}).
		Where("book_name = ?", name).First(&bookModel)
	return bookModel, result.RowsAffected
}

// GetBySearchName 搜索书名获取
func GetBySearchName(name string) (books []Book, row int64) {
	result := mysql.DB.Model(Book{}).
		Select([]string{"id", "book_name", "category_name",
			"publisher", "author", "introduce", "price", "pdate", "pic_url"}).
		Where("book_name LIKE ?", "%"+name+"%").Find(&books)

	return books, result.RowsAffected
}

// GetByIsNewBook 查询为新书项的图书
func GetByIsNewBook(c *gin.Context, count string) (books []Book, page paginator.Page) {

	page = paginator.Paginate(
		c,
		mysql.DB.Model(Book{}).Where("is_new_book = ?", true),
		"client",
		"books/by-is_new_book/"+count,
		&books,
		count,
		"in_time",
		"desc")

	return
}

// GetByRecommended 查询为推荐项的图书
func GetByRecommended(c *gin.Context, count string) (books []Book, page paginator.Page) {

	page = paginator.Paginate(
		c,
		mysql.DB.Model(Book{}).Where("is_commended = ?", true),
		"client",
		"books/by-recommended/"+count,
		&books,
		count,
		"in_time",
		"desc")

	return
}

// GetBySold 查询销售排行图书
func GetBySold(c *gin.Context, count string) (books []Book, page paginator.Page) {

	page = paginator.Paginate(
		c,
		mysql.DB.Model(Book{}),
		"client",
		"books/by-sold/"+count,
		&books,
		count,
		"sold",
		"desc")

	return
}

// GetBooksBySlice 通过 id 切片获取图书
func GetBooksBySlice(ids []int64) ([]Book, int64) {
	var books []Book
	// 按给出切片顺序排序查询
	row := mysql.DB.Model(Book{}).
		Select([]string{"id", "book_name", "price", "pic_url", "publisher", "author", "quantity"}).
		Where(ids).Clauses(clause.OrderBy{
		Expression: clause.Expr{SQL: "FIELD(id,?)",
			Vars: []interface{}{ids}, WithoutParentheses: true},
	}).Find(&books).RowsAffected
	return books, row
}

// GetBooksAll 获取所有图书
func GetBooksAll(c *gin.Context, count string) (books []Book, page paginator.Page) {

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
	row := mysql.DB.Preload("Books", func(db *gorm.DB) *gorm.DB {
		return db.Model(Book{}).
			Select([]string{"id", "category_name", "book_name", "publisher"})
	}).Order("category_id asc").Find(&categories).RowsAffected
	return categories, row
}

// ChangeQuantityAndSold 修改库存销量
func ChangeQuantityAndSold(detailModel order.OrdersDetail) (bookModel Book) {
	mysql.DB.Model(Book{}).Where("id = ?", detailModel.BookID).Select("quantity", "sold").
		Updates(map[string]interface{}{
			"quantity": gorm.Expr("quantity - ?", detailModel.BuyCount),
			"sold":     gorm.Expr("sold + ?", detailModel.BuyCount),
		})

	mysql.DB.Model(Book{}).Where("id = ?", detailModel.BookID).First(&bookModel)
	return
}
