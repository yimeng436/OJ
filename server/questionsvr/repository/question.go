package repository

import (
	"github.com/yimeng436/OJ/pkg/pb"
	"questionsvr/middleware/db"
)

func CreateQuestion(question *pb.QuestionInfo) (int64, error) {
	db := db.GetDB()
	err := db.Model(Question{}).Create(question).Error
	if err != nil {
		return 0, nil
	}
	return question.Id, nil
}

func GetQuestionById(id int64) (*pb.QuestionInfo, error) {
	db := db.GetDB()
	question := new(pb.QuestionInfo)
	err := db.Model(pb.QuestionInfo{}).Where("id = ?", id).First(question).Error
	if err != nil {
		return nil, err
	}
	return question, nil
}

func DeleteQuestionById(id int64) error {
	db := db.GetDB()
	err := db.Delete(Question{}, id).Error
	return err
}
