package dao

import (
	"context"
	"gorm.io/gorm"
	"withChat/internal/model"
)

type UserDao struct {
	*gorm.DB
}

func (d *UserDao) FindUserByUserNameAndPwd(user *model.LoginUser) (err error) {
	err = d.DB.Table("user").Where("user_uuid = ? and user_name = ? and password = ?", user.UserUuid, user.UserName, user.Password).Find(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{NewDBClient(ctx)}
}

func NewUserDaoByDB(db *gorm.DB) *UserDao {

	return &UserDao{db}
}
