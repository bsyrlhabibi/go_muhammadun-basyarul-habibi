package models

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	UserID  uint   `json:"user_id" form:"user_id"`
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
}
