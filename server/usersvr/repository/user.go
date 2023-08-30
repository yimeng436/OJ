package repository

import "usersvr/middleware/db"

func GetUserByUserName(userName string) (User, error) {
	db := db.GetDB()
	var user User
	var err error
	err = db.Model(&User{}).Where("userName = ?", userName).First(&user).Error

	if err != nil {
		return user, err
	}

	return user, err
}
