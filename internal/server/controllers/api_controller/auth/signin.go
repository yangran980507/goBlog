// Package auth 用户登陆处理函数
package auth

import (
	"blog/internal/server/controllers"
	"blog/internal/server/requests"
	"blog/pkg/auth"
	"blog/pkg/errcode"
	"blog/pkg/jwt"
	"blog/pkg/response"
	"github.com/gin-gonic/gin"
)

// LoginController 用户登陆控制器
type LoginController struct {
	controllers.BaseController
}

// LoginUser 登陆处理函数
func (lc *LoginController) LoginUser(c *gin.Context) {
	// 创建验证结构体空值实例
	request := requests.SigninUserValidation{}
	if ok := requests.BindAndValid(c, &request, requests.SigninUserValidate); !ok {
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
		response.NewResponse(c, errcode.ErrTokenInvalid.ParseCode()).
			WithResponse("登陆失败:" + err.Error())
	} else {
		// 创建 jwt 鉴权结构体实例
		userinfo := jwt.UserInfo{
			UserID:        userModel.ID,
			UserLoginName: userModel.LoginName,
		}
		// 签发令牌
		token := jwt.NewJWT().IssueToken(userinfo)
		// 返回成功码，令牌及用户数据
		response.NewResponse(c, errcode.ErrSuccess.ParseCode()).WithResponse(
			gin.H{
				"user":  userModel,
				"token": token,
			})
	}
}

// RefreshToken 刷新 Access token
func (lc *LoginController) RefreshToken(c *gin.Context) {
	token, err := jwt.NewJWT().RefreshToken(c)

	if err != nil {
		response.NewResponse(c, errcode.ErrTokenInvalid.ParseCode()).WithResponse(
			"令牌刷新无效:" + err.Error())
	} else {
		// 返回成功码，令牌及用户数据
		response.NewResponse(c, errcode.ErrSuccess.ParseCode()).WithResponse(
			gin.H{
				"token": token,
			})
	}
}
