// Package client 用户注册处理函数
package client

import (
	"blog/internal/server/models/user"
	"blog/internal/server/requests"
	"blog/pkg/errcode"
	"blog/pkg/jwt"
	"blog/pkg/response"
	"github.com/gin-gonic/gin"
)

// SignupUser 用户注册处理函数
func (uc *UserController) SignupUser(c *gin.Context) {

	// 创建验证结构体空值实例
	request := requests.SignupUserValidation{}
	// 绑定结构数据到验证结构体中
	if ok := requests.BindAndValid(c, &request, requests.SignupUserValidate); !ok {
		// 绑定 && 验证失败
		return
	}

	// 创建写数据库结构体实例
	userModel := user.User{
		LoginName: request.LoginName,
		TrueName:  request.TrueName,
		PassWord:  request.PassWord,
		Phone:     request.Phone,
	}
	// 调用写数据库函数创建数据
	userModel.Create()

	if userModel.ID > 0 {
		// 创建 jwt 鉴权结构体实例
		userinfo := jwt.UserInfo{
			UserID:        userModel.ID,
			UserLoginName: userModel.LoginName,
		}
		// 签发令牌
		token := jwt.NewJWT().IssueToken(userinfo)
		// 返回成功码，令牌及用户数据
		response.NewResponse(c, errcode.ErrSuccess).WithResponse(
			gin.H{
				"token": token,
				"user":  userModel,
			})
	} else {
		// 创建失败，返回失败信息
		response.NewResponse(c, errcode.ErrTokenInvalid, "注册失败").
			WithResponse()
	}
}
