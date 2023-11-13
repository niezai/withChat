package model

import "gorm.io/gorm"

type UserFriends struct {
	gorm.Model
	UserId   int32 `json:"user_id"`
	FriendId int32 `json:"friend_id"`
}
