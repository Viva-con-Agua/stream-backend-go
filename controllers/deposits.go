package controllers

import (
	"net/http"

	"stream-backend-go/database"
	"stream-backend-go/models"
	"stream-backend-go/utils"

	"github.com/Viva-con-Agua/echo-pool/resp"
	"github.com/labstack/echo"
)

func GetDepositById(c echo.Context) (err error) {
	uuid := c.Param("id")
	response, err := database.GetDepositById(uuid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resp.InternelServerError)
	}
	if response == nil {
		return c.JSON(http.StatusNoContent, resp.NoContent(uuid))
	}
	return c.JSON(http.StatusOK, response)
}

/**
 * Response list of models.Deposit
 */
func GetDepositCount(c echo.Context) (err error) {
	query := new(models.QueryDeposit)
	if err = c.Bind(query); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	query.Defaults()
	page := query.Page()
	sort := query.OrderBy()
	filter := query.Filter()
	response, err := database.GetDepositCount(page, sort, filter)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resp.InternelServerError)
	}
	return c.JSON(http.StatusOK, response)
}

/**
 * Response list of models.Deposit
 */
func GetDepositList(c echo.Context) (err error) {
	query := new(models.QueryDeposit)
	if err = c.Bind(query); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	query.Defaults()
	page := query.Page()
	sort := query.OrderBy()
	filter := query.Filter()
	response, err := database.GetDepositList(page, sort, filter)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resp.InternelServerError)
	}
	return c.JSON(http.StatusOK, response)
}

func UpdateDeposit(c echo.Context) (err error) {
	// create body as models.Deposit
	body := new(models.Deposit)
	// save data to body
	if err = c.Bind(body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	// validate body
	if err = c.Validate(body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	// update body into database
	if err = database.UpdateDeposit(body); err != nil {
		if err == utils.ErrorNotFound {
			return c.JSON(http.StatusNoContent, resp.NoContent(body.Uuid))
		}
		return c.JSON(http.StatusInternalServerError, resp.InternelServerError())
	}
	// response created
	return c.JSON(http.StatusOK, resp.Updated(body.Uuid))
}

func CreateDeposit(c echo.Context) (err error) {
	// create body as models.DepositCreate
	body := new(models.DepositCreate)
	// save data to body
	if err = c.Bind(body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	// validate body
	if err = c.Validate(body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	// update body into database
	if err = database.CreateDeposit(body); err != nil {
		if err == utils.ErrorConflict {
			return c.JSON(http.StatusNoContent, resp.Conflict())
		}
		return c.JSON(http.StatusInternalServerError, resp.InternelServerError())
	}
	// response created
	return c.JSON(http.StatusOK, resp.Created())
}

func ConfirmDeposit(c echo.Context) (err error) {
	// create body as models.DepositCreate
	body := new(models.DepositCreate)
	// save data to body
	if err = c.Bind(body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	// validate body
	if err = c.Validate(body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	// update body into database
	if err = database.CreateDeposit(body); err != nil {
		if err == utils.ErrorConflict {
			return c.JSON(http.StatusNoContent, resp.Conflict())
		}
		return c.JSON(http.StatusInternalServerError, resp.InternelServerError())
	}
	// response created
	return c.JSON(http.StatusOK, resp.Created())
}
