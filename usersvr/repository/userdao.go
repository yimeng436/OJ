package repository

import (
	"OJ/common/model/entity"
	"OJ/gatewaysvr/middleware/db"
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
