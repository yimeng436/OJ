package repository

import "time"

type User struct {
	UserName     string    `gorm:"column:userName"`     // 用户昵称
	Id           int64     `gorm:"column:id"`           // id
	UserAccount  string    `gorm:"column:userAccount"`  // 账号
	AvatarUrl    string    `gorm:"column:avatarUrl"`    // 用户头像
	Gender       int8      `gorm:"column:gender"`       // 性别
	UserPassword string    `gorm:"column:userPassword"` // 密码
	Phone        string    `gorm:"column:phone"`        // 电话
	Email        string    `gorm:"column:email"`        // 邮箱
	UserStatus   int32     `gorm:"column:userStatus"`   // 状态 0 - 正常
	CreateTime   time.Time `gorm:"column:createTime"`   // 创建时间
	UpdateTime   time.Time `gorm:"column:updateTime"`
	IsDelete     int8      `gorm:"column:isDelete"` // 是否删除
	UserRole     int32     `gorm:"column:userRole"` // 用户角色 0 - 普通用户 1 - 管理员
}

func (u *User) TableName() string {
	return "user"
}
