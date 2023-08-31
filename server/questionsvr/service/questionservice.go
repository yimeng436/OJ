package service

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"github.com/yimeng436/OJ/pkg/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"questionsvr/utils"
)

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
func (QuestionService) GetQuestionVoPage(ctx context.Context, request *pb.GetQuestionVoPageRequest) (*pb.GetQuestionVoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQuestionVoPage not implemented")
}
func (QuestionService) AddQuestion(ctx context.Context, request *pb.QuestionInfo) (*pb.BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddQuestion not implemented")
}
func (QuestionService) GetQuestionById(ctx context.Context, request *pb.QuestionIdRequest) (*pb.QuestionInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQuestionById not implemented")
}
func (QuestionService) DeleteQuestionById(ctx context.Context, request *pb.QuestionIdRequest) (*pb.BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteQuestionById not implemented")
}
func (QuestionService) UpdateQuestionById(ctx context.Context, request *pb.QuestionIdRequest) (*pb.BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateQuestionById not implemented")
}
func (QuestionService) GetQuestionVo(ctx context.Context, request *pb.QuestionInfo) (*pb.QuestionVo, error) {
	var questionVo *pb.QuestionVo
	var err error
	err = copier.Copy(questionVo, request)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
