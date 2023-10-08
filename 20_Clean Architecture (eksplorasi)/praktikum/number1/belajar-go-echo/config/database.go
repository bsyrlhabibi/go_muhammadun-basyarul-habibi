package config

import (
	"belajar-go-echo/features/users/data"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	MigrateDB()
}

func MigrateDB() {
	DB.Migrator().DropTable(&data.User{})
	DB.AutoMigrate(&data.User{})
}
