package route

import (
	"prakerja_kg/controller"

	"github.com/labstack/echo/v4"
)

func InitRoute() *echo.Echo {
	e := echo.New()
	e.POST("/contacts", controller.CreateContact)
	e.GET("/contacts", controller.GetContacts)
	e.GET("/contacts/:id", controller.GetContactById)
	return e
}
