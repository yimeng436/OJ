package repository

import (
	"gatewaysvr/middleware/db"
	"github.com/yimeng436/OJ/common/model/entity"
)

func GetUserByUserAccount(userAccount string) (entity.User, error) {
	db := db.GetDB()
	var user entity.User
	err := db.Model(entity.User{}).Where("userAccount=?", userAccount).First(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}
