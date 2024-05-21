// Package book 存放 book 模型钩子函数
package book

import (
	// "gorm.io/gorm"
	"gorm.io/gorm"
)

// BeforeCreate 创建图书之前调用
func (book *Book) BeforeCreate(tx *gorm.DB) (err error) {

	// 查询类别是否存在
	_, err = book.GetCategory()
	if err != nil {
		// 该类型不存在
		err = book.AddCategory()
		if err != nil {
			return err
		}
		return nil
	}
	// 该类型存在
	return nil
}

/*
// AfterCreate 创建图书之后调用
// 将关联写进分类
func (book *Book) AfterCreate(tx *gorm.DB) (err error) {
	categoryModel, err := book.GetCategory()
	err = book.Append(categoryModel)
	if err != nil {
		return err
	}
	return nil
}
*/
