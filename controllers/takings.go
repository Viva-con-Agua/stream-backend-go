package controllers

import (
    "net/http"

	"stream-backend-go/database"
	"stream-backend-go/models"
	"stream-backend-go/utils"
	"github.com/Viva-con-Agua/echo-pool/resp"
	"github.com/labstack/echo"
)

/**
 * Response list of models.Taking
 */
func GetTakingList(c echo.Context) (err error) {
<<<<<<< HEAD
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
		return c.JSON(http.StatusInternalServerError, resp.InternelServerError)
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
		return c.JSON(http.StatusInternalServerError, resp.InternelServerError)
	}
	return c.JSON(http.StatusOK, response)
}

func GetTakingById(c echo.Context) (err error) {
	uuid := c.Param("id")
	response, err := database.GetTakingById(uuid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resp.InternelServerError)
	}
	if response == nil {
		return c.JSON(http.StatusNoContent, resp.NoContent(uuid))
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
			return c.JSON(http.StatusNoContent, resp.Conflict())
		}
		return c.JSON(http.StatusInternalServerError, resp.InternelServerError())
	}
	// response created
	return c.JSON(http.StatusOK, resp.Created())
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
			return c.JSON(http.StatusNoContent, resp.NoContent(body.Uuid))
		}
		return c.JSON(http.StatusInternalServerError, resp.InternelServerError())
	}
	// response created
	return c.JSON(http.StatusOK, resp.Updated(body.Uuid))
}
