package repository

import (
	"time"
)

type Question struct {
	Id          int64     `gorm:"column:id;primary_key;AUTO_INCREMENT;comment:'id'" json:"id,omitempty"`
	Title       string    `gorm:"column:title;comment:'标题'" json:"title,omitempty"`
	Content     string    `gorm:"column:content;comment:'内容'" json:"content,omitempty"`
	Tags        string    `gorm:"column:tags;comment:'标签列表（json 数组）'" json:"tags,omitempty"`
	Answer      string    `gorm:"column:answer;comment:'题目答案'" json:"answer,omitempty"`
	SubmitNum   int32     `gorm:"column:submitNum;default:0;NOT NULL;comment:'题目提交数'" json:"submitNum,omitempty"`
	AcceptedNum int32     `gorm:"column:acceptedNum;default:0;NOT NULL;comment:'题目通过数'" json:"acceptedNum,omitempty"`
	JudgeCase   string    `gorm:"column:judgeCase;comment:'判题用例（json 数组）'" json:"judgeCase,omitempty"`
	JudgeConfig string    `gorm:"column:judgeConfig;comment:'判题配置（json 对象）'" json:"judgeConfig,omitempty"`
	ThumbNum    int32     `gorm:"column:thumbNum;default:0;NOT NULL;comment:'点赞数'" json:"thumbNum,omitempty"`
	FavourNum   int32     `gorm:"column:favourNum;default:0;NOT NULL;comment:'收藏数'" json:"favourNum,omitempty"`
	UserId      int64     `gorm:"column:userId;NOT NULL;comment:'创建用户 id'" json:"userId,omitempty"`
	CreateTime  time.Time `gorm:"column:createTime;default:CURRENT_TIMESTAMP;NOT NULL;comment:'创建时间'" json:"createTime,omitempty"`
	UpdateTime  time.Time `gorm:"column:updateTime;default:CURRENT_TIMESTAMP;NOT NULL;comment:'更新时间'" json:"updateTime,omitempty"`
	IsDelete    int8      `gorm:"column:isDelete;default:0;NOT NULL;comment:'是否删除'" json:"isDelete,omitempty"`
}

func (q *Question) TableName() string {
	return "question"
}
