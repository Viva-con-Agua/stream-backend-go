package controllers

import (
	"net/http"

	"../database"
	"../models"
	"github.com/labstack/echo"
)

func GetRolesList(c echo.Context) (err error) {
	response, err := database.GetRolesList()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.InternelServerError)
	}
	return c.JSON(http.StatusOK, response)
}

func PostRole(c echo.Context) (err error) {
	// create body as models.Role
	body := new(models.RoleCreate)
	// save data to body
	if err = c.Bind(body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	// validate body
	if err = c.Validate(body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	// insert body into database
	if err = database.PostRole(body); err != nil {
		return c.JSON(http.StatusInternalServerError, models.InternelServerError)
	}
	// response created
	return c.JSON(http.StatusCreated, models.Created())
}
