package service

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"github.com/yimeng436/OJ/pkg/pb"
	"google.golang.org/protobuf/encoding/protojson"
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

func (QuestionSubmitService) ListQuestionSubmitByPage(ctx context.Context, request *pb.QuestionSubmitQueryRequest) (*pb.QuestionSubmitQueryResponse, error) {
	questionSubmitList, err := repository.ListQuestionSubmitByPage(request)
	if err != nil {
		return nil, err
	}
	userClient := rpcservice.GetUserServiceClient()
	loginUser, err := userClient.GetLoginUser(ctx, &pb.Empty{})
	if err != nil {
		return nil, err
	}
	voList := make([]*pb.QuestionSubmitVo, 0)
	for _, v := range questionSubmitList {
		vo := toVo(v)
		if vo.UserId != loginUser.Id && loginUser.UserRole != 0 {
			vo.Code = ""
		}
		voList = append(voList, vo)
	}
	resp := new(pb.GetQuestionPageVoResponse)
	resp.QuestionVo = voList
	return voList, nil
}

func toVo(obj *repository.QuestionSubmit) *pb.QuestionSubmitVo {
	var vo *pb.QuestionSubmitVo
	copier.Copy(vo, obj)
	if obj.JudgeInfo != "" {
		protojson.Unmarshal([]byte(obj.JudgeInfo), vo.JudgeInfo)
	}
	return vo
}
