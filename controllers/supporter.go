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
 * join Supporter to role
 */
func JoinSupporterRole(c echo.Context) (err error) {

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
	if err = database.JoinSupporterRole(body); err != nil {
		return c.JSON(http.StatusInternalServerError, pool.InternelServerError)
	}
	// response created
	return c.JSON(http.StatusCreated, pool.Created())
}

func GetSupporter(c echo.Context) (err error) {
	uuid := c.Param("id")
	response, err := database.GetSupporter(uuid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, pool.InternelServerError)
	}
	if response == nil {
		return c.JSON(http.StatusNoContent, pool.NoContent(uuid))
	}
	return c.JSON(http.StatusOK, response)
}

/**
 * Response list of models.Supporter
 */
func GetSupporterList(c echo.Context) (err error) {
	query := new(models.QuerySupporter)
	if err = c.Bind(query); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	query.Defaults()
	page := query.Page()
	sort := query.OrderBy()
	filter := query.Filter()
	response, err := database.GetSupporterList(page, sort, filter)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, pool.InternelServerError)
	}
	return c.JSON(http.StatusOK, response)
}

func UpdateSupporter(c echo.Context) (err error) {
	// create body as models.Supporter
	body := new(models.Supporter)
	// save data to body
	if err = c.Bind(body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	// validate body
	if err = c.Validate(body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	// update body into database
	if err = database.UpdateSupporter(body); err != nil {
		if err == utils.ErrorNotFound {
			return c.JSON(http.StatusNoContent, pool.NoContent(body.Uuid))
		}
		return c.JSON(http.StatusInternalServerError, pool.InternelServerError())
	}
	// response created
	return c.JSON(http.StatusOK, pool.Updated(body.Uuid))
}

func DeleteSupporter(c echo.Context) (err error) {
	// create body as models.Supporter
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
	if err = database.DeleteSupporter(body); err != nil {
		if err == utils.ErrorNotFound {
			return c.JSON(http.StatusNoContent, pool.NoContent(body.Uuid))
		}
		return c.JSON(http.StatusInternalServerError, pool.InternelServerError())
	}
	// response created
	return c.JSON(http.StatusOK, pool.Deleted(body.Uuid))
}
