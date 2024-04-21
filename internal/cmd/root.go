// Package cmd 命令行根命令
package cmd

import (
	"blog/global"
	"blog/global/logger"
	"blog/pkg/console"
	"blog/pkg/errcode"
	"blog/pkg/mysql"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var RootCmd = &cobra.Command{
	Use:   "Blog",
	Short: "A simple project",
	Long:  "Default will run webserver",

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// 初始化配置变量
		global.InitializeConf()
		// 初始化数据库连接
		mysql.InitializeDB()
		// 初始化业务错误码
		errcode.InitializeErrorCode()
		// 初始化日志
		logger.InitializeLog()
	},
}

func Execute() {

	// 添加子命令
	addCmd()

	// 默认 rootCmd 执行 ServeCmd 命令
	defaultCmd(RootCmd, ServeCmd)

	if err := RootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("Failed to run app with %v: %s",
			os.Args, err.Error()))
	}
}

// 添加子命令
func addCmd() {
	RootCmd.AddCommand(
		ServeCmd,
	)
}

// 默认执行 ServeCmd 命令
func defaultCmd(rootCmd *cobra.Command, subCmd *cobra.Command) {
	cmd, _, err := rootCmd.Find(os.Args[1:])
	firstArg := firstElement(os.Args[1:])
	if err == nil && cmd.Use == rootCmd.Use && firstArg != "-h" && firstArg != "--help" {
		args := append([]string{subCmd.Use}, os.Args[1:]...)
		rootCmd.SetArgs(args)
	}
}

func firstElement(args []string) string {
	if len(args) > 0 {
		return args[0]
	}
	return ""
}
