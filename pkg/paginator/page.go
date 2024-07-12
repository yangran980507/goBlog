// Package paginator 生成分页数据
package paginator

import (
	"blog/pkg/logger"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strconv"
)

// Page 分页数据
type Page struct {
	CurrentPage  int    // 当前页数
	PrevPageURL  string // 前一页链接
	NextPageURL  string // 后一页链接
	FirstPageURL string // 首页链接
	LastPageURL  string // 尾页链接
	//CountPerPage int    // 数据每页数
	//TotalPages   int    // 总页数
	//TotalCount   int    // 数据总数
}

// Paginate 分页数据生成
func Paginate(c *gin.Context, db *gorm.DB, controller string, tableName string,
	data interface{}, countStr string, sortBy string, orderBy string) Page {

	count, err := strconv.Atoi(countStr)
	if err != nil {
		return Page{}
	}

	p := &Paginator{
		ctx:          c,
		query:        db,
		sort:         sortBy,
		order:        orderBy,
		countPerPage: count,
	}
	p.initPaginator(controller, tableName)
	err = p.query.
		Preload(clause.Associations).  // 预加载全部关联
		Order(p.sort + " " + p.order). // 升序/降序排序
		Limit(p.countPerPage).         // 查询数
		Offset(p.offset).              // 查询跳过数
		Find(data).                    // 查询结果返回
		Error                          // 错误

	if err != nil {
		logger.LogIf(err)
		return Page{}
	}

	return Page{
		CurrentPage:  p.currentPage,
		FirstPageURL: p.getFirstPage(),
		PrevPageURL:  p.getPrevPage(),
		NextPageURL:  p.getNextPage(),
		LastPageURL:  p.getLastPage(),
	}
}
