package controller

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func validateJwt(c echo.Context) bool {
	tokenString := c.Request().Header.Get("Authorization")[7:]

	if tokenString == "" {
		return false
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return false
	}

	_, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return false
	}

	return true
}
