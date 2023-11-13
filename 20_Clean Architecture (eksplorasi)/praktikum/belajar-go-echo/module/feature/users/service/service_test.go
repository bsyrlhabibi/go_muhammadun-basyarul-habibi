package service

import (
	"belajar-go-echo/mocks"
	models "belajar-go-echo/module/entities"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserService_Authenticate(t *testing.T) {
	// Inisialisasi mock userRepository
	userRepositoryMock := &mocks.UserRepository{}
	userRepositoryMock.On("GetUserByEmail", "dummy@example.com").Return(&models.User{Password: "dummyPassword"}, nil)

	// Inisialisasi userService dengan mock userRepository
	userService := NewUserService(userRepositoryMock, "your_secret_key")

	// Panggil metode Authenticate
	token, err := userService.Authenticate("dummy@example.com", "dummyPassword")

	// Verifikasi bahwa metode GetUserByEmail dari userRepository dipanggil dengan benar
	userRepositoryMock.AssertExpectations(t)

	// Verifikasi hasil yang diharapkan
	assert.Nil(t, err)
	assert.NotEmpty(t, token) // Sesuaikan verifikasi ini dengan logika bisnis Anda
}

// func TestAuthenticate(t *testing.T) {
// 	userServiceMock := mocks.NewUserService(t)

// 	userServiceMock.On("Authenticate", "dummy@example.com", "dummyPassword").Return("dummyToken", nil)

// 	token, err := userServiceMock.Authenticate("dummy@example.com", "dummyPassword")

// 	userServiceMock.AssertExpectations(t)

// 	assert.Nil(t, err)
// 	assert.Equal(t, "dummyToken", token)
// }

// func TestCreateUser(t *testing.T) {
// 	userServiceMock := mocks.NewUserService(t)

// 	dummyUser := models.User{
// 		Email:    "dummyUser",
// 		Password: "dummy@example.com",
// 	}

// 	userServiceMock.On("CreateUser", dummyUser).Return(nil)

// 	err := userServiceMock.CreateUser(dummyUser)

// 	userServiceMock.AssertExpectations(t)

// 	assert.Nil(t, err)
// }

// func TestGetAllUsers(t *testing.T) {
// 	userServiceMock := mocks.NewUserService(t)

// 	userList := []models.User{
// 		{Password: "user1", Email: "user1@example.com"},
// 		{Password: "user2", Email: "user2@example.com"},
// 	}
// 	userServiceMock.On("GetAllUsers").Return(userList, nil)

// 	users, err := userServiceMock.GetAllUsers()

// 	userServiceMock.AssertExpectations(t)

// 	assert.Nil(t, err)
// 	assert.Equal(t, userList, users)
// }

func TestUserService_CreateUser(t *testing.T) {
	// Inisialisasi mock userRepository
	userRepositoryMock := &mocks.UserRepository{}
	userRepositoryMock.On("CreateUser", mock.Anything).Return(nil)

	// Inisialisasi userService dengan mock userRepository
	userService := NewUserService(userRepositoryMock, "your_secret_key")

	// Persiapkan dummy user
	dummyUser := models.User{
		Email:    "dummyUser",
		Password: "dummy@example.com",
	}

	// Panggil metode CreateUser
	err := userService.CreateUser(dummyUser)

	// Verifikasi bahwa metode CreateUser dari userRepository dipanggil dengan benar
	userRepositoryMock.AssertExpectations(t)

	// Verifikasi hasil yang diharapkan
	assert.Nil(t, err)
}

func TestUserService_GetAllUsers(t *testing.T) {
	// Inisialisasi mock userRepository
	userRepositoryMock := &mocks.UserRepository{}
	userRepositoryMock.On("GetAllUsers").Return([]models.User{}, nil)

	// Inisialisasi userService dengan mock userRepository
	userService := NewUserService(userRepositoryMock, "your_secret_key")

	// Panggil metode GetAllUsers
	users, err := userService.GetAllUsers()

	// Verifikasi bahwa metode GetAllUsers dari userRepository dipanggil dengan benar
	userRepositoryMock.AssertExpectations(t)

	// Verifikasi hasil yang diharapkan
	assert.Nil(t, err)
	assert.NotNil(t, users) // Sesuaikan verifikasi ini dengan logika bisnis Anda
}
