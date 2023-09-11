package service

import (
	"context"
	"github.com/yimeng436/OJ/pkg/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ThirdPartCodeSandService struct {
	pb.UnimplementedCodeSandServiceServer
}

func (ThirdPartCodeSandService) ExecuteCode(ctx context.Context, request *pb.ExecuteCodeRequest) (*pb.ExecuteCodeResponse, error) {

	return nil, status.Errorf(codes.Unimplemented, "method ExecuteCode not implemented")
}
