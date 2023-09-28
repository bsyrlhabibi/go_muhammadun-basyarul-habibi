package controllers

import (
	"net/http"
	"strconv"

	"orm/configs"
	"orm/models"

	"github.com/labstack/echo/v4"
)

// get all books
func GetBooksController(c echo.Context) error {
	var books []models.Book

	if err := configs.DB.Find(&books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all books",
		"books":   books,
	})
}

// get book by id
func GetBookController(c echo.Context) error {
	bookID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid book ID")
	}

	var book models.Book
	if err := configs.DB.First(&book, bookID).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Book not found")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get book by ID",
		"book":    book,
	})
}

// create new book
func CreateBookController(c echo.Context) error {
	book := models.Book{}
	c.Bind(&book)

	if err := configs.DB.Save(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new book",
		"book":    book,
	})
}

// delete book by id
func DeleteBookController(c echo.Context) error {
	bookID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid book ID")
	}

	var book models.Book
	if err := configs.DB.First(&book, bookID).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Book not found")
	}

	if err := configs.DB.Delete(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete book",
	})
}

// update book by id
func UpdateBookController(c echo.Context) error {
	bookID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid book ID")
	}

	var book models.Book
	if err := configs.DB.First(&book, bookID).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Book not found")
	}

	newBook := new(models.Book)
	if err := c.Bind(newBook); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	book.Title = newBook.Title
	book.Author = newBook.Author
	book.Publisher = newBook.Publisher

	if err := configs.DB.Save(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update book",
		"book":    book,
	})
}
