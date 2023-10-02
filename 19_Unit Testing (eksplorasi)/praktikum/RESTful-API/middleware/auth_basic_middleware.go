package middleware

import (
	"restful-api/configs"
	"restful-api/models"

	"github.com/labstack/echo/v4"
)

func BasicAuthDB(username, password string, c echo.Context) (bool, error) {
	var user models.User
	err := configs.DB.Where("email = ? AND password = ?", username, password).First(&user).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
