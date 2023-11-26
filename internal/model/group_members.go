package model

import "gorm.io/gorm"

type GroupMember struct {
	gorm.Model
	UserId   int    `json:"user_id"`
	GroupId  int    `json:"group_id"`
	NickName string `json:"nick_name"`
	Mute     int    `json:"mute"`
}
