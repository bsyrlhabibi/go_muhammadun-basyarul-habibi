package controllers

import (
	"net/http"
	"strconv"

	"orm/configs"
	"orm/models"

	"github.com/labstack/echo/v4"
)

// create new blog
func CreateBlogController(c echo.Context) error {
	blog := models.Blog{}
	c.Bind(&blog)

	if err := configs.DB.Save(&blog).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new blog",
		"blog":    blog,
	})
}

// get all blogs
func GetBlogsController(c echo.Context) error {
	var blogs []models.Blog

	if err := configs.DB.Find(&blogs).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all blogs",
		"blogs":   blogs,
	})
}

// get blog by id
func GetBlogController(c echo.Context) error {
	blogID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid blog ID")
	}

	var blog models.Blog
	if err := configs.DB.Preload("User").First(&blog, blogID).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Blog not found")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get blog by ID",
		"blog":    blog,
	})
}

// update blog by id
func UpdateBlogController(c echo.Context) error {
	blogID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid blog ID")
	}

	var blog models.Blog
	if err := configs.DB.First(&blog, blogID).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Blog not found")
	}

	newBlog := new(models.Blog)
	if err := c.Bind(newBlog); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	updates := make(map[string]interface{})

	if newBlog.Title != "" {
		updates["Title"] = newBlog.Title
	}

	if newBlog.Content != "" {
		updates["Content"] = newBlog.Content
	}

	if newBlog.UserID != 0 {
		updates["UserID"] = newBlog.UserID
	}

	if err := configs.DB.Model(&blog).Updates(updates).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update blog",
		"blog":    blog,
	})
}

// delete blog by id
func DeleteBlogController(c echo.Context) error {
	blogID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid blog ID")
	}

	var blog models.Blog
	if err := configs.DB.First(&blog, blogID).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Blog not found")
	}

	if err := configs.DB.Delete(&blog).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete blog",
	})
}
