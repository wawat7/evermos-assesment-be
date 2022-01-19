package main

import (
	"evermos-assessment-be/app"
	"evermos-assessment-be/exception"
	"evermos-assessment-be/helper"
	"evermos-assessment-be/service/user"
	"github.com/gin-gonic/gin"
)

func main() {

	config := app.New()
	db := app.NewDB(config)

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userController := user.NewController(userService)

	// Setup Gin
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.CustomRecovery(exception.ErrorHandler))

	// Setup Route
	userController.Route(router)

	// Start App
	err := router.Run(":8000")
	helper.PanicIfError(err)

}
