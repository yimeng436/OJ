package service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/jinzhu/copier"
	"github.com/yimeng436/OJ/common/constant"
	"github.com/yimeng436/OJ/pkg/pb"
	"google.golang.org/protobuf/encoding/protojson"
	"questionsvr/repository"
	"questionsvr/utils"
)

var questionService QuestionService

type QuestionService struct {
	pb.UnimplementedQuestionServiceServer
}

func (QuestionService) ValidQuestion(ctx context.Context, request *pb.ValidQuestionRequest) (*pb.Empty, error) {
	//用于判断是否为添加操作
	add := request.Add
	questionInfo := request.QuestionInfo

	if add {
		if utils.IsAnyBlank(questionInfo.Title, questionInfo.Content) {
			return &pb.Empty{}, errors.New("题目、内容、标签字段不能为空")
		}
	}

	if !utils.IsAnyBlank(questionInfo.Title) && len(questionInfo.Title) > 80 {
		return &pb.Empty{}, errors.New("题目过长")
	}
	if !utils.IsAnyBlank(questionInfo.Content) && len(questionInfo.Content) > 8192 {
		return &pb.Empty{}, errors.New("内容过长")
	}

	if !utils.IsAnyBlank(questionInfo.Answer) && len(questionInfo.Answer) > 8192 {
		return &pb.Empty{}, errors.New("答案过长")
	}

	if !utils.IsAnyBlank(questionInfo.JudgeCase) && len(questionInfo.JudgeCase) > 8192 {
		return &pb.Empty{}, errors.New("测试用例过长")
	}
	if !utils.IsAnyBlank(questionInfo.JudgeConfig) && len(questionInfo.JudgeConfig) > 80 {
		return &pb.Empty{}, errors.New("配置过长")
	}

	return &pb.Empty{}, nil
}
func (QuestionService) GetQuestionVoPage(ctx context.Context, request *pb.GetQuestionVoPageRequest) (*pb.GetQuestionPageVoResponse, error) {
	page := request.Page
	questionInfo := request.Question

	question := new(repository.Question)
	if questionInfo != nil {
		copier.Copy(question, questionInfo)
	}
	questionList, err := repository.GetQuestionList(question, int(page.Page), int(page.PageSize))
	if err != nil {
		return nil, err
	}
	resp := new(pb.GetQuestionPageVoResponse)
	for _, q := range questionList {
		resQuestion := new(pb.QuestionInfo)
		copier.Copy(resQuestion, q)
		resp.Question = append(resp.Question, resQuestion)
	}

	return resp, nil
}
func (QuestionService) AddQuestion(ctx context.Context, request *pb.QuestionAddRequest) (*pb.BoolResponse, error) {
	c := request.Content
	title := request.Title
	answer := request.Answer
	if utils.IsAnyBlank(c, answer, title) {
		return &pb.BoolResponse{Res: false}, errors.New("context、title、answer 不能为空")
	}
	tags := request.Tags
	judgeCase := request.JudgeCase
	judgeconfig := request.JudgeConfig
	if tags == nil {
		return &pb.BoolResponse{Res: false}, errors.New("tags 不能为空")
	}
	if judgeCase == nil {
		return &pb.BoolResponse{Res: false}, errors.New("judgeCase 不能为空")
	}

	if judgeconfig == nil {
		return &pb.BoolResponse{Res: false}, errors.New("judgeconfig 不能为空")
	}
	question := new(pb.QuestionInfo)
	copier.Copy(question, request)
	questionTags, err := json.Marshal(tags)
	if err != nil {
		return &pb.BoolResponse{Res: false}, errors.New("tags 参数异常")
	}
	questionjudgeCase, err := json.Marshal(judgeCase)
	if err != nil {
		return &pb.BoolResponse{Res: false}, errors.New("tags 参数异常")
	}
	questionjudgeconfig, err := json.Marshal(judgeconfig)
	if err != nil {
		return &pb.BoolResponse{Res: false}, errors.New("tags 参数异常")
	}

	question.Tags = string(questionTags)
	question.JudgeCase = string(questionjudgeCase)
	question.JudgeConfig = string(questionjudgeconfig)
	q := new(repository.Question)
	copier.Copy(q, question)
	_, err = repository.CreateQuestion(q)
	if err != nil {
		return &pb.BoolResponse{Res: false}, errors.New("添加失败")
	}
	return &pb.BoolResponse{Res: true}, nil

}
func (QuestionService) GetQuestionVoById(ctx context.Context, request *pb.QuestionIdRequest) (*pb.QuestionVo, error) {
	if request.Id == 0 {
		return nil, errors.New("id 不能为空")
	}
	questionInfo, err := repository.GetQuestionById(request.Id)
	if err != nil {
		return nil, err
	}
	quetionTem := new(pb.QuestionInfo)
	copier.Copy(quetionTem, questionInfo)
	questionVo, err := questionService.GetQuestionVo(ctx, quetionTem)
	questionVo.CreateTime = questionInfo.CreateTime.String()
	questionVo.UpdateTime = questionInfo.UpdateTime.String()
	err = json.Unmarshal([]byte(questionInfo.Tags), &questionVo.Tags)
	err = json.Unmarshal([]byte(questionInfo.JudgeConfig), questionVo.JudgeConfig)
	if err != nil {
		return nil, err
	}
	return questionVo, nil
}
func (QuestionService) DeleteQuestionById(ctx context.Context, request *pb.QuestionIdRequest) (*pb.BoolResponse, error) {
	if request.Id == 0 {
		return nil, errors.New("id 不能为空")
	}
	if err := repository.DeleteQuestionById(request.Id); err != nil {
		return &pb.BoolResponse{Res: false}, errors.New("删除失败")
	}
	return &pb.BoolResponse{Res: true}, nil
}
func (QuestionService) UpdateQuestion(ctx context.Context, request *pb.QuestionInfo) (*pb.BoolResponse, error) {
	question := new(repository.Question)
	copier.Copy(question, request)
	err := repository.UpdateQuestionById(question)
	if err != nil {
		return nil, err
	}
	return &pb.BoolResponse{Res: true}, nil
}
func (QuestionService) GetQuestionVo(ctx context.Context, request *pb.QuestionInfo) (*pb.QuestionVo, error) {
	questionVo := new(pb.QuestionVo)
	var err error
	err = copier.Copy(questionVo, request)
	if err != nil {
		return nil, err
	}
	if request.JudgeConfig != "" {
		protojson.Unmarshal([]byte(request.JudgeConfig), questionVo.JudgeConfig)
	}
	if request.Tags != "" {
		json.Unmarshal([]byte(request.Tags), &questionVo.Tags)
	}
	return questionVo, nil
}
func (QuestionService) GetQuestionById(ctx context.Context, request *pb.QuestionIdRequest) (*pb.QuestionInfo, error) {
	userState := request.Ctx.Context[constant.UserLoginState]
	if userState == "" {
		return nil, errors.New("请先登录")
	}
	loginUser := new(pb.UserVo)
	json.Unmarshal([]byte(userState), loginUser)
	questionInfo, err := repository.GetQuestionById(request.Id)
	if err != nil {
		return nil, err
	}
	if questionInfo.UserId != loginUser.Id && loginUser.Id != 1 {
		return nil, errors.New("没有权限")
	}
	resp := new(pb.QuestionInfo)
	copier.Copy(resp, questionInfo)
	return resp, nil

}
