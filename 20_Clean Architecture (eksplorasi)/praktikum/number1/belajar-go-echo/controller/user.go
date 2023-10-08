package controller

import (
	"belajar-go-echo/features/users/data"
	"belajar-go-echo/repository"

	"github.com/labstack/echo/v4"
)

func GetAllUsers(c echo.Context) error {
	users, err := repository.GetAllUsers()
	if err != nil {
		return c.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(200, echo.Map{
		"data": users,
	})
}

func CreateUser(c echo.Context) error {
	user := data.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(400, echo.Map{
			"error": err.Error(),
		})
	}
	err := repository.CreateUsers(user)
	if err != nil {
		return c.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(201, echo.Map{
		"message": "success create user",
		"data":    user,
	})
}
