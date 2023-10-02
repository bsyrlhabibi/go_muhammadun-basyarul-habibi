package main

import (
	"restful-api/configs"
	"restful-api/routes"
)

func init() {
	configs.InitDB()
	configs.InitialMigration()
}

func main() {
	e := routes.New()

	e.Logger.Fatal(e.Start(":8000"))
}
