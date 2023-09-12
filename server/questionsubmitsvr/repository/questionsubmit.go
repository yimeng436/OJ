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

func ListQuestionSubmitByPage(request *pb.QuestionSubmitQueryRequest) ([]*QuestionSubmit, error) {
	page := int(request.Page.Page)
	pageSize := int(request.Page.PageSize)
	db := db.GetDB()
	query := buildCondition(db, request)
	var questionSubmit []*QuestionSubmit
	err := query.Limit(pageSize).Offset(page).Find(&questionSubmit).Error
	if err != nil {
		return nil, err
	}
	return questionSubmit, nil
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
	var questionSubmit *QuestionSubmit
	err := db.Model(QuestionSubmit{}).Where("id = ?", id).Find(questionSubmit).Error
	if err != nil {
		return nil, err
	}
	return questionSubmit, nil
}

func buildCondition(db *gorm.DB, request *pb.QuestionSubmitQueryRequest) *gorm.DB {
	query := db.Model(QuestionSubmit{})
	language := request.Language
	questionId := request.QuestionId
	userId := request.UserId
	status := request.Status
	if request.Language != "" {
		query = query.Where("language = ?", language)
	}
	if request.QuestionId != 0 {
		query = query.Where("questionId = ?", questionId)
	}
	if request.UserId != 0 {
		query = query.Where("userId = ?", userId)
	}
	query = query.Where("status = ?", status)
	return query
}
