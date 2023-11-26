package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserUuid string `json:"user_uuid"`
	UserName string `json:"user_name"`
	Password string `json:"pass_word"`
	Nickname string `json:"nick_name"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
}

type LoginUser struct {
	UserUuid string `json:"user_uuid"`
	UserName string `json:"user_name"`
	Password string `json:"pass_word"`
	Nickname string `json:"nick_name"`
}
