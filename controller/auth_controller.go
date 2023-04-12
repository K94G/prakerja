package controller

import (
	"net/http"
	"prakerja_kg/config"
	"prakerja_kg/model"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	var user model.User

	res := config.DB.First(&user, "username = ?", username)

	if res.Error != nil {
		return c.JSON(http.StatusUnauthorized, model.Response{
			Message: "failed", Data: nil,
		})
	}

	if username == user.Username && password == user.Password {
		token := jwt.New(jwt.SigningMethodHS256)

		claims := token.Claims.(jwt.MapClaims)
		claims["id"] = user.Id
		claims["name"] = user.Name
		claims["username"] = user.Username
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		t, err := token.SignedString([]byte(secretKey))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, model.Response{
				Message: "login failed", Data: nil,
			})
		}

		return c.JSON(http.StatusOK, model.Response{
			Message: "login success", Data: t,
		})
	}

	return c.JSON(http.StatusUnauthorized, model.Response{
		Message: "unauthorized", Data: nil,
	})
}
