// Package static 应用控制层
package static

import (
	"blog/internal/server/controllers/client"
	"github.com/gin-gonic/gin"
)

type StaticController struct {
	client.BaseAPIController
}

func (sc *StaticController) HomePage(c *gin.Context) {

	c.HTML(200, "主页.html", nil)
}
