package entity

import (
	"database/sql"
)

type User struct {
	UserName     sql.NullString `gorm:"column:userName;type:varchar(256);comment:用户昵称" json:"userName,omitempty"`
	Id           int64          `gorm:"column:id;type:bigint(20);AUTO_INCREMENT;comment:id;primary_key" json:"id,omitempty"`
	UserAccount  sql.NullString `gorm:"column:userAccount;type:varchar(256);comment:账号" json:"userAccount,omitempty"`
	AvatarUrl    sql.NullString `gorm:"column:avatarUrl;type:varchar(1024);comment:用户头像" json:"avatarUrl,omitempty"`
	Gender       sql.NullInt32  `gorm:"column:gender;type:tinyint(4);comment:性别" json:"gender,omitempty"`
	UserPassword string         `gorm:"column:userPassword;type:varchar(512);comment:密码;NOT NULL" json:"userPassword,omitempty"`
	Phone        sql.NullString `gorm:"column:phone;type:varchar(128);comment:电话" json:"phone,omitempty"`
	Email        sql.NullString `gorm:"column:email;type:varchar(512);comment:邮箱" json:"email,omitempty"`
	UserStatus   int            `gorm:"column:userStatus;type:int(11);default:0;comment:状态 0 - 正常;NOT NULL" json:"userStatus,omitempty"`
	CreateTime   sql.NullTime   `gorm:"column:createTime;type:datetime;default:CURRENT_TIMESTAMP;comment:创建时间" json:"createTime,omitempty"`
	UpdateTime   sql.NullTime   `gorm:"column:updateTime;type:datetime;default:CURRENT_TIMESTAMP" json:"updateTime,omitempty"`
	IsDelete     int            `gorm:"column:isDelete;type:tinyint(4);default:0;comment:是否删除;NOT NULL" json:"isDelete,omitempty"`
	UserRole     int            `gorm:"column:userRole;type:int(11);default:0;comment:用户角色 0 - 普通用户 1 - 管理员;NOT NULL" json:"userRole,omitempty"`
}

func (m *User) TableName() string {
	return "user"
}
