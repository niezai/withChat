package model

import "gorm.io/gorm"

type Like struct {
	gorm.Model
	LikeId   string `json:"like_id"`
	TypeId   int    `json:"type_id"`
	ToLikeId string `json:"to_like_id"`
	LikeFlag int    `json:"like_flag"`
}
