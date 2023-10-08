package routes

import (
	"belajar-go-echo/controller"

	"github.com/labstack/echo/v4"
)

func RouteUser(e *echo.Group) {
	e.GET("", controller.GetAllUsers)
	e.POST("", controller.CreateUser)
}
