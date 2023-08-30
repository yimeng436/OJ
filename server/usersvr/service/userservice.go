package service

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"github.com/yimeng436/OJ/pkg/pb"
	"gorm.io/gorm"
	"usersvr/repository"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func (UserService) UserLogin(ctx context.Context, request *pb.UserLoginRequest) (*pb.UserVo, error) {
	userName := request.UserName
	userPassword := request.UserPassword
	if userName == "" || userPassword == "" {
		return nil, errors.New("用户密码为空")
	}
	user, err := repository.GetUserByUserName(userName)

	if err != nil && err != gorm.ErrRecordNotFound {
		return &pb.UserVo{}, err
	}

	if userPassword != user.UserPassword {
		return &pb.UserVo{}, errors.New("账号密码错误")
	}

	userVo := new(pb.UserVo)
	err = copier.Copy(userVo, &user)
	if err != nil {
		return &pb.UserVo{}, err
	}

	return userVo, nil
}
func (UserService) GetLoginUser(ctx context.Context, request *pb.Empty) (*pb.UserVo, error) {

	return nil, nil
}
