package model

import "gorm.io/gorm"

type Comments struct {
	gorm.Model
	TypeId        int    `json:"type_id"`
	CommentId     int    `json:"comment_id"`
	UserId        string `json:"user_id"`
	EventId       string `json:"event_id"`
	RootCommentId string `json:"root_comment_id"`
	ToCommentId   string `json:"to_comment_id"`
	Content       string `json:"content"`
}
