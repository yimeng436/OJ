package repository

import (
	"fmt"
	"testing"
	"usersvr/config"
	"usersvr/log"
	"usersvr/middleware/db"
)

func TestGorm(t *testing.T) {
	err := config.Init()
	if err != nil {
		log.Fatalf("init config failed, err:%v\n", err)
	}
	log.InitLog()
	log.Infof("config loadding success")
	var user User
	db := db.GetDB()
	db.Model(User{}).Where("userName=?", "yimeng").First(&user)
	fmt.Println(user)
}
