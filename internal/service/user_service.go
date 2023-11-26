package service

import (
	"context"
	"withChat/internal/dao"
	"withChat/internal/model"
)

type UserService struct {
	UserUuid string `json:"user_uuid"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (userService *UserService) Register(ctx context.Context) {

}

func (userService *UserService) Login(user *model.LoginUser, ctx context.Context) (err error) {
	userDao := dao.NewUserDao(ctx)
	return userDao.FindUserByUserNameAndPwd(user)
}
