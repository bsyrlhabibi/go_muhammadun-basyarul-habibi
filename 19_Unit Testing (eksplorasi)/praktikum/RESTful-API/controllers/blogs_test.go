package controllers

import (
	"restful-api/configs"
	"restful-api/models"
)

type BlogResponse struct {
	Message string
	Data    []models.Blog
}

func InsertDataBlogforGetBlogs() error {
	blog := models.Blog{
		UserID:  1,
		Title:   "master fu",
		Content: "go.org",
	}

	var err error
	if err = configs.DB.Save(&blog).Error; err != nil {
		return err
	}
	return nil
}
