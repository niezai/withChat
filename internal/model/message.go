package model

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	FromUserId  int32  `json:"from_user_id"`
	ToUserId    int32  `json:"to_user_id"`
	Content     string `json:"content"`
	Url         string `json:"url"`
	Picture     string `json:"picture"`
	MessageType int    `json:"message_type"`
	ContentType int    `json:"content_type"`
}
