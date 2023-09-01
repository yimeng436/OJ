package repository

import (
	"questionsvr/middleware/db"
)

func CreateQuestion(question *Question) (int64, error) {
	db := db.GetDB()
	err := db.Model(Question{}).Create(question).Error
	if err != nil {
		return 0, nil
	}
	return question.Id, nil
}

func GetQuestionById(id int64) (*Question, error) {
	db := db.GetDB()
	question := new(Question)
	err := db.Model(Question{}).Where("id = ?", id).First(question).Error
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
