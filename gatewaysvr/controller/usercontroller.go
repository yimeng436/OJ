package controller

import (
	"OJ/common"
	"OJ/common/constant"
	"OJ/common/log"
	"OJ/common/model/request"
	"OJ/serivice"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// @Summary		用户登录
// @Description	用户登录
// @Accept json
// @Param user body request.UserLoginRequest true "user"
// @Produce  json
// @Success		200	{object}	common.Response
// @Router			/user/login  [post]
func Login(ctx *gin.Context) {

	var loginuser request.UserLoginRequest
	if err := ctx.ShouldBind(&loginuser); err != nil {
		log.Fatal(err)
		common.Fail(ctx, "参数错误")
		return
	}

	userVo, err := serivice.UserLogin(loginuser.UserAccount, loginuser.UserPassword)
	if err != nil {
		log.Fatal(err)
		common.Fail(ctx, err.Error())
		return
	}

	common.Success(ctx, userVo, "success")
}

// @Summary		获取登录用户
// @Description	获取登录用户
// @Produce  json
// @Success		200	{object}	common.Response
// @Router			/user/getLoginUser  [get]
func GetLoginUser(ctx *gin.Context) {
	session := sessions.Default(ctx)
	loginUser := session.Get(constant.UserLoginState)
	if loginUser == nil {
		common.Fail(ctx, "未登录")
		return
	}
	common.Success(ctx, loginUser, "success")
}
