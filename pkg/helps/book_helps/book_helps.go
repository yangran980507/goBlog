// Package book_helps book 模型相关辅助函数
package book_helps

// RequestStrToBool 将请求中的字符串类型转换为布尔类型
func RequestStrToBool(str string) bool {
	var b bool
	switch str {
	case "true":
		b = true
	case "false":
		b = false
	}
	return b
}
