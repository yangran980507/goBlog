// Package admin 存放 admin 对 book 的操作
package admin

import (
	"blog/internal/server/models"
	"blog/internal/server/models/book"
	"blog/internal/server/requests"
	"blog/pkg/errcode"
	"blog/pkg/helps"
	"blog/pkg/helps/book_helps"
	"blog/pkg/logger"
	"blog/pkg/response"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

// BookStorage 图书入库
func (ac *AdminController) BookStorage(c *gin.Context) {
	request := requests.BookStorageValidation{}
	if ok := requests.BindAndValid(c, &request, requests.BookStorageValidate); !ok {
		return
	}

	bookModel := book.Book{
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

	if bookModel.ID == 0 {
		response.NewResponse(c, errcode.ErrUnknown.ParseCode()).
			WithResponse("入库失败，请稍后重试")
	} else {
		response.NewResponse(c, errcode.ErrSuccess.ParseCode()).
			WithResponse("入库成功")
	}
}

// GetBooksAllByPaginator 通过分类控制器获取入库图书信息
func (ac *AdminController) GetBooksAllByPaginator(c *gin.Context) {

	books := make([]book.Book, 10)
	books, page := book.GetBooksAll(c, 10)
	response.NewResponse(c, errcode.ErrSuccess.ParseCode()).WithResponse(gin.H{
		"data": books,
		"page": page,
	})
}

func (ac *AdminController) DeleteBook(c *gin.Context) {
	// 解析接口
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.LogIf(err)
		response.NewResponse(c, errcode.ErrUnknown.ParseCode()).WithResponse("删除失败")
		return
	}

	idB := models.BaseMode{
		ID: uint(id),
	}
	bookModel := book.Book{
		BaseMode: idB,
	}

	row := bookModel.Delete()
	if row != 1 {
		response.NewResponse(c, errcode.ErrUnknown.ParseCode()).WithResponse("删除失败")
	} else {
		response.NewResponse(c, errcode.ErrSuccess.ParseCode()).WithResponse("删除成功")
	}
}
