package strategy

import (
	"encoding/json"
	"errors"
	"github.com/yimeng436/OJ/common/enum"
	"github.com/yimeng436/OJ/pkg/pb"
)

var (
	Java_Program_Time int64 = 10000
)

type JavaStrategy struct {
}

func (JavaStrategy) ExecuteJudge(ctx *JudgeContext) (*pb.JudgeInfo, error) {
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
		errCase := new(pb.JudgeCase)
		errCase.Inputs = inputList[len(inputList)-1]
		errCase.Outputs = judgeCaseList[len(inputList)-1].Outputs
		judgeInfoResp.ErrCase = errCase
		return judgeInfoResp, nil
	}
	//校验结果是否正确
	for i, jude := range judgeCaseList {
		if jude.Outputs != outputList[i] {
			judgeStatus = enum.GetJudegeInfo(enum.WrongAnswer)
			judgeInfoResp.Message = judgeStatus
			errCase := new(pb.JudgeCase)
			errCase.Inputs = jude.Inputs
			errCase.Outputs = jude.Outputs
			judgeInfoResp.ErrCase = errCase
			return judgeInfoResp, nil
		}
	}
	var err error
	time := judgeInfo.Time
	memory := judgeInfo.Memory
	judgeConfigStr := question.JudgeConfig
	judgeConfig := new(pb.JudgeConfig)
	err = json.Unmarshal([]byte(judgeConfigStr), judgeConfig)
	if err != nil {
		return nil, errors.New("判题配置反序列化异常:" + err.Error())
	}
	timeLimit := judgeConfig.TimeLimit
	memoryLimit := judgeConfig.MemoryLimit
	if time-Java_Program_Time > timeLimit {
		judgeStatus = enum.GetJudegeInfo(enum.TimeLimitExceeded)
		judgeInfoResp.Message = judgeStatus
		return judgeInfoResp, nil
	}

	if (memory / (1024 * 1024)) > float32(memoryLimit) {
		judgeStatus = enum.GetJudegeInfo(enum.MemoryLimitExceeded)
		judgeInfoResp.Message = judgeStatus
		return judgeInfoResp, nil
	}

	judgeStatus = enum.GetJudegeInfo(enum.Accept)
	judgeInfoResp.Message = judgeStatus
	judgeInfoResp.Time = time
	judgeInfoResp.Memory = memory / (1024 * 1024)
	return judgeInfoResp, nil

}
