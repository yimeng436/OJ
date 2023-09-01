package repository

import (
	"time"
)

type QuestionSubmit struct {
	Id         int64     `gorm:"column:id;primary_key;AUTO_INCREMENT;comment:'id'" json:"id,omitempty"`
	Language   string    `gorm:"column:language;NOT NULL;comment:'编程语言'" json:"language,omitempty"`
	Code       string    `gorm:"column:code;NOT NULL;comment:'用户代码'" json:"code,omitempty"`
	JudgeInfo  string    `gorm:"column:judgeInfo;comment:'判题信息（json 对象）'" json:"judgeInfo,omitempty"`
	Status     int32     `gorm:"column:status;default:0;NOT NULL;comment:'判题状态（0 - 待判题、1 - 判题中、2 - 成功、3 - 失败）'" json:"status,omitempty"`
	QuestionId int64     `gorm:"column:questionId;NOT NULL;comment:'题目 id'" json:"questionId,omitempty"`
	UserId     int64     `gorm:"column:userId;NOT NULL;comment:'创建用户 id'" json:"userId,omitempty"`
	CreateTime time.Time `gorm:"column:createTime;default:CURRENT_TIMESTAMP;NOT NULL;comment:'创建时间'" json:"createTime,omitempty"`
	UpdateTime time.Time `gorm:"column:updateTime;default:CURRENT_TIMESTAMP;NOT NULL;comment:'更新时间'" json:"updateTime,omitempty"`
	IsDelete   int8      `gorm:"column:isDelete;default:0;NOT NULL;comment:'是否删除'" json:"isDelete,omitempty"`
}

func (q *QuestionSubmit) TableName() string {
	return "question_submit"
}
