package main

import (
	"./controllers"
	"./models"
	"./utils"
	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

type (
	CustomValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	Role := new(models.Role)
	Role.Name = "admin"
	Roles := new(models.Roles)
	Roles.AddRole(*Role)
	// intial loading function
	utils.LoadConfig()
	utils.ConnectDatabase()
	store := utils.RedisSession()

	//create echo server
	e := echo.New()
	e.Use(store)
	e.Validator = &CustomValidator{validator: validator.New()}

	e.POST("/auth/signin", controllers.SignIn)
	auth := e.Group("/auth")
	auth.Use(utils.SessionAuth)
	auth.GET("/signout", controllers.SignOut)
	auth.POST("/signup", controllers.SignUp, Roles.Permission)

	apiV1 := e.Group("/api/v1")
	apiV1.Use(utils.SessionAuth)
	apiV1.GET("/users", controllers.GetUserList)
	apiV1.POST("/users/role", controllers.JoinUserRole)
	apiV1.POST("/roles", controllers.PostRole)
	apiV1.GET("/roles", controllers.GetRolesList)

	e.Logger.Fatal(e.Start(":1323"))
}
