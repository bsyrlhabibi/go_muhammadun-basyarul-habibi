package controllers

import (
	"net/http"
	"strconv"

	"restful-api/configs"
	"restful-api/middleware"
	"restful-api/models"

	"github.com/labstack/echo/v4"
)

// get all users
func GetUsersController(c echo.Context) error {
	var users []models.User

	if err := configs.DB.Preload("Blogs").Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	// your solution here
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}

	var user models.User
	if err := configs.DB.Preload("Blogs").First(&user, userID).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user by ID",
		"user":    user,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	if err := configs.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	// your solution here
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}

	var user models.User
	if err := configs.DB.First(&user, userID).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}

	if err := configs.DB.Delete(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user",
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	// your solution here
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}

	var user models.User
	if err := configs.DB.First(&user, userID).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}

	newUser := new(models.User)
	if err := c.Bind(newUser); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	updates := make(map[string]interface{})

	if newUser.Name != "" {
		updates["name"] = newUser.Name
	}
	if newUser.Email != "" {
		updates["email"] = newUser.Email
	}
	if newUser.Password != "" {
		updates["password"] = newUser.Password
	}

	if err := configs.DB.Model(&user).Updates(updates).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user",
		"user":    user,
	})
}

func LoginUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	err := configs.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "fail login",
			"error":   err.Error(),
		})
	}

	token, err := middleware.CreateToken(int(user.ID), user.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "fail login",
			"error":   err.Error(),
		})
	}

	userResponse := models.UserResponse{int(user.ID), user.Name, user.Email, token}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    userResponse,
	})
}
