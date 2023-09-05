package controller

import (
	"encoding/json"
	"gatewaysvr/log"
	"gatewaysvr/rpcservice"
	"github.com/gin-gonic/gin"
	"github.com/yimeng436/OJ/common"
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
	// 前端一致请求不会携带 cookie,换成用简单版的token实现的
	//session := sessions.Default(ctx)
	//session.Set(constant.UserLoginState, userVo)
	//err = session.Save()
	//userState, err := json.Marshal(userVo)
	//ctx.SetCookie(constant.UserLoginState, string(userState), 3600, "/", "localhost", false, true)
	common.Success(ctx, userVo, "success")
}

// @Summary		获取登录用户
// @Description	获取登录用户
// @Produce  json
// @Success		200	{object}	common.Response
// @Router			/user/getLoginUser  [get]
func GetLoginUser(ctx *gin.Context) {
	CheckLogin(ctx)
	loginUser, e := ctx.Get("loginUser")
	if !e {
		common.Fail(ctx, "未登录")
		ctx.Abort()
		return
	}
	common.Success(ctx, loginUser, "success")
}

func CheckLogin(ctx *gin.Context) {

	token := ctx.Request.Header.Get("Authorization")

	if token == "" {
		ctx.Abort()
		return
	}
	token = token[7:]
	loginUser := &pb.UserVo{}
	err := json.Unmarshal([]byte(token), loginUser)
	if err != nil {
		ctx.Abort()
		return
	}
	if loginUser.Id == 0 {
		ctx.Abort()
		return
	}
	ctx.Set("loginUser", loginUser)
	ctx.Next()
}
