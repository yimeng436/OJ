package strategy

import "github.com/yimeng436/OJ/pkg/pb"

type JudgeContext struct {
	JudgeInfo      *pb.JudgeInfo          `json:"judgeInfo"`
	InputList      []string               `json:"inputList"`
	OutputList     []string               `json:"outputList"`
	JudgeCaseList  []pb.JudgeCase         `json:"judgeCaseList"`
	Question       *pb.QuestionInfo       `json:"question"`
	QuestionSubmit *pb.QuestionSubmitInfo `json:"questionSubmit"`
}

type JudgeStrategy interface {
	ExecuteJudge(*JudgeContext) (*pb.JudgeInfo, error)
}
