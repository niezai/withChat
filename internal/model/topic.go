package model

import "gorm.io/gorm"

type Topic struct {
	gorm.Model
	TopicUuid    string `json:"topic_uuid"`
	TopicTitle   string `json:"topic_title"`
	TopicContent string `json:"topic_content"`
}
