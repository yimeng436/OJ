package service

import (
	"context"
	"github.com/yimeng436/OJ/pkg/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"judgesvr/rpcservice"
)

type JudgeService struct {
	pb.UnimplementedJudgeServiceServer
}

func (JudgeService) DoJudge(ctx context.Context, request *pb.DoJudgeRequest) (*pb.QuestionSubmitInfo, error) {
	questionsubmitid := request.Questionsubmitid
	questionSubmitClient := rpcservice.GetQuestionSubmitClient()
	var questionsubmit *pb.QuestionSubmitInfo
	quesionsubmit, err := questionSubmitClient.GetQuestionSubmitById(&pb.IdReQuest{Id: questionsubmitid})
	if err != nil {
		return nil, err
	}
	questionId := questionsubmit.QuestionId
	questionClient := rpcservice.GetQuestionSvrClient()
	id, err2 := questionClient.GetQuestionById(ctx, &pb.QuestionIdRequest{Id: questionId, Ctx: nil})
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteCode not implemented")
}
