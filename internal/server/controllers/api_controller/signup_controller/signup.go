package signup_controller

import (
	"blog/internal/server/controllers"
	"blog/internal/server/models/user"
	"blog/internal/server/requests"
	"blog/pkg/errcode"
	"blog/pkg/jwt"
	"blog/pkg/response"
	"github.com/gin-gonic/gin"
)

type SignupController struct {
	controllers.BaseController
}

func (sc *SignupController) SignupUser(c *gin.Context) {

	request := requests.SignupUserValidation{}
	if ok := requests.BindAndValid(c, &request, requests.SignupUserValidate); !ok {
		return
	}

	userModel := user.User{
		LoginName: request.LoginName,
		TrueName:  request.TrueName,
		PassWord:  request.PassWord,
		Phone:     request.Phone,
	}
	userModel.Create()

	if userModel.ID > 0 {
		userinfo := jwt.UserInfo{
			UserID:        userModel.ID,
			UserLoginName: userModel.LoginName,
		}
		token := jwt.NewJWT().IssueToken(userinfo)
		response.NewResponse(c, errcode.ErrSuccess.ParseCode()).WithResponse(
			gin.H{
				"token": token,
				"data":  userModel,
			})
	} else {
		response.NewResponse(c, errcode.ErrTokenInvalid.ParseCode()).WithResponse("创建用户失败")
	}
}
