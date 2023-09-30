package service

import (
	"codesandsvr/log"
	"context"
	"github.com/yimeng436/OJ/pkg/pb"
)

type CodeSandProxy struct {
	pb.UnimplementedCodeSandServiceServer
	CodeSandService pb.CodeSandServiceServer
}

func (p CodeSandProxy) ExecuteCode(ctx context.Context, request *pb.ExecuteCodeRequest) (*pb.ExecuteCodeResponse, error) {
	log.Infof("before")
	resp, err := p.CodeSandService.ExecuteCode(ctx, request)
	log.Infof("after")
	return resp, err
}
