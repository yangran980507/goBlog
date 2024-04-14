package main

import (
	"blog/global"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	// 初始化 gin 实例
	r := gin.Default()

	// 设置路由
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "welcome to goBlog!",
		})
	})

	fmt.Printf("values: %s\n", global.ServerSetting)

	// 运行服务
	err := r.Run(":8080")
	if err != nil {
		fmt.Println(err.Error())
	}
}
