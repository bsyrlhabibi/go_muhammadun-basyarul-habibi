// routes/user_routes.go
package routes

import (
	"belajar-go-echo/module/feature/users/handler"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SetupUserRoutes(e *echo.Echo, uh *handler.UserHandler, db *gorm.DB, jwtMiddleware echo.MiddlewareFunc) {
	userGroup := e.Group("/users")

	userGroup.Use(jwtMiddleware)

	userGroup.GET("", uh.GetAllUsers)

	e.POST("/register", uh.CreateUser)
	e.POST("/login", uh.Login)
}
