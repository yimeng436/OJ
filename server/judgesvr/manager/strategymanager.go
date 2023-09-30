package manager

import (
	"github.com/yimeng436/OJ/pkg/pb"
	"judgesvr/strategy"
)

type StrategyManager struct {
}

func (StrategyManager) ExecuteJudge(ctx *strategy.JudgeContext) (*pb.JudgeInfo, error) {
	questionSubmit := ctx.QuestionSubmit
	language := questionSubmit.Language
	var useStrategy strategy.JudgeStrategy
	useStrategy = &strategy.DefaultStrategy{}
	if language == "java" {
		useStrategy = &strategy.JavaStrategy{}
	}
	return useStrategy.ExecuteJudge(ctx)
}
