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

    apiV1 := e.Group("/backend/stream/api/v1")
    // TODO REENABLE AUTHENTICATION
    //apiV1.Use(pool.SessionAuth)

    /*
     * Routes for takings
     */
    // GET routes for takings
    // TODO: Combine takings get and takings count
    apiV1.GET("/takings", controllers.GetTakingList)
    apiV1.GET("/takings/count", controllers.GetTakingCount)
    // TODO: Remove id from routes
    apiV1.GET("/takings/id/:id", controllers.GetTakingById)
    apiV1.GET("/takings/:id", controllers.GetTakingById)

    // POST routes for takings
    apiV1.POST("/takings/create", controllers.CreateTaking)

    // PUT routes for takings
    apiV1.PUT("/takings/update", controllers.UpdateTaking)

    /*
     * Routes for deposits
     */
    // GET routes for takings
    // TODO: Combine deposits get and deposits count
    apiV1.GET("/deposits", controllers.GetDepositList)
    apiV1.GET("/deposits/count", controllers.GetDepositCount)
    // TODO: Implement get of single deposit
    //apiV1.GET("/depoits/:id", controllers.GetDepositById)

    // POST routes for deposits
    apiV1.POST("/deposits/create", controllers.CreateDeposit)
    apiV1.POST("/deposits/confirm", controllers.ConfirmDeposit)

    // PUT routes for deposits
    // TODO: Implement Update of deposits
    // apiV1.PUT("/deposits/update", controllers.UpdateDeposit)

    // TODO: Add Household routes

    e.Logger.Fatal(e.Start(":1323"))
}
