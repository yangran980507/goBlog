// Package book 图书模型
package book

import (
	"blog/internal/server/models"
	"blog/pkg/mysql"
)

// Book 图书模型
type Book struct {
	// 图书编号
	models.BaseMode
	// 书号：
	BookNumber string `json:"book_number,omitempty" gorm:"column:book_number;index;unique"`
	// 书名
	BookName string `json:"book_name,omitempty"`
	// 图书类型 ID
	CategoryName string `json:"category_name" gorm:"column:category_name;not null"`
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
	Selled int `json:"selled,omitempty"`
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

// Get 显示图书
func (book *Book) Get() (Book, int64) {
	bookModel := Book{}
	row := mysql.DB.First(&bookModel, book.ID).RowsAffected
	return bookModel, row
}

// Update 更新图书
func (book *Book) Update() int64 {
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
			"is_new_book":   book.IsNewBook,
			"is_commended":  book.IsCommended,
		}).
		RowsAffected
}

// Category 图书类别模型
type Category struct {
	// 分类编号
	CategoryID uint `json:"category_id" gorm:"column:category_id;primaryKey;autoIncrement"`
	// 类别名
	Name string `json:"name" gorm:"column:name"`
	// 对应图书
	Books []Book `json:"books" gorm:"foreignKey:category_name;references:name"`
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

// DeleteCategory 删除分类
func (category *Category) DeleteCategory() {
	mysql.DB.Delete(&category)
}

// CountAssociation 关联计数
func (book *Book) CountAssociation(category *Category) bool {

	return mysql.DB.Model(category).
		Association("Books").Count() == 0
}

/*
// AddAssociation 添加关联
func (book *Book) AddAssociation(category *Category) error {

	if err := mysql.DB.Model(category).Association("Books").
		Append(&book); err != nil {
		return err
	}

	return nil
}


// FindAssociation 查找关联
func (category *Category) FindAssociation(book *[]Book) error {

	if err := mysql.DB.Model(category).Association("Books").
		Find(&book); err != nil {
		return err
	}

	return nil
}

// ReplaceAssociation 替换关联
func (category *Category) ReplaceAssociation(book *Book) error {
	if err := mysql.DB.Model(category).Association("Books").
		Replace(&book); err != nil {
		return err
	}
	return nil
}

// DropAssociation 删除关联
func (category *Category) DropAssociation(book *Book) {
	mysql.DB.Model(category).Association("Books").
		Delete(&book)

}
*/
