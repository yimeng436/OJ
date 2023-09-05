package controller

import (
	"encoding/json"
	"gatewaysvr/log"
	"gatewaysvr/rpcservice"
	"github.com/gin-gonic/gin"
	"github.com/yimeng436/OJ/common"
	"github.com/yimeng436/OJ/common/constant"
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

// @Summary		分页查询问题
// @Description	分页查询问题
// @Accept json
// @Param user body pb.GetQuestionVoPageRequest true "QuestionVoPage"
// @Produce  json
// @Success		200	{object}	common.Response
// @Router			/question/list  [post]
func ListQuestion(ctx *gin.Context) {
	request := new(pb.GetQuestionVoPageRequest)

	if err := ctx.ShouldBind(&request); err != nil {
		log.Fatal("参数错误")
		common.Fail(ctx, "请求参数错误")
		return
	}

	if request.Page.Page <= 0 {
		request.Page.Page = 0
	}
	if request.Page.PageSize <= 0 {
		request.Page.Page = 5
	}
	if request.Question == nil {
		request.Question = new(pb.QuestionInfo)
	}
	questionClient := rpcservice.GetQuestionServiceClient()
	resp, err := questionClient.GetQuestionVoPage(ctx, request)
	if err != nil {
		log.Fatal("GetQuestionVoPage rpc服务器调用异常：", err.Error())
		common.Fail(ctx, err.Error())
		return
	}
	common.Success(ctx, resp, "success")
}

// @Summary		查询问题提交信息
// @Description	查询问题提交信息
// @Accept json
// @Param user body pb.QuestionSubmitQueryRequest true "QuestionSubmit"
// @Produce  json
// @Success		200	{object}	common.Response
// @Router			/question/submit/query [post]
func QueryQuestionSubmit(ctx *gin.Context) {
	request := new(pb.QuestionSubmitQueryRequest)

	if err := ctx.ShouldBind(&request); err != nil {
		log.Fatal("参数错误")
		common.Fail(ctx, "请求参数错误")
		return
	}
	userState, exists := ctx.Get("loginUser")
	if !exists {
		common.Fail(ctx, "未登录")
		return
	}
	vo := userState.(pb.UserVo)
	serisedVo, err := json.Marshal(vo)
	if err != nil {
		log.Fatal("序列化异常")
		common.Fail(ctx, err.Error())
		return
	}
	if request.Ctx == nil {
		c := new(pb.Context)
		c.Context = make(map[string]string)
		request.Ctx = c
	}
	request.Ctx.Context[constant.UserLoginState] = string(serisedVo)
	questionSubmitClient := rpcservice.GetQuestionSubmitServiceClient()
	resp, err := questionSubmitClient.ListQuestionSubmitByPage(ctx, request)
	if err != nil {
		log.Fatal("ListQuestionSubmitByPage rpc服务器调用异常：", err.Error())
		common.Fail(ctx, err.Error())
		return
	}
	common.Success(ctx, resp, "success")
}

// @Summary		提交问题
// @Description	提交问题
// @Accept json
// @Param user body pb.QuestionSubmitAddRequest true "QuestionSubmitAddRequest"
// @Produce  json
// @Success		200	{object}	common.Response
// @Router			/question/submit/do [post]
func DoSubmit(ctx *gin.Context) {
	request := new(pb.QuestionSubmitAddRequest)

	if err := ctx.ShouldBind(&request); err != nil {
		log.Fatal("参数错误")
		common.Fail(ctx, "请求参数错误")
		return
	}
	loginUser, exists := ctx.Get("loginUser")
	if !exists {
		common.Fail(ctx, "未登录")
		return
	}
	vo := loginUser.(pb.UserVo)
	serisedVo, err := json.Marshal(vo)
	if err != nil {
		log.Fatal("序列化异常")
		common.Fail(ctx, err.Error())
		return
	}
	if request.Ctx == nil {
		c := new(pb.Context)
		c.Context = make(map[string]string)
		request.Ctx = c
	}
	request.Ctx.Context[constant.UserLoginState] = string(serisedVo)
	questionSubmitClient := rpcservice.GetQuestionSubmitServiceClient()
	resp, err := questionSubmitClient.DoQuestionSubmit(ctx, request)
	if err != nil {
		log.Fatal("DoQuestionSubmit rpc服务器调用异常：", err.Error())
		common.Fail(ctx, err.Error())
		return
	}
	common.Success(ctx, resp, "success")
}
