// Package client 用户授权 handlerFunc
package client

import (
	"blog/internal/server/models/user"
	"blog/internal/server/requests"
	"blog/pkg/auth"
	"blog/pkg/errcode"
	"blog/pkg/jwt"
	"blog/pkg/logger"
	"blog/pkg/response"
	"errors"
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
		Address:   request.Address,
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
		response.NewResponse(c, errcode.ErrTokenInvalid).
			WithResponse("注册失败")
	}
}

// LoginUser 登陆处理函数
func (uc *UserController) LoginUser(c *gin.Context) {

	// 创建验证结构体空值实例
	request := requests.SignInValidation{}
	if ok := requests.BindAndValid(c, &request, requests.SignInValidate); !ok {
		// 绑定 && 验证失败
		return
	}

	// 创建登陆控制器实例
	loginInfo := auth.LoginInfo{
		LoginName: request.LoginName,
		Password:  request.Password,
	}
	userModel, err := loginInfo.Login()
	if err != nil {

		// 登陆失败
		switch {
		case errors.Is(err, errcode.ErrAccountAbsent):
			response.NewResponse(c, errcode.ErrAccountAbsent).
				WithResponse("账户不存在")
		case errors.Is(err, errcode.ErrPassWord):
			response.NewResponse(c, errcode.ErrPassWord).
				WithResponse("密码输入错误")
		}

	} else {
		if allow := userModel.IsManager; allow {
			response.NewResponse(c, errcode.ErrNotAdmin).
				WithResponse("请使用用户账号登录!")
		} else {
			// 创建 jwt 鉴权结构体实例
			userinfo := jwt.UserInfo{
				UserID:        userModel.ID,
				UserLoginName: userModel.LoginName,
			}
			// 签发令牌
			token := jwt.NewJWT().IssueToken(userinfo)
			// 返回成功码，令牌及用户数据
			response.NewResponse(c, errcode.ErrSuccess, "登录成功!").
				WithResponse(
					gin.H{
						"user":  userModel,
						"token": token,
					})
		}
	}
}

// RefreshToken 刷新 Access token
func (uc *UserController) RefreshToken(c *gin.Context) {
	token, err := jwt.NewJWT().RefreshToken(c)

	if err != nil {
		logger.LogIf(err)
		if errors.Is(err, errcode.ErrTokenTimeOut) {
			response.NewResponse(c, errcode.ErrTokenTimeOut).WithResponse(
				"令牌刷新无效:" + err.Error())
			return
		}

		response.NewResponse(c, errcode.ErrTokenInvalid, "令牌刷新无效:"+err.Error()).
			WithResponse()
		return
	} else {
		// 返回成功码，令牌及用户数据
		response.NewResponse(c, errcode.ErrSuccess).WithResponse(
			gin.H{
				"token": token,
			})
	}
}
