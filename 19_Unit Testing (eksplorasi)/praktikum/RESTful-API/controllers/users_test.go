package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"restful-api/configs"
	"restful-api/models"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func InitEchoTestAPI() *echo.Echo {
	configs.InitDBTest()
	e := echo.New()
	return e
}

type UserResponse struct {
	Message string
	Data    []models.User
}

func InsertDataUserForGetUsers() error {
	user := models.User{
		Name:     "mamat",
		Password: "mamat",
		Email:    "mamat@gmail.com",
	}

	var err error
	if err = configs.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

// get all users
func TestGetUsers(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "get users",
			path:       "/users",
			expectCode: http.StatusOK,
		},
	}

	e := InitEchoTestAPI()
	InsertDataUserForGetUsers()

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, testCase.path, nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			if assert.NoError(t, GetUsersController(c)) {
				assert.Equal(t, testCase.expectCode, rec.Code)
			}
		})
	}
}

// get user by id
func TestGetUserById(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "get users",
			path:       "/users/:id",
			expectCode: http.StatusOK,
		},
	}

	e := InitEchoTestAPI()
	users := InsertDataUserForGetUsers()

	userJSON, _ := json.Marshal(users)
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/users/:id", strings.NewReader(string(userJSON)))
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath(testCase.path)
			c.SetParamNames("id")
			c.SetParamValues("1")

			if assert.NoError(t, GetUserController(c)) {
				assert.Equal(t, testCase.expectCode, rec.Code)
			}
		})
	}
}

// get user by id invalid
func TestGetUserByIdInvalid(t *testing.T) {
	e := InitEchoTestAPI()

	t.Run("Invalid User ID", func(t *testing.T) {
		// Menggunakan ID yang tidak valid
		req := httptest.NewRequest(http.MethodGet, "/users/invalid_id", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if !assert.Error(t, UpdateUserController(c)) {
			t.Errorf("Expected an error, but got nil")
		}
	})
}

// create new user
func TestCreateUser(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "create user",
			path:       "/users/:id",
			expectCode: http.StatusOK,
		},
	}
	userJSON := `{"name":"mamat create","email":"mamatcreage@gmail.com","password":"123456"}`
	e := InitEchoTestAPI()
	InsertDataUserForGetUsers()

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/users/:id", strings.NewReader(userJSON))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath(testCase.path)
			c.SetParamNames("id")
			c.SetParamValues("2")

			if assert.NoError(t, CreateUserController(c)) {
				assert.Equal(t, testCase.expectCode, rec.Code)
			}
		})
	}
}

// delete user
func TestDeleteUser(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "delete users",
			path:       "/users:id",
			expectCode: http.StatusOK,
		},
	}

	e := InitEchoTestAPI()
	users := InsertDataUserForGetUsers()

	userJSON, _ := json.Marshal(users)
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodDelete, "/users/:id", strings.NewReader(string(userJSON)))
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath(testCase.path)
			c.SetParamNames("id")
			c.SetParamValues("1")

			if assert.NoError(t, DeleteUserController(c)) {
				assert.Equal(t, testCase.expectCode, rec.Code)
			}
		})
	}
}

// delete user invalid
func TestDeleteUserInvalid(t *testing.T) {
	e := InitEchoTestAPI()

	t.Run("User Not Found", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/users/5", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("5")

		if !assert.Error(t, UpdateUserController(c)) {
			t.Errorf("Expected an error, but got nil")
		}
	})
}

// update user by id
func TestUpdateUser(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "update 1 user",
			path:       "/users/:id",
			expectCode: http.StatusOK,
		},
	}
	userJSON := `{"name":"Mamat Update","email":"mamatupdate@gmail.com","password":"mamatupdate"}`
	e := InitEchoTestAPI()
	InsertDataUserForGetUsers()
	req := httptest.NewRequest(http.MethodPut, "/users/:id", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues("1")

		if assert.NoError(t, UpdateUserController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		}
	}
}

// update user invalid
func TestUpdateUserInvalid(t *testing.T) {
	e := InitEchoTestAPI()

	t.Run("User Not Found", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/users/5", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("5")

		if !assert.Error(t, UpdateUserController(c)) {
			t.Errorf("Expected an error, but got nil")
		}
	})
}
