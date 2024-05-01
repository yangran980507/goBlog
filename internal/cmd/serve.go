// Package cmd 服务开启命令行
package cmd

import (
	"blog/global"
	blogRouter "blog/initialize/router"
	"blog/pkg/console"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"net/http"
	"time"
)

var ServeCmd = &cobra.Command{
	Use:   "server",
	Short: "Start a web server",
	Run:   runWeb,
	Args:  cobra.NoArgs,
}

func runWeb(cmd *cobra.Command, args []string) {

	// 设置 gin.mode 为 release 模式
	gin.SetMode(global.ServerSetting.RunMode)

	// 初始化 gin 实例
	router := gin.New()

	//初始化 router
	blogRouter.SetupRouter(router)

	srv := &http.Server{
		Addr:              ":" + global.ServerSetting.HttpPort,
		Handler:           router,
		ReadHeaderTimeout: global.ServerSetting.ReadTimeOut * time.Second,
		WriteTimeout:      global.ServerSetting.WriteTimeOut * time.Second,
	}

	// 启动服务监听 :8080 端口
	if err := srv.ListenAndServe(); err != nil {
		// 发生错误时打印错误消息，退出程序 os.Exit(1)
		console.Exit("start web server failed,err:" + fmt.Sprint(err.Error()))
	}
}
