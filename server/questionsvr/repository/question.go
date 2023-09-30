package repository

import (
	"encoding/json"
	"gorm.io/gorm"
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
	err := db.Delete(&Question{}, id).Error
	return err
}

func UpdateQuestionById(question *Question) error {
	db := db.GetDB()
	err := db.Where("id = ?", question.Id).Updates(question).Error
	return err
}

func GetQuestionList(question *Question, page, pageSize int) ([]*Question, error) {
	db := db.GetDB()
	query := buildSearchCondition(db, question)
	var questions []*Question
	if page > 0 {
		page--
	}
	err := query.Limit(pageSize).Offset(page * pageSize).Find(&questions).Error
	if err != nil {
		return nil, err
	}
	return questions, nil
}

func buildSearchCondition(db *gorm.DB, question *Question) (tx *gorm.DB) {
	query := db.Model(Question{})
	id := question.Id
	content := question.Content
	tags := question.Tags
	title := question.Title
	userId := question.UserId
	level := question.Level

	if content != "" {
		query = query.Or("content LIKE ?", "%"+content+"%")
	}
	if title != "" {
		query = query.Or("title LIKE ?", "%"+title+"%")
	}
	if userId != 0 {
		query = query.Or("userId = ?", userId)
	}
	if id != 0 {
		query = query.Or("id = ?", id)
	}
	if tags != "" {
		var targetTag []string
		err := json.Unmarshal([]byte(tags), &targetTag)
		if err != nil {
			return
		}

		for _, tag := range targetTag {
			query = query.Or("tags LIKE ?", "%"+tag+"%")
		}
	}
	if level != "" {
		query = query.Or("level = ?", level)
	}
	return query
}

func GetTotal() (int64, error) {
	db := db.GetDB()
	var total int64
	err := db.Model(Question{}).Count(&total).Error
	if err != nil {
		return 0, err
	}
	return total, nil
}
