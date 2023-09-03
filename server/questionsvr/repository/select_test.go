package repository

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/copier"
	"testing"
)

type QuestionRequest struct {
	Id      int64    `gorm:"column:id;primary_key;AUTO_INCREMENT;comment:'id'" json:"id,omitempty"`
	Title   string   `gorm:"column:title;comment:'标题'" json:"title,omitempty"`
	Content string   `gorm:"column:content;comment:'内容'" json:"content,omitempty"`
	Tags    []string `gorm:"column:tags;comment:'标签列表（json 数组）'" json:"tags,omitempty"`
	Answer  string   `gorm:"column:answer;comment:'题目答案'" json:"answer,omitempty"`
}

func TestMethod(t *testing.T) {
	question := &Question{
		Id:      1,
		Title:   "A",
		Content: "asd",
		Tags:    "[\"数组\",\"简单\"]",
		Answer:  "1,2,3",
	}
	request := new(QuestionRequest)
	copier.Copy(request, question)
	err := json.Unmarshal([]byte(question.Tags), &request.Tags)
	fmt.Println(err)
	fmt.Println(request)
}
