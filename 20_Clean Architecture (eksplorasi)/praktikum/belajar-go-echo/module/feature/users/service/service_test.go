package service

import (
	"belajar-go-echo/mocks"
	models "belajar-go-echo/module/entities"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserService_Authenticate(t *testing.T) {
	userRepositoryMock := &mocks.UserRepository{}
	userRepositoryMock.On("GetUserByEmail", "dummy@example.com").Return(&models.User{Password: "dummyPassword"}, nil)

	userService := NewUserService(userRepositoryMock, "your_secret_key")

	token, err := userService.Authenticate("dummy@example.com", "dummyPassword")
	
	userRepositoryMock.AssertExpectations(t)

	assert.Nil(t, err)
	assert.NotEmpty(t, token)
}

func TestUserService_CreateUser(t *testing.T) {
	userRepositoryMock := &mocks.UserRepository{}
	userRepositoryMock.On("CreateUser", mock.Anything).Return(nil)

	userService := NewUserService(userRepositoryMock, "your_secret_key")

	dummyUser := models.User{
		Email:    "dummyUser",
		Password: "dummy@example.com",
	}

	err := userService.CreateUser(dummyUser)

	userRepositoryMock.AssertExpectations(t)

	assert.Nil(t, err)
}

func TestUserService_GetAllUsers(t *testing.T) {
	userRepositoryMock := &mocks.UserRepository{}
	userRepositoryMock.On("GetAllUsers").Return([]models.User{}, nil)

	userService := NewUserService(userRepositoryMock, "your_secret_key")

	users, err := userService.GetAllUsers()

	userRepositoryMock.AssertExpectations(t)

	assert.Nil(t, err)
	assert.NotNil(t, users)
}
