package model

import "gorm.io/gorm"

type UserFriend struct {
	gorm.Model
	UserId   int32 `json:"user_id"`
	FriendId int32 `json:"friend_id"`
}
