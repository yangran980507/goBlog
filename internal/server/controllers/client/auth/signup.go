// Package auth 注册相关逻辑
package auth

import (
	"blog/internal/server/controllers/client"
	//"blog/pkg/errcode"
	//"blog/pkg/response"
	"github.com/gin-gonic/gin"
)

// SignupController 注册相关控制器
type SignupController struct {
	client.BaseAPIController
}

func (sc *SignupController) IsUserExist(c *gin.Context) {
	//c.String(200, "signal arrived")
	c.HTML(200, "主页.user_view", gin.H{})
}
