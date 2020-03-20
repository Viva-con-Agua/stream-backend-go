package controllers

import (
	"net/http"

	"../database"
	"../models"
	"../utils"
	"github.com/Viva-con-Agua/echo-pool/pool"
	"github.com/labstack/echo"
)

/**
 * Response list of models.Taking
 */
func GetTakingList(c echo.Context) (err error) {
	query := new(models.QueryTaking)
	if err = c.Bind(query); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	query.Defaults()
	page := query.Page()
	sort := query.OrderBy()
	filter := query.Filter()
	response, err := database.GetTakingList(page, sort, filter)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, pool.InternelServerError)
	}
	return c.JSON(http.StatusOK, response)
}

/**
 * Response list of models.Taking
 */
func GetTakingCount(c echo.Context) (err error) {
	query := new(models.QueryTaking)
	if err = c.Bind(query); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	query.Defaults()
	page := query.Page()
	sort := query.OrderBy()
	filter := query.Filter()
	response, err := database.GetTakingList(page, sort, filter)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, pool.InternelServerError)
	}
	return c.JSON(http.StatusOK, response)
}

func GetTakingById(c echo.Context) (err error) {
	uuid := c.Param("id")
	response, err := database.GetTakingById(uuid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, pool.InternelServerError)
	}
	if response == nil {
		return c.JSON(http.StatusNoContent, pool.NoContent(uuid))
	}
	return c.JSON(http.StatusOK, response)
}

func CreateTaking(c echo.Context) (err error) {
	// create body as models.TakingCreate
	body := new(models.TakingCreate)
	// save data to body
	if err = c.Bind(body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	// validate body
	if err = c.Validate(body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	// update body into database
	if err = database.CreateTaking(body); err != nil {
		if err == utils.ErrorConflict {
			return c.JSON(http.StatusNoContent, pool.Conflict())
		}
		return c.JSON(http.StatusInternalServerError, pool.InternelServerError())
	}
	// response created
	return c.JSON(http.StatusOK, pool.Created())
}

func UpdateTaking(c echo.Context) (err error) {
	// create body as models.Taking
	body := new(models.Taking)
	// save data to body
	if err = c.Bind(body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	// validate body
	if err = c.Validate(body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	// update body into database
	if err = database.UpdateTaking(body); err != nil {
		if err == utils.ErrorNotFound {
			return c.JSON(http.StatusNoContent, pool.NoContent(body.Uuid))
		}
		return c.JSON(http.StatusInternalServerError, pool.InternelServerError())
	}
	// response created
	return c.JSON(http.StatusOK, pool.Updated(body.Uuid))
}
