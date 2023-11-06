package service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/jinzhu/copier"
	"github.com/yimeng436/OJ/common/constant"
	"github.com/yimeng436/OJ/common/enum"
	"github.com/yimeng436/OJ/pkg/pb"
	"google.golang.org/protobuf/encoding/protojson"
	mq "questionsubmitsvr/middleware/rabbitmq"
	"questionsubmitsvr/middleware/rediscli"
	"questionsubmitsvr/remotemodel"
	"questionsubmitsvr/repository"
	"questionsubmitsvr/rpcservice"
	"strconv"
	// 必须要导入这个包，否则grpc会报错
	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
	redis "github.com/redis/go-redis/v9"
)

type QuestionSubmitService struct {
	pb.UnimplementedQuestionSubmitServiceServer
}

func (QuestionSubmitService) GetQuestionSubmitById(ctx context.Context, request *pb.IdReQuest) (*pb.QuestionSubmitInfo, error) {
	questionSubmitId := request.Id
	questionSubmit, err := repository.GetById(questionSubmitId)
	if err != nil {
		return nil, err
	}
	p := new(pb.QuestionSubmitInfo)
	err = copier.Copy(p, questionSubmit)
	if err != nil {
		return nil, err
	}
	return p, nil
}
func (QuestionSubmitService) DoQuestionSubmit(ctx context.Context, request *pb.QuestionSubmitAddRequest) (*pb.IdResponse, error) {
	language := request.Language
	_, err2 := enum.GetLanguage(language)
	if err2 != nil {
		return nil, err2
	}

	var err error
	userState := request.Ctx.Context[constant.UserLoginState]
	if userState == "" {
		return nil, errors.New("请先登录")
	}
	loginUser := new(pb.UserVo)
	json.Unmarshal([]byte(userState), loginUser)
	questionSvrClient := rpcservice.GetQuestionSvrClient()
	questionIdRequest := &pb.QuestionIdRequest{
		Id:  request.QuestionId,
		Ctx: request.Ctx,
	}

	redisCli := rediscli.GetRedisCli()
	questionStr, err := redisCli.Get(context.Background(), constant.QuestionKey).Result()
	switch err {
	case redis.Nil:

	}
	question, err := questionSvrClient.GetQuestionById(context.Background(), questionIdRequest)
	if err != nil {
		return nil, err
	}
	if question == nil {
		return nil, errors.New("题目不存在")
	}
	questionSubmit := new(repository.QuestionSubmit)
	questionSubmit.QuestionId = request.QuestionId
	questionSubmit.Code = request.Code
	questionSubmit.Language = request.Language
	questionSubmit.UserId = loginUser.Id
	questionSubmit.Status = enum.SubmitWaiting
	questionSubmit.JudgeInfo = "{}"

	err = repository.Create(questionSubmit)
	if err != nil {
		return nil, err
	}
	questionSubmitId := questionSubmit.Id

	mqProducer := mq.GetMQ("questionsubmit_exchange", "submit.question")
	mqProducer.PublishRouting(strconv.FormatInt(questionSubmitId, 10))

	return &pb.IdResponse{Id: questionSubmitId}, nil
}

func (QuestionSubmitService) ListQuestionSubmitByPage(ctx context.Context, request *pb.QuestionSubmitQueryRequest) (*pb.QuestionSubmitQueryResponse, error) {
	userState := request.Ctx.Context[constant.UserLoginState]
	if userState == "" {
		return nil, errors.New("请先登录")
	}
	questionSubmitList, total, err := repository.ListQuestionSubmitByPage(request)
	if err != nil {
		return nil, err
	}
	loginUser := new(pb.UserVo)
	json.Unmarshal([]byte(userState), loginUser)

	voList := make([]*pb.QuestionSubmitVo, 0)
	for _, v := range questionSubmitList {
		vo := toVo(v)
		if vo.UserId != loginUser.Id && loginUser.UserRole != 1 {
			vo.Code = ""
		}
		voList = append(voList, vo)
	}
	resp := new(pb.QuestionSubmitQueryResponse)

	resp.QuestionSubmitVO = voList
	resp.Total = total
	return resp, nil
}
func (QuestionSubmitService) GetQuestionSubmitTotal(context.Context, *pb.Empty) (*pb.TotalResponse, error) {
	total, err := repository.GetTotal()
	if err != nil {
		return nil, err
	}
	return &pb.TotalResponse{Total: total}, nil
}
func (QuestionSubmitService) UpdateQuestionStatusById(ctx context.Context, request *pb.QuestionSubmitInfo) (*pb.BoolResponse, error) {
	var questionsubmit = new(repository.QuestionSubmit)

	copier.Copy(questionsubmit, request)
	err := repository.UpdateQuestionSubmit(questionsubmit)
	if err != nil {
		return nil, err
	}
	return &pb.BoolResponse{Res: true}, nil
}
func toVo(obj *repository.QuestionSubmit) *pb.QuestionSubmitVo {
	vo := new(pb.QuestionSubmitVo)
	copier.Copy(vo, obj)
	vo.CreateTime = obj.CreateTime.Format("2006-01-02 15:04:05")
	vo.UpdateTime = obj.CreateTime.Format("2006-01-02 15:04:05")
	vo.SubmitStatus = enum.GetValue(int(vo.Status))
	if obj.JudgeInfo != "" {
		protojson.Unmarshal([]byte(obj.JudgeInfo), vo.JudgeInfo)
	}
	if !obj.Question.IsEmpty() {
		questionVo, err := toQuestionVo(obj.Question)
		if err != nil {
			return nil
		}
		vo.QuestionVo = questionVo
	}

	if !obj.User.IsEmpty() {
		userVo := new(pb.UserVo)
		err := copier.Copy(userVo, obj.User)
		if err != nil {
			return nil
		}
		vo.SubmitUser = userVo
	}
	return vo
}

func toQuestionVo(question remotemodel.Question) (*pb.QuestionVo, error) {
	questionVo := new(pb.QuestionVo)
	err := copier.Copy(questionVo, question)
	if err != nil {
		return nil, err
	}
	if question.Tags != "" {
		json.Unmarshal([]byte(question.Tags), &questionVo.Tags)
	}
	if question.JudgeConfig != "" {
		json.Unmarshal([]byte(question.JudgeConfig), &questionVo.JudgeConfig)
	}
	return questionVo, nil
}
