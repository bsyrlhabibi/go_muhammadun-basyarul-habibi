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

type BookResponse struct {
	Message string
	Data    []models.Book
}

func InsertDataBookForGetBooks() error {
	book := models.Book{
		Title:     "Unit Testing",
		Author:    "master fu",
		Publisher: "go.org",
	}

	var err error
	if err = configs.DB.Save(&book).Error; err != nil {
		return err
	}
	return nil
}

// get all books
func TestGetBooks(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "get books",
			path:       "/books",
			expectCode: http.StatusOK,
		},
	}

	e := InitEchoTestAPI()
	InsertDataBookForGetBooks()

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, testCase.path, nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			if assert.NoError(t, GetBooksController(c)) {
				assert.Equal(t, testCase.expectCode, rec.Code)
			}
		})
	}
}

// get book by id
func TestGetBookById(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "get users",
			path:       "/books/:id",
			expectCode: http.StatusOK,
		},
	}

	e := InitEchoTestAPI()
	books := InsertDataBookForGetBooks()

	bookJSON, _ := json.Marshal(books)
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/books/:id", strings.NewReader(string(bookJSON)))
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath(testCase.path)
			c.SetParamNames("id")
			c.SetParamValues("1")

			if assert.NoError(t, GetBookController(c)) {
				assert.Equal(t, testCase.expectCode, rec.Code)
			}
		})
	}
}

// get book by id invalid
func TestGetBookByIdInvalid(t *testing.T) {
	e := InitEchoTestAPI()

	t.Run("Invalid Book ID", func(t *testing.T) {
		// Menggunakan ID yang tidak valid
		req := httptest.NewRequest(http.MethodGet, "/books/invalid_id", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if !assert.Error(t, GetBookController(c)) {
			t.Errorf("Expected an error, but got nil")
		}
	})
}

// create new book
func TestCreateBook(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "create user",
			path:       "/books/:id",
			expectCode: http.StatusOK,
		},
	}
	bookJSON := `{"title":"Create Book","author":"Booksi","publisher":"bok.com"}`
	e := InitEchoTestAPI()
	InsertDataBookForGetBooks()

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/books/:id", strings.NewReader(bookJSON))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath(testCase.path)
			c.SetParamNames("id")
			c.SetParamValues("2")

			if assert.NoError(t, CreateBookController(c)) {
				assert.Equal(t, testCase.expectCode, rec.Code)
			}
		})
	}
}

// delete book
func TestDeleteBook(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "delete users",
			path:       "/books:id",
			expectCode: http.StatusOK,
		},
	}

	e := InitEchoTestAPI()
	books := InsertDataBookForGetBooks()

	bookJSON, _ := json.Marshal(books)
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodDelete, "/books/:id", strings.NewReader(string(bookJSON)))
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath(testCase.path)
			c.SetParamNames("id")
			c.SetParamValues("1")

			if assert.NoError(t, DeleteBookController(c)) {
				assert.Equal(t, testCase.expectCode, rec.Code)
			}
		})
	}
}

// delete book invalid
func TestDeleteBookInvalid(t *testing.T) {
	e := InitEchoTestAPI()

	t.Run("User Not Found", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/books/5", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("5")

		if !assert.Error(t, UpdateBookController(c)) {
			t.Errorf("Expected an error, but got nil")
		}
	})
}

// update book by id
func TestUpdateBook(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "update 1 book",
			path:       "/books/:id",
			expectCode: http.StatusOK,
		},
	}
	bookJSON := `{"title":"Update Unit Testing","author":"unit update","publisher":"go.org update"}`
	e := InitEchoTestAPI()
	InsertDataBookForGetBooks()
	req := httptest.NewRequest(http.MethodPut, "/books/:id", strings.NewReader(bookJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues("1")

		if assert.NoError(t, UpdateBookController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		}
	}
}

// update book invalid
func TestUpdateBookInvalid(t *testing.T) {
	e := InitEchoTestAPI()

	t.Run("Book Not Found", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/books/5", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("5")

		if !assert.Error(t, UpdateBookController(c)) {
			t.Errorf("Expected an error, but got nil")
		}
	})
}
