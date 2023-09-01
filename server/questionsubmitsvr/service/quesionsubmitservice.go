package service

import (
	"context"
	"errors"
	"github.com/yimeng436/OJ/pkg/pb"
	"questionsubmitsvr/enum"
	"questionsubmitsvr/repository"
	"questionsubmitsvr/rpcservice"
)

type QuestionSubmitService struct {
	pb.UnimplementedQuestionSubmitServiceServer
}

func (QuestionSubmitService) DoQuestionSubmit(ctx context.Context, request *pb.QuestionSubmitAddRequest) (*pb.BoolResponse, error) {
	language := request.Language
	_, err2 := enum.GetLanguage(language)
	if err2 != nil {
		return nil, err2
	}

	userServiceClient := rpcservice.GetUserServiceClient()
	var err error
	var loginUser *pb.UserVo
	loginUser, err = userServiceClient.GetLoginUser(ctx, &pb.Empty{})
	if err != nil {
		return nil, err
	}

	if loginUser == nil {
		return nil, errors.New("请先登录")
	}

	questionSvrClient := rpcservice.GetQuestionSvrClient()
	question, err := questionSvrClient.GetQuestionById(ctx, &pb.QuestionIdRequest{Id: request.QuestionId})
	if err != nil {
		return nil, err
	}
	if question == nil {
		return nil, errors.New("题目不存在")
	}
	if loginUser == nil {
		return nil, errors.New("未登录")
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
	return &pb.BoolResponse{Res: true}, nil
}
