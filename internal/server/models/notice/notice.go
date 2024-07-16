// Package notice 公告模型
package notice

import (
	"blog/pkg/mysql"
)

type Notice struct {
	// 公告编号
	ID uint `json:"id,omitempty" gorm:"column:id;primaryKey;autoIncrement"`
	// 正文
	Content string `json:"content,omitempty"`
	// 发布时间
	ShowTime int64 `json:"show_time,omitempty"`
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
func Get() (notices []Notice, rows int64) {
	rows = mysql.DB.Order("show_time  desc").Find(&notices).RowsAffected
	return notices, rows
}

// ClientGet 获取公告
func ClientGet() (notices []Notice, rows int64) {
	rows = mysql.DB.Select([]string{"content"}).Order("show_time  desc").Find(&notices).RowsAffected
	return notices, rows
}
