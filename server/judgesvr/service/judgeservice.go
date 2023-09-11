package service

import (
	"context"
	"github.com/yimeng436/OJ/pkg/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type JudgeService struct {
	pb.UnimplementedStreamGreeterServer
}

func (JudgeService) DoJudge() (*pb.QuestionSubmitInfo, error) {

	return nil, status.Errorf(codes.Unimplemented, "method ExecuteCode not implemented")
}
func (JudgeService) DoJudge(context.Context, *QuestionSubmitInfo) (*DoJudgeRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoJudge not implemented")
}
