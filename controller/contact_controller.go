package controller

import (
	"net/http"
	"prakerja_kg/config"
	"prakerja_kg/model"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateContact(c echo.Context) error {
	validate := validateJwt(c)
	if !validate {
		return c.JSON(http.StatusUnauthorized, model.Response{
			Message: "unauthorized", Data: nil,
		})
	}

	var contact model.Contact
	c.Bind(&contact)

	res := config.DB.Create(&contact)

	if res.Error != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			Message: "create failed", Data: nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		Message: "create success", Data: contact,
	})
}

func GetContacts(c echo.Context) error {
	validate := validateJwt(c)
	if !validate {
		return c.JSON(http.StatusUnauthorized, model.Response{
			Message: "unauthorized", Data: nil,
		})
	}

	var contacts []model.Contact

	res := config.DB.Find(&contacts)

	if res.Error != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			Message: "get failed", Data: nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		Message: "get success", Data: contacts,
	})
}

func GetContactById(c echo.Context) error {
	validate := validateJwt(c)
	if !validate {
		return c.JSON(http.StatusUnauthorized, model.Response{
			Message: "unauthorized", Data: nil,
		})
	}

	var contact model.Contact

	id, _ := strconv.Atoi(c.Param("id"))
	res := config.DB.First(&contact, id)

	if res.Error != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			Message: "get failed", Data: nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		Message: "get success", Data: contact,
	})
}

func UpdateContactById(c echo.Context) error {
	validate := validateJwt(c)
	if !validate {
		return c.JSON(http.StatusUnauthorized, model.Response{
			Message: "unauthorized", Data: nil,
		})
	}

	var contact model.Contact

	var body model.Contact
	c.Bind(&body)

	id, _ := strconv.Atoi(c.Param("id"))
	res := config.DB.First(&contact, "id = ?", id)

	if res.Error != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			Message: "get failed", Data: nil,
		})
	}

	contact.Name = body.Name
	contact.Phone = body.Phone
	contact.Email = body.Email
	contact.Address = body.Address
	contact.City = body.City
	contact.Province = body.Province
	contact.PostalCode = body.PostalCode
	contact.GoogleMapsUrl = body.GoogleMapsUrl
	contact.Category = body.Category
	contact.Active = body.Active

	config.DB.Save(&contact)

	if config.DB.Error != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			Message: "update failed", Data: nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		Message: "update success", Data: contact,
	})
}

func DeactivateContactById(c echo.Context) error {
	validate := validateJwt(c)
	if !validate {
		return c.JSON(http.StatusUnauthorized, model.Response{
			Message: "unauthorized", Data: nil,
		})
	}

	var contact model.Contact

	id, _ := strconv.Atoi(c.Param("id"))
	res := config.DB.First(&contact, "id = ?", id)

	if res.Error != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			Message: "get failed", Data: nil,
		})
	}

	contact.Active = false

	config.DB.Save(&contact)

	if config.DB.Error != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			Message: "deactivate failed", Data: nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		Message: "deactivate success", Data: contact,
	})
}

func DestroyContactById(c echo.Context) error {
	validate := validateJwt(c)
	if !validate {
		return c.JSON(http.StatusUnauthorized, model.Response{
			Message: "unauthorized", Data: nil,
		})
	}

	var contact model.Contact

	id, _ := strconv.Atoi(c.Param("id"))
	res := config.DB.Delete(&contact, id)

	if res.Error != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			Message: "get failed", Data: nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		Message: "delete success", Data: nil,
	})
}
