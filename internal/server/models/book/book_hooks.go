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
func (book *Book) AfterCreate(tx *gorm.DB) (err error) {
	category, _ := book.GetCategory()
	if err = book.AddAssociation(&category); err != nil {
		return
	}
	return nil
}


// BeforeUpdate 修改图书之前调用
func (book *Book) BeforeUpdate(tx *gorm.DB) (err error) {
	// 查询原数据
	bookOld, _ := book.Get()
	// 判断类型是否修改
	if bookOld.CategoryName != book.CategoryName {
		// 未修改
		return nil
	}

	// 修改
	// 判断类型是否存在
	_, err = book.GetCategory()
	if err != nil {
		// 类型不存在，创建新类型
		err = book.AddCategory()
		if err != nil {
			return err
		}
		return nil
	}

	// 获取原类型
	//category, _ := bookOld.GetCategory()
	// 删除原关联
	//category.DropAssociation(&bookOld)

	return nil
}

/*
// AfterUpdate 修改图书之后调用
func (book *Book) AfterUpdate(tx *gorm.DB) (err error) {
	category, _ := book.GetCategory()
	count := CountCategory()
	if int64(category.CategoryID) < count {
		// 新类型，添加关联
		if err = book.AddAssociation(&category); err != nil {
			return err
		}
		return nil
	}

	// 原类型，替换关联
	if err = category.ReplaceAssociation(book); err != nil {
		return err
	}
	return nil
}
*/
