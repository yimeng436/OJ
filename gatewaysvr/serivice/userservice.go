package serivice

import (
	"errors"
	"gatewaysvr/repository"
	"github.com/yimeng436/OJ/common/model/entity"
	"github.com/yimeng436/OJ/common/model/vo"
)

func UserLogin(userAccount, password string) (*vo.UserVo, error) {

	var loginuser entity.User
	var err error
	loginuser, err = repository.GetUserByUserAccount(userAccount)
	if err != nil {
		return nil, err
	}

	if password != loginuser.UserPassword {
		return nil, errors.New("用户名密码错误")
	}
	userVo := new(vo.UserVo)
	userVo.User2UserVo(loginuser)

	return userVo, nil
}
