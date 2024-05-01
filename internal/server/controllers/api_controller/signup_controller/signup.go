package signup_controller

import (
	"blog/internal/server/controllers"
	"blog/internal/server/models/user"
	"blog/internal/server/requests"
	"blog/pkg/errcode"
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

	userCreated := user.User{
		LoginName: request.LoginName,
		TrueName:  request.TrueName,
		PassWord:  request.PassWord,
		Phone:     request.Phone,
	}
	userCreated.Create()

	if userCreated.ID > 0 {
		response.NewResponse(c, errcode.ErrSuccess.ParseCode()).WithResponse("注册成功")
	} else {
		response.NewResponse(c, errcode.ErrUnknown.ParseCode()).WithResponse("注册失败")
	}
}
