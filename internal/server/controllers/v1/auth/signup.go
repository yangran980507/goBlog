// Package auth 注册相关逻辑
package auth

import (
	"blog/internal/server/controllers/v1"
	"blog/pkg/errcode"
	"blog/pkg/response"
	"github.com/gin-gonic/gin"
)

// SignupController 注册相关控制器
type SignupController struct {
	v1.BaseAPIController
}

func (sc *SignupController) IsUserExist(c *gin.Context) {
	response.NewResponse(c, errcode.ErrTest.ParseCode()).WithResponse(gin.H{
		"name": "yangran",
		"age":  25,
	})
}
