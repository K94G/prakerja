package controller

import (
	"net/http"
	"prakerja_kg/config"
	"prakerja_kg/model"

	"github.com/labstack/echo/v4"
)

func CreateContact(c echo.Context) error {
	var contact model.Contact
	c.Bind(&contact)

	res := config.DB.Create(&contact)

	if res.Error != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			Message: "failed", Data: nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		Message: "sucess", Data: contact,
	})
}

func GetContacts(c echo.Context) error {
	var contacts []model.Contact

	res := config.DB.Find(&contacts)

	if res.Error != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			Message: "failed", Data: nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		Message: "sucess", Data: contacts,
	})
}

func GetContactById(c echo.Context) error {
	var contacts []model.Contact

	// id := c.Param("id")
	res := config.DB.Find(&contacts)

	if res.Error != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			Message: "failed", Data: nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		Message: "sucess", Data: contacts,
	})
}
