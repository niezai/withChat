package model

import (
	"gorm.io/gorm"
)

type Groups struct {
	gorm.Model
	UserId      int    `json:"user_id"`
	GroupName   string `json:"name"`
	GroupNotice string `json:"group_notice"`
	GroupUuid   string `json:"group_uuid"`
}
