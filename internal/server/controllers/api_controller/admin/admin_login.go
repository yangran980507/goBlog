// Package admin 管理员登录 handlerFunc
package admin

import (
	"blog/internal/server/requests"
	"blog/pkg/auth"
	"blog/pkg/errcode"
	"blog/pkg/jwt"
	"blog/pkg/response"
	"errors"
	"github.com/gin-gonic/gin"
)

func (ac *AdminController) LoginAdmin(c *gin.Context) {
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
			response.NewResponse(c, errcode.ErrAccountAbsent, "账户不存在").
				WithResponse()
		case errors.Is(err, errcode.ErrPassWord):
			response.NewResponse(c, errcode.ErrPassWord, "密码输入错误").
				WithResponse()
		}
	} else {
		if allow := userModel.IsManager; !allow {
			response.NewResponse(c, errcode.ErrTokenInvalid, "请使用管理员账号登录!").
				WithResponse()
		} else {
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
				})
		}
	}
}
