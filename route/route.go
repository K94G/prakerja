package route

import (
	"prakerja_kg/controller"
	"prakerja_kg/docs"

	"github.com/labstack/echo/v4"
)

func InitRoute() *echo.Echo {
	e := echo.New()

	docs.Initialize(e)

	e.POST("/contacts", controller.CreateContact)
	e.GET("/contacts", controller.GetContacts)
	e.GET("/contacts/:id", controller.GetContactById)
	e.PUT("/contacts/:id", controller.UpdateContactById)
	e.PATCH("/contacts/:id", controller.DeactivateContactById)
	e.DELETE("/contacts/:id", controller.DestroyContactById)

	e.POST("/login", controller.Login)

	return e
}
