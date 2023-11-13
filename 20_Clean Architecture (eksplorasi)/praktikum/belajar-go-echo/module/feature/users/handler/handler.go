// module/feature/users/handler/user_handler.go

package handler

import (
	models "belajar-go-echo/module/entities"
	"belajar-go-echo/module/feature/users/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (uh *UserHandler) GetAllUsers(c echo.Context) error {
	users, err := uh.userService.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to get users"})
	}

	response := models.NewListUsersResponse(users)

	return c.JSON(http.StatusOK, response)
}

func (uh *UserHandler) CreateUser(c echo.Context) error {
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	err := uh.userService.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to create user"})
	}

	return c.JSON(http.StatusCreated, echo.Map{"message": "User created successfully", "data": user})
}

func (uh *UserHandler) Login(c echo.Context) error {
	var loginRequest models.User
	if err := c.Bind(&loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	token, err := uh.userService.Authenticate(loginRequest.Email, loginRequest.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Authentication failed"})
	}

	response := echo.Map{"token": token, "email": loginRequest.Email}

	return c.JSON(http.StatusOK, response)
}
