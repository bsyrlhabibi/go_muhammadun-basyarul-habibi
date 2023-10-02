package routes

import (
	"restful-api/constants"
	"restful-api/controllers"
	m "restful-api/middleware"

	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	m.LogMiddleware(e)

	//jwt route users
	eJwt := e.Group("")
	eJwt.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	eJwt.GET("/users", controllers.GetUsersController)
	eJwt.DELETE("/users/:id", controllers.DeleteUserController)
	eJwt.PUT("/users/:id", controllers.UpdateUserController)
	eJwt.GET("/users/:id", controllers.GetUserController)
	eJwt.GET("/books", controllers.GetBooksController)

	//Not Authenticated
	e.POST("/users", controllers.CreateUserController)
	e.POST("/users/login", controllers.LoginUserController)

	//jwt route books
	e.GET("/books", controllers.GetBooksController, m.JWTMiddleware)
	e.GET("/books/:id", controllers.GetBookController, m.JWTMiddleware)
	e.POST("/books", controllers.CreateBookController, m.JWTMiddleware)
	e.DELETE("/books/:id", controllers.DeleteBookController, m.JWTMiddleware)
	e.PUT("/books/:id", controllers.UpdateBookController, m.JWTMiddleware)

	//route blogs
	e.POST("/blogs", controllers.CreateBlogController)
	e.GET("/blogs", controllers.GetBlogsController)
	e.GET("/blogs/:id", controllers.GetBlogController)
	e.PUT("/blogs/:id", controllers.UpdateBlogController)
	e.DELETE("/blogs/:id", controllers.DeleteBlogController)

	//auth basic
	eAuthBasic := e.Group("/auth")
	eAuthBasic.Use(mid.BasicAuth(m.BasicAuthDB))
	eAuthBasic.GET("/users", controllers.GetUsersController)

	return e
}
