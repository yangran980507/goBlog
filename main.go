package main

import (
	"blog/global"
	"blog/global/logger"
	"blog/global/mysql"
	"blog/pkg/errcode"
	"blog/pkg/response"
	"fmt"

	"github.com/gin-gonic/gin"
)

func init() {
	mysql.InitializeDB()
	errcode.InitializeErrorCode()
	logger.InitializeLog()
}
func main() {

	// 初始化 gin 实例
	r := gin.Default()

	// 设置路由
	r.GET("/", func(c *gin.Context) {
		response.NewResponse(c, errcode.ErrTest.ParseCode()).WithResponse(gin.H{
			"name": "yangran",
			"age":  25,
		}, "test01", "123456789")
	})

	// 运行服务
	err := r.Run(":" + global.ServerSetting.HttpPort)
	if err != nil {
		fmt.Println(err.Error())
	}
}
