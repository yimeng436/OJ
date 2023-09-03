package controller

import (
	"gatewaysvr/log"
	"gatewaysvr/rpcservice"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/yimeng436/OJ/common"
	"github.com/yimeng436/OJ/common/constant"
	"github.com/yimeng436/OJ/pkg/pb"
)

// @Summary		用户登录
// @Description	用户登录
// @Accept json
// @Param user body pb.UserLoginRequest true "user"
// @Produce  json
// @Success		200	{object}	common.Response
// @Router			/user/login  [post]
func Login(ctx *gin.Context) {
	var request pb.UserLoginRequest

	if err := ctx.ShouldBind(&request); err != nil {
		log.Fatal(err)
		common.Fail(ctx, "参数错误")
		return
	}
	loginUser := &pb.UserLoginRequest{
		UserName:     request.UserName,
		UserPassword: request.UserPassword,
	}
	userServiceClient := rpcservice.GetUserServiceClient()
	userVo, err := userServiceClient.UserLogin(ctx, loginUser)
	if err != nil {
		log.Fatal(err)
		common.Fail(ctx, err.Error())
		return
	}
	session := sessions.Default(ctx)
	session.Set(constant.UserLoginState, userVo)
	err = session.Save()
	common.Success(ctx, userVo, "success")
}

// @Summary		获取登录用户
// @Description	获取登录用户
// @Produce  json
// @Success		200	{object}	common.Response
// @Router			/user/getLoginUser  [get]
func GetLoginUser(ctx *gin.Context) {
	CheckLogin(ctx)
	loginUser, _ := ctx.Get("loginUser")
	common.Success(ctx, loginUser, "success")
}

func CheckLogin(ctx *gin.Context) {
	session := sessions.Default(ctx)
	loginUser := session.Get(constant.UserLoginState)
	if loginUser == nil {
		common.Fail(ctx, "未登录")
		ctx.Abort()
	}
	ctx.Set("loginUser", loginUser)
	ctx.Next()
}
