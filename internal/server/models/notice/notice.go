// Package notice 公告模型
package notice

import (
	"blog/internal/server/models"
	"blog/pkg/mysql"
)

type Notice struct {
	// 公告编号
	models.BaseMode
	// 标题
	Title string
	// 正文
	Content string
	// 发布时间
	ShowTime int64
}

// Create 添加公告
func (notice *Notice) Create() {
	mysql.DB.Create(&notice)
}

// Delete 删除公告
func (notice *Notice) Delete() int64 {
	return mysql.DB.Model(Notice{}).Delete(&notice).RowsAffected
}

// Get 获取公告
func Get() ([]Notice, int64) {
	notices := make([]Notice, 5)
	rows := mysql.DB.Order("show_time  desc").Find(&notices).RowsAffected
	return notices, rows
}
