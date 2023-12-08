package service

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"time"
	"withChat/internal/dao"
	"withChat/internal/db/redis"
	"withChat/internal/model"
	"withChat/internal/util"
)

type UserService struct {
}

func (userService *UserService) Register(user *model.User, ctx context.Context) error {
	return dao.NewUserDao(ctx).CreateUser(user)
}

func (userService *UserService) Login(user *model.User, ctx context.Context) (string, error) {
	userDao := dao.NewUserDao(ctx)
	var authToken string
	var err error
	if user != nil {
		err = userDao.FindUserByUserNameAndPwd(user)
		if err != nil {
			return "", errors.New("请核对登录信息")
		}
		// 有该用户生成 token 存入redis
		authToken, err = util.GenerateToken(user.UserUuid, user.UserName)
		if err != nil {
			return "", err
		}
		// 将token存入到redis中
		r := redis.Redis.Set(user.UserUuid, authToken, 60*60*time.Second)
		if !r {
			logrus.Error("token存入redis失败，err =", r)
		}
		return authToken, nil
	}
	return "", nil
}
