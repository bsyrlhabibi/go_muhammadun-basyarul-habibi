package controllers_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"restful-api/configs"
	"restful-api/controllers"
	"restful-api/models"
)

// create blog
func TestCreateBlog(t *testing.T) {
	configs.InitDBTest()
	e := echo.New()
	controllers.InsertDataUserForGetUsers()
	blogJSON := `{
		"title": "Test Blog",
		"content": "This is a test blog content",
		"user_id": 1
	}`

	req := httptest.NewRequest(http.MethodPost, "/blogs", bytes.NewBufferString(blogJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, controllers.CreateBlogController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

// get all blogs
func TestGetBlogs(t *testing.T) {

	configs.InitDBTest()

	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/blogs", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, controllers.GetBlogsController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

// get blog by id
func TestGetBlogById(t *testing.T) {

	configs.InitDBTest()

	e := echo.New()

	user := models.User{
		Name:     "John Doe",
		Password: "password",
		Email:    "john.doe@example.com",
	}
	if err := configs.DB.Create(&user).Error; err != nil {
		panic(err)
	}

	blog := models.Blog{
		Title:   "Test Blog",
		Content: "This is a test blog content",
		UserID:  1,
	}
	if err := configs.DB.Save(&blog).Error; err != nil {
		t.Fatal(err)
	}

	blogID := strconv.Itoa(int(blog.ID))
	req := httptest.NewRequest(http.MethodGet, "/blogs/"+blogID, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(blogID)

	if assert.NoError(t, controllers.GetBlogController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

// update blog
func TestUpdateBlog(t *testing.T) {

	configs.InitDBTest()

	e := echo.New()

	user := models.User{
		Name:     "John Doe",
		Password: "password",
		Email:    "john.doe@example.com",
	}
	if err := configs.DB.Create(&user).Error; err != nil {
		panic(err)
	}

	blog := models.Blog{
		Title:   "Test Blog",
		Content: "This is a test blog content",
		UserID:  1,
	}
	if err := configs.DB.Save(&blog).Error; err != nil {
		t.Fatal(err)
	}

	blogID := strconv.Itoa(int(blog.ID))

	updateJSON := `{
		"title": "Updated Blog",
		"content": "This is an updated blog content"
	}`

	req := httptest.NewRequest(http.MethodPut, "/blogs/"+blogID, bytes.NewBufferString(updateJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(blogID)

	if assert.NoError(t, controllers.UpdateBlogController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

// delete blog
func TestDeleteBlog(t *testing.T) {

	configs.InitDBTest()

	e := echo.New()

	user := models.User{
		Name:     "John Doe",
		Password: "password",
		Email:    "john.doe@example.com",
	}
	if err := configs.DB.Create(&user).Error; err != nil {
		panic(err)
	}

	blog := models.Blog{
		Title:   "Test Blog",
		Content: "This is a test blog content",
		UserID:  1,
	}
	if err := configs.DB.Save(&blog).Error; err != nil {
		t.Fatal(err)
	}

	blogID := strconv.Itoa(int(blog.ID))

	req := httptest.NewRequest(http.MethodDelete, "/blogs/"+blogID, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(blogID)

	if assert.NoError(t, controllers.DeleteBlogController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
