package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	ArticleUuid    string `json:"article_uuid"`
	ArticleTitle   string `json:"article_title"`
	ArticleContent string `json:"article_content"`
	ArticleLabel   string `json:"article_label"`
}
