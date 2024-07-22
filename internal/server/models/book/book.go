// Package book 图书模型
package book

import (
	"blog/pkg/mysql"
)

// Carts 接受切片请求 JSON
type Carts struct {
	Books []Book `json:"books"`
}

// Book 图书模型
type Book struct {
	// 图书编号
	ID uint `json:"id,omitempty" gorm:"column:id;primaryKey;autoIncrement"`
	// 书号：
	BookNumber string `json:"book_number,omitempty" gorm:"column:book_number;index;unique"`
	// 书名
	BookName string `json:"book_name,omitempty" gorm:"column:book_name;index;unique"`
	// 图书类型 ID
	CategoryName string `json:"category_name,omitempty" gorm:"column:category_name;not null"`
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
	PicURL string `json:"pic_url,omitempty"`
	// 入库时间
	InTime int64 `json:"in_time,omitempty"`
	// 是否新书 ？
	IsNewBook bool `json:"is_new_book,omitempty"`
	// 是否推荐 ？
	IsCommended bool `json:"is_commended,omitempty"`
	// 库存
	Quantity int `json:"quantity,omitempty"`
	// 已售
	Sold int `json:"sold,omitempty"`
}

// Create 创建图书
func (book *Book) Create() error {
	if err := mysql.DB.Create(&book).Error; err != nil {
		return err
	}
	return nil
}

// Delete 删除图书
func (book *Book) Delete() int64 {

	return mysql.DB.Delete(Book{}, book.ID).RowsAffected
	/*bookModel := Book{}
	// 通过书号获取图书编号
	mysql.DB.Where("book_number = ?", book.BookNumber).First(&bookModel)
	row := mysql.DB.Delete(&Book{}, bookModel.ID).RowsAffected
	return row*/
}

// Get 获取图书
func (book *Book) Get() (Book, int64) {
	bookModel := Book{}
	row := mysql.DB.Select([]string{"id",
		"book_number",
		"book_name",
		"category_name",
		"publisher",
		"author",
		"introduce",
		"price",
		"pdate",
		"pic_url",
		"is_new_book",
		"is_commended",
		"quantity",
	}).First(&bookModel, book.ID).RowsAffected
	return bookModel, row
}

// Update 更新图书
func (book *Book) Update() error {
	return mysql.DB.Table("books").
		Where("id = ?", book.ID).
		Omit("id", "book_number", "in_time", "quantity", "selled").
		Updates(map[string]interface{}{
			"book_name":     book.BookName,
			"category_name": book.CategoryName,
			"publisher":     book.Publisher,
			"author":        book.Author,
			"introduce":     book.Introduce,
			"price":         book.Price,
			"pdate":         book.Pdate,
			"pic_url":       book.PicURL,
			"quantity":      book.Quantity,
			"is_new_book":   book.IsNewBook,
			"is_commended":  book.IsCommended,
		}).Error

}

// Category 图书类别模型
type Category struct {
	// 分类编号
	CategoryID uint `json:"category_id,omitempty" gorm:"column:category_id;primaryKey;autoIncrement"`
	// 类别名
	Name string `json:"name,omitempty" gorm:"column:name"`
	// 对应图书
	Books []Book `json:"books,omitempty" gorm:"foreignKey:category_name;references:name"`
}

// AddCategory 添加分类
func (book *Book) AddCategory() error {
	return mysql.DB.Model(&Category{}).Select("name").
		Create(map[string]interface{}{"name": book.CategoryName}).Error
}

// GetCategory 获取分类
func (book *Book) GetCategory() (Category, error) {
	categoryModel := Category{}
	err := mysql.DB.Model(&Category{}).Where("name = ?", book.CategoryName).
		First(&categoryModel).Error
	return categoryModel, err
}

// GetCategory 获取所有分类项
func GetCategory() ([]Category, error) {
	var categoryModel []Category
	err := mysql.DB.Model(&Category{}).Select("name").
		Find(&categoryModel).Error
	return categoryModel, err
}

// DeleteCategory 删除分类
func (category *Category) DeleteCategory() {
	mysql.DB.Delete(&category)
}

// CountAssociation 关联计数
func (book *Book) CountAssociation(category *Category) bool {

	return mysql.DB.Model(category).
		Association("Books").Count() == 0
}
