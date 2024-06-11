// Package admin 管理员图书 handlerFunc
package admin

import (
	"blog/internal/server/models"
	"blog/internal/server/models/book"
	"blog/internal/server/requests"
	"blog/pkg/app"
	"blog/pkg/errcode"
	"blog/pkg/helps"
	"blog/pkg/logger"
	"blog/pkg/response"
	"github.com/gin-gonic/gin"
	"time"
)

// BookStorage 图书入库
func (ac *AdminController) BookStorage(c *gin.Context) {
	request := requests.BookStorageValidation{}
	if ok := requests.BindAndValid(c, &request, requests.BookStorageValidate); !ok {
		return
	}

	bookModel := book.Book{
		BookNumber:   request.BookNumber,
		BookName:     request.BookName,
		CategoryName: request.CategoryName,
		Publisher:    request.Publisher,
		Author:       request.Author,
		Introduce:    request.Introduce,
		Price:        request.Price,
		Pdate:        helps.StrToTimeUnix(request.Pdate),
		PicURL:       "../../assets/images/" + request.PicURL,
		InTime:       time.Now().Unix(),
		Quantity:     request.Quantity,
		IsNewBook:    request.IsNewBook,
		IsCommended:  request.IsCommended,
	}

	err := bookModel.Create()

	if err != nil {
		logger.LogIf(err)
		response.NewResponse(c, errcode.ErrServer, "服务器错误：入库失败").
			WithResponse()
	} else {
		response.NewResponse(c, errcode.ErrSuccess, "入库成功").
			WithResponse()
	}
}

// GetBooksAllByPaginator 通过分类控制器获取入库图书信息
func (ac *AdminController) GetBooksAllByPaginator(c *gin.Context) {

	books := make([]book.Book, 10)
	books, page := book.GetBooksAll(c, 10)
	response.NewResponse(c, errcode.ErrSuccess).WithResponse(gin.H{
		"books": books,
		"page":  page,
	})
}

// DeleteBook 删除库中图书
func (ac *AdminController) DeleteBook(c *gin.Context) {

	// 解析接口数据
	id := app.GetIDFromAPI(c, "id")

	// 图书编号实例
	bookModel := book.Book{
		BaseMode: models.BaseMode{ID: uint(id)},
	}

	row := bookModel.Delete()
	if row == 1 {
		response.NewResponse(c, errcode.ErrSuccess, "删除成功").
			WithResponse()
	} else {
		response.NewResponse(c, errcode.ErrServer, "服务器错误：删除失败").
			WithResponse()
	}
}

// GetBook 获取图书信息
func (ac *AdminController) GetBook(c *gin.Context) {

	// 解析接口数据
	id := app.GetIDFromAPI(c, "id")

	// 图书编号实例
	bookModel := book.Book{
		BaseMode: models.BaseMode{ID: uint(id)},
	}

	b, row := bookModel.Get()
	if row == 1 {
		response.NewResponse(c, errcode.ErrSuccess).WithResponse(gin.H{
			"book": b,
		})
	} else {
		response.NewResponse(c, errcode.ErrServer, "服务器出错，请稍后重试").
			WithResponse()
	}

}

// BookUpdate 图书信息修改
func (ac *AdminController) BookUpdate(c *gin.Context) {

	request := requests.BookStorageValidation{}
	if ok := requests.BindAndValid(c, &request, requests.BookUpdateValidate); !ok {
		return
	}

	// 解析接口数据
	id := app.GetIDFromAPI(c, "id")

	bookModel := book.Book{
		BaseMode:     models.BaseMode{ID: uint(id)},
		BookName:     request.BookName,
		CategoryName: request.CategoryName,
		Publisher:    request.Publisher,
		Author:       request.Author,
		Introduce:    request.Introduce,
		Price:        request.Price,
		Pdate:        helps.StrToTimeUnix(request.Pdate),
		PicURL:       "../../assets/images/" + request.PicURL,
		Quantity:     request.Quantity,
		IsNewBook:    request.IsNewBook,
		IsCommended:  request.IsCommended,
	}

	var row int64

	bookOld, _ := bookModel.Get()
	// 判断类型是否修改
	if bookOld.CategoryName != bookModel.CategoryName {
		// 修改,判断类型是否存在
		_, err := bookModel.GetCategory()
		if err != nil {
			// 类型不存在，创建新类型
			err = bookModel.AddCategory()
			if err != nil {
				response.NewResponse(c, errcode.ErrServer, "服务器出错：修改失败").
					WithResponse()
				return
			}
		}
	}

	row = bookModel.Update()

	if row == 1 {
		response.NewResponse(c, errcode.ErrSuccess, "修改成功").
			WithResponse()
	} else {

		response.NewResponse(c, errcode.ErrSuccess, "未作任何修改").
			WithResponse()
	}

	// 获取修改前分类
	categoryModel, _ := bookOld.GetCategory()
	if bookOld.CountAssociation(&categoryModel) {
		// 删除修改前分类
		categoryModel.DeleteCategory()
	}
}
