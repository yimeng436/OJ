package service

import (
	"context"
	"github.com/yimeng436/OJ/pkg/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type JudgeService struct {
	pb.UnimplementedJudgeServiceServer
}

func (JudgeService) DoJudge(ctx context.Context, request *pb.DoJudgeRequest) (*pb.QuestionSubmitInfo, error) {

	return nil, status.Errorf(codes.Unimplemented, "method ExecuteCode not implemented")
}
