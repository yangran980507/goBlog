package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "welcome to goBlog!",
		})
	})

	err := r.Run(":8080")
	if err != nil {
		fmt.Println(err.Error())
	}
}
