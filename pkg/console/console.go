// Package console 命令行高亮方法
package console

import (
	"fmt"
	"os"

	"github.com/mgutz/ansi"
)

// Success 打印成功消息，绿色
func Success(msg string) {
	colorOut(msg, "green")
}

// Error 打印报错消息，红色
func Error(msg string) {
	colorOut(msg, "red")
}

// Warn 打印警告消息，黄色
func Warn(msg string) {
	colorOut(msg, "yellow")
}

// Exit 打印报错消息，退出程序
func Exit(msg string) {
	colorOut(msg, "red")
	os.Exit(1)
}

// ExitIf 语法糖，自带 err != nil 判断
func ExitIf(err error) {
	if err != nil {
		Exit(err.Error())
	}
}

// 设置高亮
func colorOut(message, color string) {
	fmt.Fprintln(os.Stdout, ansi.Color(message, color))
}
