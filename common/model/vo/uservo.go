package vo

import (
	"OJ/common/model/entity"
	"time"
)

type UserVo struct {
	UserName    string    `gorm:"column:userName;type:varchar(256);comment:用户昵称" json:"userName"`
	Id          int64     `gorm:"column:id;type:bigint(20);AUTO_INCREMENT;comment:id;primary_key" json:"id"`
	UserAccount string    `gorm:"column:userAccount;type:varchar(256);comment:账号" json:"userAccount"`
	AvatarUrl   string    `gorm:"column:avatarUrl;type:varchar(1024);comment:用户头像" json:"avatarUrl"`
	Gender      int32     `gorm:"column:gender;type:tinyint(4);comment:性别" json:"gender"`
	Phone       string    `gorm:"column:phone;type:varchar(128);comment:电话" json:"phone"`
	Email       string    `gorm:"column:email;type:varchar(512);comment:邮箱" json:"email"`
	UserStatus  int       `gorm:"column:userStatus;type:int(11);default:0;comment:状态 0 - 正常;NOT NULL" json:"userStatus"`
	CreateTime  time.Time `gorm:"column:createTime;type:datetime;default:CURRENT_TIMESTAMP;comment:创建时间" json:"createTime"`
	UpdateTime  time.Time `gorm:"column:updateTime;type:datetime;default:CURRENT_TIMESTAMP" json:"updateTime"`
	UserRole    int       `gorm:"column:userRole;type:int(11);default:0;comment:用户角色 0 - 普通用户 1 - 管理员;NOT NULL" json:"userRole"`
}

func (u *UserVo) User2UserVo(m entity.User) {
	u.Id = m.Id
	u.UserAccount = m.UserAccount.String
	u.UserName = m.UserName.String
	u.AvatarUrl = m.AvatarUrl.String
	u.Email = m.Email.String
	u.Gender = m.Gender.Int32
	u.Phone = m.Phone.String
	u.UserRole = m.UserRole
	u.UserStatus = m.UserStatus
	u.CreateTime = m.CreateTime.Time
	u.UpdateTime = m.UpdateTime.Time
}
