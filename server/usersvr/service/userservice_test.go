package service

import (
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/yimeng436/OJ/pkg/pb"
	"google.golang.org/protobuf/encoding/protojson"
	"testing"
)

type ListString struct {
	Tags []string
}

type String struct {
	Tags string
}

func TestJson(t *testing.T) {
	question := new(pb.QuestionInfo)
	question.Id = 1
	question.Title = "test"
	question.Content = "test"
	question.Tags = "[1,2,3]"
	question.Answer = "test"

	question.JudgeConfig = "JudgeConfig:{" +
		"TimeLimit:12," +
		"MemoryLimit:12," +
		"StackLimit:12" +
		"}"
	questionVo := new(pb.QuestionVo)
	copier.Copy(questionVo, question)

	question.JudgeCase = `{
	"inputs": "test",
	"outputs": "test"
}`
	judgeCase := new(pb.JudgeCase)
	err := protojson.Unmarshal([]byte(question.JudgeCase), judgeCase)
	fmt.Println(err)
	fmt.Println(questionVo)
}
