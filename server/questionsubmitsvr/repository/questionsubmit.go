package repository

import "questionsubmitsvr/middleware/db"

func Create(submit *QuestionSubmit) error {
	db := db.GetDB()
	err := db.Model(QuestionSubmit{}).Create(submit).Error
	if err != nil {
		return err
	}
	return nil
}
