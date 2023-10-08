package main

import (
	"belajar-go-echo/config"
	"belajar-go-echo/routes"

	"github.com/labstack/echo/v4"
)

func main() {

	config.ConnectDB()
	config.MigrateDB()

	app := echo.New()
	routes.RouteUser(app.Group("/users"))

	app.Start(":8080")
}
