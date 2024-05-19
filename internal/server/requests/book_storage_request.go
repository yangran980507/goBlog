// Package requests 图书创建验证
package requests

import (
	"github.com/thedevsaddam/govalidator"
)

type BookStorageValidation struct {

	// 书号：
	BookNumber string `json:"book_number" valid:"book_number"`
	// 书名
	BookName string `json:"book_name" valid:"book_name"`
	// 图书类型
	BookType string `json:"book_type" valid:"book_type"`
	// 出版社
	Publisher string `json:"publisher" valid:"publisher"`
	// 作者
	Author string `json:"author" valid:"author"`
	// 简介
	Introduce string `json:"introduce" valid:"introduce"`
	// 价格
	Price float64 `json:"price" valid:"price"`
	// 出版日期
	Pdate string `json:"pdate" valid:"pdate"`
	// 封面路径
	PicURL string `json:"picURL" valid:"picURL"`
	// 是否新书 ？
	IsNewBook string `json:"isNewBook" valid:"isNewBook"`
	// 是否推荐 ？
	IsCommended string `json:"isCommended" valid:"isCommended"`
}

// BookStorageValidate 图书入库验证函数
func BookStorageValidate(data interface{}) map[string][]string {
	// 验证规则
	rules := govalidator.MapData{
		"book_number": []string{"required", "exists:books-book_number"},
		"book_name":   []string{"required"},
		"book_type":   []string{"required"},
		"publisher":   []string{"required"},
		"author":      []string{"required"},
		"introduce":   []string{"required"},
		"price":       []string{"required"},
		"pdate":       []string{"required"},
		"isNewBook":   []string{"required"},
		"isCommended": []string{"required"},
	}

	// 返回错误信息
	messages := govalidator.MapData{

		"book_number": []string{
			"required: 书号为必填",
		},
		"book_name":   []string{"required: 书名为必填"},
		"book_type":   []string{"required: 类别名称为必填"},
		"publisher":   []string{"required: 出版社为必填"},
		"author":      []string{"required: 作者为必填"},
		"introduce":   []string{"required: 图书简介为必填"},
		"price":       []string{"required: 价格为必填"},
		"pdate":       []string{"required: 发行日期为必填"},
		"isNewBook":   []string{"required: 新书选项为必选"},
		"isCommended": []string{"required: 推荐选项为必选"},
	}

	// 传入设置的验证规则，错误消息参数，返回错误信息
	return validate(data, rules, messages)
}

type BookUpdateValidation struct {

	// 书名
	BookName string `json:"book_name" valid:"book_name"`
	// 图书类型
	BookType string `json:"book_type" valid:"book_type"`
	// 出版社
	Publisher string `json:"publisher" valid:"publisher"`
	// 作者
	Author string `json:"author" valid:"author"`
	// 简介
	Introduce string `json:"introduce" valid:"introduce"`
	// 价格
	Price float64 `json:"price" valid:"price"`
	// 出版日期
	Pdate string `json:"pdate" valid:"pdate"`
	// 封面路径
	PicURL string `json:"picURL" valid:"picURL"`
	// 是否新书 ？
	IsNewBook string `json:"isNewBook" valid:"isNewBook"`
	// 是否推荐 ？
	IsCommended string `json:"isCommended" valid:"isCommended"`
}

// BookUpdateValidate 图书修改验证函数
func BookUpdateValidate(data interface{}) map[string][]string {
	// 验证规则
	rules := govalidator.MapData{
		"book_name":   []string{"required"},
		"book_type":   []string{"required"},
		"publisher":   []string{"required"},
		"author":      []string{"required"},
		"introduce":   []string{"required"},
		"price":       []string{"required"},
		"pdate":       []string{"required"},
		"isNewBook":   []string{"required"},
		"isCommended": []string{"required"},
	}

	// 返回错误信息
	messages := govalidator.MapData{

		"book_name":   []string{"required: 书名为必填"},
		"book_type":   []string{"required: 类别名称为必填"},
		"publisher":   []string{"required: 出版社为必填"},
		"author":      []string{"required: 作者为必填"},
		"introduce":   []string{"required: 图书简介为必填"},
		"price":       []string{"required: 价格为必填"},
		"pdate":       []string{"required: 发行日期为必填"},
		"isNewBook":   []string{"required: 新书选项为必选"},
		"isCommended": []string{"required: 推荐选项为必选"},
	}

	// 传入设置的验证规则，错误消息参数，返回错误信息
	return validate(data, rules, messages)
}
