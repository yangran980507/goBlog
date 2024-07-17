// Package factories 生成 book 数据
package factories

import (
	"blog/internal/server/models/book"
	"blog/pkg/helps"
	"github.com/bxcodec/faker/v3"
	"strconv"
	"time"
)

func MakeBooks(count int) []book.Book {

	var objs []book.Book

	// 设定唯一值
	faker.SetGenerateUniqueValues(true)

	nums, _ := strconv.Atoi(helps.RandomNumber(2))

	for i := 0; i < count; i++ {
		model := book.Book{
			BookNumber:   faker.Phonenumber(),
			BookName:     faker.Name(),
			CategoryName: "类别8",
			Publisher:    faker.Name(),
			Author:       faker.Username(),
			Introduce:    faker.Paragraph(),
			Price:        float64(nums),
			Pdate:        0,
			PicURL:       faker.Name(),
			InTime:       time.Now().Unix(),
			IsNewBook:    true,
			IsCommended:  true,
			Quantity:     nums,
		}

		objs = append(objs, model)
	}
	return objs
}
