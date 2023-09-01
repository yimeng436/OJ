package controller

import (
	"gatewaysvr/config"
	"gatewaysvr/log"
	"gatewaysvr/rpcservice"
	"github.com/gin-gonic/gin"
	"github.com/yimeng436/OJ/common"
	"github.com/yimeng436/OJ/pkg/pb"
)

// @Summary		添加问题
// @Description	添加问题
// @Accept json
// @Param user body pb.QuestionAddRequest true "question"
// @Produce  json
// @Success		200	{object}	common.Response
// @Router			/question/add  [post]
func AddQuestion(ctx *gin.Context) {
	var addRequest = new(pb.QuestionAddRequest)

	if err := ctx.ShouldBind(addRequest); err != nil {
		log.Fatal("参数错误")
		common.Fail(ctx, "请求参数错误")
		return
	}
	questionClient := rpcservice.NewQuestionSvrClient(config.GetGlobalConfig().SvrConfig.QuestionSvrName)
	if questionClient == nil {
		log.Fatal("QuestionSvrClient rpc服务初始化异常")
		common.Fail(ctx, "系统错误")
		return
	}

	res, err := questionClient.AddQuestion(ctx, addRequest)
	if err != nil {
		log.Fatal("AddQuestion rpc服务器调用异常：", err.Error())
		common.Fail(ctx, err.Error())
		return
	}
	if res == nil || !res.Res {
		log.Fatal()
	}
}

// @Summary		根据id获取题目
// @Description	根据id获取题目
// @Param			id	path		int		true	"Question ID"
// @Produce  json
// @Success		200	{object}	common.Response
// @Router			/question/get/{id} [get]
func GetQuestion(ctx *gin.Context) {

	common.Success(ctx)
}
