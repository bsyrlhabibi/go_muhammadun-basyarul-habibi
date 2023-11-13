// main.go
package main

import (
	models "belajar-go-echo/module/entities"
	"belajar-go-echo/module/feature/users/handler"
	"belajar-go-echo/module/feature/users/middleware"
	"belajar-go-echo/module/feature/users/repository"
	"belajar-go-echo/module/feature/users/service"
	"belajar-go-echo/routes"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.User{})

	userRepository := repository.NewUserRepository(db)

	secretKey := "your_secret_key"
	userService := service.NewUserService(userRepository, secretKey)

	userHandler := handler.NewUserHandler(userService)

	jwtMiddleware := middleware.JWTMiddleware([]byte(secretKey))

	e := echo.New()
	routes.SetupUserRoutes(e, userHandler, db, jwtMiddleware)

	e.Start(":8080")
}
