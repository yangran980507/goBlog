// Package paginator 控制分页的函数
package paginator

import (
	"blog/global"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"math"
	"strconv"
)

// Paginator 分页控制器
type Paginator struct {
	baseURL      string       // 请求页码的链接
	ctx          *gin.Context // 请求上下文
	query        *gorm.DB     // 数据库句柄
	currentPage  int          // 当前页码
	countPerPage int          // 单页数据条数：请求中设置
	totalPages   int          // 总页数
	totalCount   int          // 数据库中数据条数
	offset       int          // 查询数据自动跳过行数
	sort         string       // 查询后排序column
	order        string       // 顺序："asc" 或者 "desc"
}

func (p *Paginator) initPaginator(controller, tableName string) {

	// 获取基础请求链接：`http://locaohost:8080/api/controller/tableName?page=`
	p.baseURL = fmt.Sprintf("%s/api/%s/%s?page=",
		global.AppSetting.URL, controller, tableName)

	p.totalCount = p.getCount()
	p.totalPages = p.getTotalPages()
	p.currentPage = p.getCurrentPage()

	p.offset = (p.currentPage - 1) * p.countPerPage
}

// 返回当前页码
func (p *Paginator) getCurrentPage() int {

	if p.totalCount == 0 {
		// 总页数等于0,没有数据
		return 0
	}

	// 获取用户请求的页码
	page, err := strconv.Atoi(p.ctx.Query("page"))
	if err != nil || page <= 0 {
		// 默认为1
		return 1
	}

	if page > p.totalPages {
		// 请求页数大于总页数
		return p.totalPages
	}

	return page
}

// 返回数据总条数
func (p *Paginator) getCount() int {
	var count int64
	if err := p.query.Count(&count).Error; err != nil {
		return 0
	}
	return int(count)
}

// 返回总页码数
func (p *Paginator) getTotalPages() int {

	// 无数据返回0
	if p.totalCount == 0 {
		return 0
	}

	// math.Cell 返回大于输入值的最小整数
	num := math.Ceil(float64(p.totalCount) / float64(p.countPerPage))

	// 数据数小于设置单页数据数，按一页
	if num == 0 {
		num = 1
	}

	return int(num)
}

// 返回查询页码链接
func (p *Paginator) getPageLink(page int) string {
	return fmt.Sprintf("%s%v", p.baseURL, page)
}

// 返回首页链接
func (p *Paginator) getFirstPage() string {
	return p.getPageLink(1)
}

// 返回尾页链接
func (p *Paginator) getLastPage() string {
	return p.getPageLink(p.totalPages)
}

// 返回上一页链接
func (p *Paginator) getPrevPage() string {
	if p.currentPage <= 1 || p.currentPage > p.totalPages {
		return ""
	}
	return p.getPageLink(p.currentPage - 1)
}

// 返回下一页链接
func (p *Paginator) getNextPage() string {
	if p.totalPages > p.currentPage {
		return p.getPageLink(p.currentPage + 1)
	}
	return ""
}
