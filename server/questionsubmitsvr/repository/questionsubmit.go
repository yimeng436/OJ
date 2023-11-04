package repository

import (
	"github.com/yimeng436/OJ/pkg/pb"
	"gorm.io/gorm"
	"questionsubmitsvr/middleware/db"
)

func Create(submit *QuestionSubmit) error {
	db := db.GetDB()
	err := db.Model(QuestionSubmit{}).Create(submit).Error
	if err != nil {
		return err
	}
	return nil
}

func ListQuestionSubmitByPage(request *pb.QuestionSubmitQueryRequest) ([]*QuestionSubmit, int64, error) {
	page := int(request.Page.Page)
	pageSize := int(request.Page.PageSize)
	field := request.Page.SortField
	db := db.GetDB()
	query := buildCondition(db, request)
	var total int64
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	if err != nil {
		return nil, 0, err
	}
	var questionSubmit []*QuestionSubmit
	query = query.Preload("Question").Preload("User")
	if field != "" {
		query = query.Order(field + " desc")
	} else {
		query = query.Order("createTime desc")
	}
	err = query.Limit(pageSize).Offset((page - 1) * pageSize).Find(&questionSubmit).Error
	if err != nil {
		return nil, 0, err
	}
	return questionSubmit, total, nil
}
func GetTotal() (int64, error) {
	db := db.GetDB()
	var total int64
	err := db.Model(QuestionSubmit{}).Count(&total).Error
	if err != nil {
		return 0, err
	}
	return total, nil
}

func GetById(id int64) (*QuestionSubmit, error) {
	db := db.GetDB()
	var questionSubmit = new(QuestionSubmit)
	err := db.Model(QuestionSubmit{}).Where("id = ?", id).Find(questionSubmit).Error
	if err != nil {
		return nil, err
	}
	return questionSubmit, nil
}

func UpdateQuestionSubmit(questionsubmit *QuestionSubmit) error {
	db := db.GetDB()
	err := db.Model(QuestionSubmit{}).Where("id = ?", questionsubmit.Id).Updates(questionsubmit).Error
	return err
}

func buildCondition(db *gorm.DB, request *pb.QuestionSubmitQueryRequest) *gorm.DB {
	query := db.Model(QuestionSubmit{})
	language := request.Language
	questionId := request.QuestionId
	userId := request.UserId
	status := request.Status
	if language != "" {
		query = query.Where("language = ?", language)
	}
	if questionId != 0 {
		query = query.Where("questionId = ?", questionId)
	}
	if userId != 0 {
		query = query.Where("userId = ?", userId)
	}
	if status != 0 {
		query = query.Where("status = ?", status)
	}
	return query
}
