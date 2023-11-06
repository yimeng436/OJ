package service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/yimeng436/OJ/common/enum"
	"github.com/yimeng436/OJ/pkg/pb"
	"judgesvr/constant"
	"judgesvr/manager"
	"judgesvr/middleware/db"
	"judgesvr/rpcservice"
	"judgesvr/strategy"
	// 必须要导入这个包，否则grpc会报错
	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
)

type JudgeService struct {
	pb.UnimplementedJudgeServiceServer
}

func (JudgeService) DoJudge(ctx context.Context, request *pb.DoJudgeRequest) (*pb.QuestionSubmitInfo, error) {
	questionsubmitid := request.Questionsubmitid
	return Judge(context.Background(), questionsubmitid)
}

func Judge(ctx context.Context, questionsubmitid int64) (*pb.QuestionSubmitInfo, error) {
	questionSubmitClient := rpcservice.GetQuestionSubmitClient()
	quesionsubmit, err := questionSubmitClient.GetQuestionSubmitById(context.Background(), &pb.IdReQuest{Id: questionsubmitid})
	if err != nil {
		return nil, err
	}
	questionId := quesionsubmit.QuestionId
	questionClient := rpcservice.GetQuestionSvrClient()
	question, err := questionClient.GetById(context.Background(), &pb.QuestionIdRequest{Id: questionId})
	if err != nil {
		return nil, err
	}
	if quesionsubmit.Status == enum.SubmitRunning {
		return nil, errors.New("判题中，请等待")
	}
	quesionsubmit.Status = enum.SubmitRunning
	// TODO 调用questionSubmitClient的跟新状态函数
	_, err = questionSubmitClient.UpdateQuestionStatusById(context.Background(), quesionsubmit)
	if err != nil {
		return nil, err
	}

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
	inputlist := make([]string, 0)
	for _, judgeCase := range judgeCaseObj {

		inputlist = append(inputlist, judgeCase.Inputs)
	}
	excudeResp, err := codeSandSvrClient.ExecuteCode(ctx,
		&pb.ExecuteCodeRequest{
			Code:     quesionsubmit.Code,
			Language: quesionsubmit.Language,
			Inputs:   inputlist,
		})

	// 用户代码执行结果
	outputs := excudeResp.Outputs
	if outputs != nil && len(outputs) == len(inputlist) {
		judgeContent := &strategy.JudgeContext{
			OutputList:     outputs,
			InputList:      inputlist,
			JudgeCaseList:  judgeCaseObj,
			Question:       question,
			JudgeInfo:      excudeResp.JudgeInfo,
			QuestionSubmit: quesionsubmit,
		}
		strategyManager := manager.StrategyManager{}
		// 拿到判题信息
		judgeInfo, err := strategyManager.ExecuteJudge(judgeContent)
		if err != nil {
			return nil, err
		}

		// 跟新这次提交记录的判题信息和状态
		judgeInfoStr, err := json.Marshal(judgeInfo)
		if err != nil {
			return nil, errors.New("序列化异常:" + err.Error())
		}
		quesionsubmit.JudgeInfo = string(judgeInfoStr)
		quesionsubmit.Status = enum.SubmitSucceed
		quesionsubmit.JudgeStatus = judgeInfo.JudgeStatus
	} else {
		judgeInfo := pb.JudgeInfo{Memory: 0, Time: 0, Message: excudeResp.Message, JudgeStatus: excudeResp.JudgeInfo.JudgeStatus}
		judgeInfoStr, err := json.Marshal(judgeInfo)
		if err != nil {
			return nil, errors.New("序列化异常:" + err.Error())
		}
		quesionsubmit.JudgeInfo = string(judgeInfoStr)
		quesionsubmit.Status = 2
		quesionsubmit.JudgeStatus = constant.Failed
	}
	Db := db.GetDB()
	tx := Db.Begin()

	_, err = questionSubmitClient.UpdateQuestionStatusById(context.Background(), quesionsubmit)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	if quesionsubmit.JudgeStatus == constant.Success {
		question.AcceptedNum++
		resp, err := questionClient.UpdateQuestion(context.Background(), question)
		if err != nil || !resp.Res {
			tx.Rollback()
			return nil, err
		}
	}

	return quesionsubmit, nil
}
