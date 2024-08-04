// Package factories 生成 book 数据
package factories

import (
	"blog/internal/server/models/book"
	"blog/pkg/helps"
	"github.com/bxcodec/faker/v3"
	"time"
)

func MakeBooks(count int) []book.Book {

	var objs []book.Book

	// 设定唯一值
	faker.SetGenerateUniqueValues(true)

	for i := 0; i < count; i++ {
		model := book.Book{
			BookNumber:   faker.Phonenumber(),
			BookName:     faker.Word(),
			CategoryName: "类别" + helps.RandomNumber(1),
			Publisher:    helps.RandomPublisher(),
			Author:       faker.ChineseName(),
			Introduce:    faker.Paragraph(),
			Price:        helps.RandomPrice(),
			Pdate:        faker.UnixTime(),
			PicURL:       "0" + helps.RandomNumber(2) + ".jpg",
			InTime:       time.Now().Unix(),
			IsNewBook:    true,
			IsCommended:  true,
			Quantity:     helps.RandomQuantity(),
		}

		objs = append(objs, model)
	}
	return objs
}
