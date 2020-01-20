package controllers

import (
	"net/http"

	"../database"
	"../models"
	"../utils"
	"github.com/labstack/echo"
)

type ()

/**
 * join user to role
 */
func JoinUserRole(c echo.Context) (err error) {

	// create body as models.Role
	body := new(models.AssignBody)
	// save data to body
	if err = c.Bind(body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	// validate body
	if err = c.Validate(body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	// insert body into database
	if err = database.JoinUserRole(body); err != nil {
		return c.JSON(http.StatusInternalServerError, models.InternelServerError)
	}
	// response created
	return c.JSON(http.StatusCreated, models.Created())
}

func GetUser(c echo.Context) (err error) {
	uuid := c.Param("id")
	response, err := database.GetUser(uuid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.InternelServerError)
	}
	if response == nil {
		return c.JSON(http.StatusNoContent, models.NoContent(uuid))
	}
	return c.JSON(http.StatusOK, response)
}

/**
 * Response list of models.User
 */
func GetUserList(c echo.Context) (err error) {
	query := new(models.QueryUser)
	if err = c.Bind(query); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	query.Defaults()
	page := query.Page()
	sort := query.OrderBy()
	filter := query.Filter()
	response, err := database.GetUserList(page, sort, filter)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.InternelServerError)
	}
	return c.JSON(http.StatusOK, response)
}

func UpdateUser(c echo.Context) (err error) {
	// create body as models.User
	body := new(models.User)
	// save data to body
	if err = c.Bind(body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	// validate body
	if err = c.Validate(body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	// update body into database
	if err = database.UpdateUser(body); err != nil {
		if err == utils.ErrorNotFound {
			return c.JSON(http.StatusNoContent, models.NoContent(body.Uuid))
		}
		return c.JSON(http.StatusInternalServerError, models.InternelServerError())
	}
	// response created
	return c.JSON(http.StatusOK, models.Updated(body.Uuid))
}

func DeleteUser(c echo.Context) (err error) {
	// create body as models.User
	body := new(models.DeleteBody)
	// save data to body
	if err = c.Bind(body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	// validate body
	if err = c.Validate(body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	// update body into database
	if err = database.DeleteUser(body); err != nil {
		if err == utils.ErrorNotFound {
			return c.JSON(http.StatusNoContent, models.NoContent(body.Uuid))
		}
		return c.JSON(http.StatusInternalServerError, models.InternelServerError())
	}
	// response created
	return c.JSON(http.StatusOK, models.Deleted(body.Uuid))
}
