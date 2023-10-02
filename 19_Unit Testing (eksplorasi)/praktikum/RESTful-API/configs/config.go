package configs

import (
	"fmt"
	"restful-api/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

const (
	DB_Username = "root"
	DB_Password = "root123"
	DB_Port     = "3306"
	DB_Host     = "localhost"
	DB_Name     = "crud_go"

	//DB_TEST
	DB_USER_TEST = "root"
	DB_PASS_TEST = "root123"
	DB_HOST_TEST = "127.0.0.1"
	DB_PORT_TEST = "3306"
	DB_NAME_TEST = "acp7_test"
)

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func InitDB() {
	config := Config{
		DB_Username: DB_Username,
		DB_Password: DB_Password,
		DB_Port:     DB_Port,
		DB_Host:     DB_Host,
		DB_Name:     DB_Name,
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func InitialMigration() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Book{})
	DB.AutoMigrate(&models.Blog{})
}

func InitDBTest() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DB_USER_TEST, DB_PASS_TEST, DB_HOST_TEST, DB_PORT_TEST, DB_NAME_TEST)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitMigrateTest()
}

func InitMigrateTest() {
	DB.Migrator().DropTable(&models.User{})
	DB.Migrator().DropTable(&models.Book{})
	DB.Migrator().DropTable(&models.Blog{})
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Book{})
	DB.AutoMigrate(&models.Blog{})
}
