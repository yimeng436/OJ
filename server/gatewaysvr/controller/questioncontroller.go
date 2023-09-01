package controller

import (
	"gatewaysvr/log"
	"gatewaysvr/rpcservice"
	"github.com/gin-gonic/gin"
	"github.com/yimeng436/OJ/common"
	"github.com/yimeng436/OJ/pkg/pb"
	"strconv"
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
	questionClient := rpcservice.GetQuestionServiceClient()

	res, err := questionClient.AddQuestion(ctx, addRequest)
	if err != nil {
		log.Fatal("AddQuestion rpc服务器调用异常：", err.Error())
		common.Fail(ctx, err.Error())
		return
	}
	if res == nil || !res.Res {
		log.Fatal("添加失败", err.Error())
		common.Fail(ctx, err.Error())
		return
	}
	common.Success(ctx, res, "success")
}

// @Summary		根据id获取题目
// @Description	根据id获取题目
// @Param			id	path		int		true	"Question ID"
// @Produce  json
// @Success		200	{object}	common.Response
// @Router			/question/get/{id} [get]
func GetQuestion(ctx *gin.Context) {
	request := new(pb.QuestionIdRequest)
	param := ctx.Param("id")
	if param == "" {
		common.Fail(ctx, "id不能为空")
		return
	}
	var err error
	request.Id, err = strconv.ParseInt(param, 10, 64)
	if err != nil {
		log.Fatal("id 错误")
		common.Fail(ctx, err.Error())
		return
	}

	questionServiceClient := rpcservice.GetQuestionServiceClient()
	question, err := questionServiceClient.GetQuestionById(ctx, request)
	if err != nil {
		log.Fatal("GetQuestionById rpc服务器调用异常：", err.Error())
		common.Fail(ctx, err.Error())
		return
	}
	common.Success(ctx, question, "success")
}
