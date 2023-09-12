package service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/yimeng436/OJ/common/enum"
	"github.com/yimeng436/OJ/pkg/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"judgesvr/rpcservice"
	"judgesvr/strategy"
)

type JudgeService struct {
	pb.UnimplementedJudgeServiceServer
}

func (JudgeService) DoJudge(ctx context.Context, request *pb.DoJudgeRequest) (*pb.QuestionSubmitInfo, error) {
	questionsubmitid := request.Questionsubmitid
	questionSubmitClient := rpcservice.GetQuestionSubmitClient()
	quesionsubmit, err := questionSubmitClient.GetQuestionSubmitById(ctx, &pb.IdReQuest{Id: questionsubmitid})
	if err != nil {
		return nil, err
	}
	questionId := quesionsubmit.QuestionId
	questionClient := rpcservice.GetQuestionSvrClient()
	question, err := questionClient.GetById(ctx, &pb.QuestionIdRequest{Id: questionId})
	if err != nil {
		return nil, err
	}
	if quesionsubmit.Status == enum.Waiting {
		return nil, errors.New("判题中，请等待")
	}
	quesionsubmit.Status = enum.Running
	// TODO 调用questionSubmitClient的跟新状态函数

	// 调用沙箱服务
	codeSandSvrClient := rpcservice.GetCodeSandSvrClient()
	judgeCase := question.JudgeCase
	var judgeCaseObj []pb.JudgeCase

	// string --> JudgeCase数组
	err = json.Unmarshal([]byte(judgeCase), &judgeCaseObj)
	if err != nil {
		return nil, errors.New("测试用例反序列化异常:" + err.Error())
	}
	//取出每一个JudgeCase里面的inputs
	var inputlist []string
	for _, judgeCase := range judgeCaseObj {
		var inputs string
		err = json.Unmarshal([]byte(judgeCase.Inputs), &inputs)
		if err != nil {
			return nil, errors.New("测试用例反序列化异常:" + err.Error())
		}
		inputlist = append(inputlist, inputs)
	}
	excudeResp, err := codeSandSvrClient.ExecuteCode(ctx,
		&pb.ExecuteCodeRequest{
			Code:     quesionsubmit.Code,
			Language: quesionsubmit.Language,
			Inputs:   inputlist,
		})

	// 用户代码执行结果
	outputs := excudeResp.Outputs
	judgeStatus := enum.GetJudegeInfo(enum.Accept)
	if len(outputs) != len(inputlist) {
		judgeStatus = enum.GetJudegeInfo(enum.WrongAnswer)
		return nil, errors.New(judgeStatus)
	}
	//校验结果是否正确
	for i, jude := range judgeCaseObj {
		if jude.Outputs != outputs[i] {
			judgeStatus = enum.GetJudegeInfo(enum.WrongAnswer)
			return nil, errors.New(judgeStatus)
		}
	}

	content := strategy.JudgeContent{
		OutputList:     outputs,
		InputList:      inputlist,
		JudgeCaseList:  judgeCaseObj,
		Question:       question,
		JudgeInfo:      excudeResp.JudgeInfo,
		QuestionSubmit: quesionsubmit,
	}

	return nil, status.Errorf(codes.Unimplemented, "method ExecuteCode not implemented")
}
