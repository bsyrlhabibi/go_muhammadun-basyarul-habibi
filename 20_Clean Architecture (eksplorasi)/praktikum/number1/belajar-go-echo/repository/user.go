package repository

import (
	"belajar-go-echo/config"
	"belajar-go-echo/features/users/data"
	"fmt"
)

func CreateUsers(user data.User) error {
	res := config.DB.Create(&user)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func GetAllUsers() ([]data.User, error) {
	var users []data.User
	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
