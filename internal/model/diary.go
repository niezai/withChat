package model

import "gorm.io/gorm"

type Diary struct {
	gorm.Model
	DiaryId      string `json:"diary_id"`
	UserId       string `json:"user_id"`
	DiaryTitle   string `json:"diary_title"`
	DiaryLabel   string `json:"diary_label"`
	DiaryContent string `json:"diary_content"`
}
