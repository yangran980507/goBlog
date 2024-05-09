// Package admin 图书入库handler
package admin

import (
	"blog/internal/server/models/books"
	"blog/internal/server/requests"
	"blog/pkg/errcode"
	"blog/pkg/helps"
	"blog/pkg/helps/book_helps"
	"blog/pkg/response"
	"github.com/gin-gonic/gin"
	"time"
)

func (ac *AdminController) BookStorage(c *gin.Context) {
	request := requests.BookStorageValidation{}
	if ok := requests.BindAndValid(c, &request, requests.BookStorageValidate); !ok {
		return
	}

	bookModel := books.Book{
		BookNumber:  request.BookNumber,
		BookName:    request.BookName,
		BookType:    request.BookType,
		Publisher:   request.Publisher,
		Author:      request.Author,
		Introduce:   request.Introduce,
		Price:       request.Price,
		Pdate:       helps.StrToTimeUnix(request.Pdate),
		PicURL:      "../../assets/images/" + request.PicURL,
		InTime:      time.Now().Unix(),
		IsNewBook:   book_helps.RequestStrToBool(request.IsNewBook),
		IsCommended: book_helps.RequestStrToBool(request.IsCommended),
	}

	bookModel.Create()

	if bookModel.ID > 0 {
		response.NewResponse(c, errcode.ErrSuccess.ParseCode()).
			WithResponse("入库成功")
	} else {
		response.NewResponse(c, errcode.ErrUnknown.ParseCode()).
			WithResponse("入库失败，请稍后重试")
	}
}
