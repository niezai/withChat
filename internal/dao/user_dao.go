package dao

import (
	"context"
	"gorm.io/gorm"
	"withChat/internal/db/mysql"
	"withChat/internal/model"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{mysql.NewDBClient(ctx)}
}

func NewUserDaoByDB(db *gorm.DB) *UserDao {

	return &UserDao{db}
}

func (d *UserDao) FindUserByUserNameAndPwd(user *model.User) (err error) {
	err = d.DB.Table("user").Where("email = ? and user_name = ? and password = ?", user.Email, user.UserName, user.Password).Find(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *UserDao) CreateUser(user *model.User) error {
	return d.Table("user").Save(&user).Error
}

func (d *UserDao) FindUserByUuid(uuid string) (*model.User, error) {
	var user *model.User
	err := d.Table("user").Where("user_uuid = ", uuid).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
