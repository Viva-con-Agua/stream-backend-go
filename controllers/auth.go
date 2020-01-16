package controllers

import (
	"net/http"

	"../database"
	"../models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
)

func SignIn(c echo.Context) (err error) {
	u := new(models.UserSignIn)
	if err = c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err = c.Validate(u); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	sess, _ := session.Get("session", c)
	if sess.Values["valid"] == nil || sess.Values["valid"] == false {
		sess.Values["valid"] = true
		sess.Save(c.Request(), c.Response())
	}
	return c.JSON(http.StatusOK, "login")
}

func SignOut(c echo.Context) (err error) {
	sess, _ := session.Get("session", c)
	if sess.Values["valid"] != nil || sess.Values["valid"] == true {
		sess.Values["valid"] = false
		sess.Save(c.Request(), c.Response())
	}
	return c.JSON(http.StatusOK, "logout")
}

/**
 * SignUp User
 */
func SignUp(c echo.Context) (err error) {
	//u is new UserSignUp
	u := new(models.UserCreate)

	//Bind Body to u
	if err = c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	// validate u as UserSignUp
	if err = c.Validate(u); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	// insert u into database
	if err = database.SignUp(u); err != nil {
		return c.JSON(http.StatusInternalServerError, models.InternelServerError())
	}
	//create new uuid
	return c.JSON(http.StatusCreated, models.Created())
}
