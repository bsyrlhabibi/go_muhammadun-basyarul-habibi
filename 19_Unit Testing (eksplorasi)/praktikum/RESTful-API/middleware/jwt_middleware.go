package middleware

import (
	"net/http"
	"restful-api/constants"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func CreateToken(userId int, name string) (string, error) {
	claims := jwt.MapClaims{}
	claims["userId"] = userId
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SECRET_JWT))
}

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenHeader := c.Request().Header.Get("Authorization")
		if tokenHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Token is missing")
		}

		tokenString := strings.TrimPrefix(tokenHeader, "Bearer ")

		_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(constants.SECRET_JWT), nil
		})

		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token or expired token")
		}

		return next(c)
	}
}
