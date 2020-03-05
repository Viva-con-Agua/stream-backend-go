package main

import (
    "./controllers"
    "./utils"
    "github.com/Viva-con-Agua/echo-pool/pool"
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

    // intial loading function
    utils.LoadConfig()
    utils.ConnectDatabase()
    store := pool.RedisSession("172.2.150.2:6379")

    //create echo server
    e := echo.New()
    e.Use(store)
    e.Validator = &CustomValidator{validator: validator.New()}

    // TODO: Listen for user creation on nats

    apiV1 := e.Group("/drops-backend/api/v1")
    apiV1.Use(pool.SessionAuth)

    apiV1.GET("/supporter", controllers.GetSupporterList)

    apiV1.GET("/supporter/:id", controllers.GetSupporter)
    apiV1.PUT("/supporter", controllers.UpdateSupporter)
    apiV1.DELETE("/supporter", controllers.DeleteSupporter)
    // apiV1.POST("/supporter/role", controllers.JoinSupporterRole)

    apiV1.GET("/crew/:id", controllers.GetCrew)

    apiV1.GET("/crew", controllers.GetCrewList)
    apiV1.PUT("/crew", controllers.UpdateCrew)
    apiV1.DELETE("/crew", controllers.DeleteCrew)

    e.Logger.Fatal(e.Start(":1323"))
}
