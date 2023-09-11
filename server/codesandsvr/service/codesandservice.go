package service

import (
	"context"
	"github.com/yimeng436/OJ/pkg/pb"
)

type CodeSandService struct {
	pb.UnimplementedCodeSandServiceServer
}

func (CodeSandService) ExecuteCode(ctx context.Context, request *pb.ExecuteCodeRequest) (*pb.ExecuteCodeResponse, error) {
	code := request.Code
	inputs := request.Inputs
	language := request.Language

	// 简单测试一下 输入做输出
	resp := new(pb.ExecuteCodeResponse)
	resp.Outputs = inputs
	resp.Status = "success"
	resp.JudgeInfo = &pb.JudgeInfo{
		Message: code,
		Time:    0,
		Memory:  0,
	}
	resp.Message = language

	return resp, nil
}
