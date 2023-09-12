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
	"questionsubmitsvr/middleware/log"
	"questionsubmitsvr/repository"
	"questionsubmitsvr/rpcservice"

	// 必须要导入这个包，否则grpc会报错
	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
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
func (QuestionSubmitService) DoQuestionSubmit(ctx context.Context, request *pb.QuestionSubmitAddRequest) (*pb.BoolResponse, error) {
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
	question, err := questionSvrClient.GetQuestionById(ctx, questionIdRequest)
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
	questionSubmit.Status = enum.Waiting
	questionSubmit.JudgeInfo = "{}"

	err = repository.Create(questionSubmit)
	if err != nil {
		return nil, err
	}
	questionSubmitId := questionSubmit.Id
	judgeServiceClient := rpcservice.GetJudgeServiceClient()
	go func() {
		defer func() {
			if r := recover(); r != nil {
				// 发生 panic，进行恢复处理
				log.Fatal("Recovered from panic:", r)
			}
		}()
		_, err = judgeServiceClient.DoJudge(ctx, &pb.DoJudgeRequest{Questionsubmitid: questionSubmitId})
	}()

	return &pb.BoolResponse{Res: true}, nil
}

func (QuestionSubmitService) ListQuestionSubmitByPage(ctx context.Context, request *pb.QuestionSubmitQueryRequest) (*pb.QuestionSubmitQueryResponse, error) {
	questionSubmitList, err := repository.ListQuestionSubmitByPage(request)
	if err != nil {
		return nil, err
	}
	userState := request.Ctx.Context[constant.UserLoginState]
	if userState == "" {
		return nil, errors.New("请先登录")
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

	resp.QuestionVO = voList
	return resp, nil
}
func (QuestionSubmitService) GetQuestionSubmitTotal(context.Context, *pb.Empty) (*pb.TotalResponse, error) {
	total, err := repository.GetTotal()
	if err != nil {
		return nil, err
	}
	return &pb.TotalResponse{Total: total}, nil
}
func toVo(obj *repository.QuestionSubmit) *pb.QuestionSubmitVo {
	vo := new(pb.QuestionSubmitVo)
	copier.Copy(vo, obj)
	if obj.JudgeInfo != "" {
		protojson.Unmarshal([]byte(obj.JudgeInfo), vo.JudgeInfo)
	}
	return vo
}
