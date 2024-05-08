// Package helps 公共辅助函数
package helps

import (
	"blog/pkg/logger"
	"time"
)

// StrToTimeUnix 字符串格式日期转时间戳
func StrToTimeUnix(dateStr string) int64 {
	date, err := time.Parse("2006-01", dateStr)
	logger.LogIf(err)
	return date.Unix()
}
