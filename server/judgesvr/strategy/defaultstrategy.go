package strategy

import (
	"encoding/json"
	"errors"
	"github.com/yimeng436/OJ/common/enum"
	"github.com/yimeng436/OJ/pkg/pb"
)

type DefaultStrategy struct {
}

func (DefaultStrategy) ExecuteJudge(ctx *JudgeContent) (*pb.JudgeInfo, error) {
	judgeCaseList := ctx.JudgeCaseList
	question := ctx.Question
	judgeInfo := ctx.JudgeInfo
	inputList := ctx.InputList
	outputList := ctx.OutputList
	judgeStatus := enum.GetJudegeInfo(enum.Accept)
	judgeInfoResp := &pb.JudgeInfo{
		Message: judgeStatus,
		Memory:  0,
		Time:    0,
	}

	if len(outputList) != len(inputList) {
		judgeStatus = enum.GetJudegeInfo(enum.WrongAnswer)
		judgeInfoResp.Message = judgeStatus
		return judgeInfoResp, errors.New(judgeStatus)
	}
	//校验结果是否正确
	for i, jude := range judgeCaseList {
		if jude.Outputs != outputList[i] {
			judgeInfoResp.Message = judgeStatus
			return judgeInfoResp, errors.New(judgeStatus)
		}
	}
	var err error
	time := judgeInfo.Time
	memory := judgeInfo.Memory
	judgeConfigStr := question.JudgeConfig
	var judgeConfig *pb.JudgeConfig
	err = json.Unmarshal([]byte(judgeConfigStr), judgeConfig)
	if err != nil {
		return nil, errors.New("判题配置反序列化异常:" + err.Error())
	}
	timeLimit := judgeConfig.TimeLimit
	memoryLimit := judgeConfig.MemoryLimit
	if time > timeLimit {
		judgeStatus = enum.GetJudegeInfo(enum.TimeLimitExceeded)
		judgeInfoResp.Message = judgeStatus
		return judgeInfoResp, errors.New(judgeStatus)
	}

	if memory > memoryLimit {
		judgeStatus = enum.GetJudegeInfo(enum.MemoryLimitExceeded)
		judgeInfoResp.Message = judgeStatus
		return judgeInfoResp, errors.New(judgeStatus)
	}

	judgeStatus = enum.GetJudegeInfo(enum.Accept)
	judgeInfoResp.Message = judgeStatus
	return judgeInfoResp, nil

}
