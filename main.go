package main

import (
	"evermos-assessment-be/app"
	"evermos-assessment-be/exception"
	"evermos-assessment-be/helper"
	"evermos-assessment-be/middleware"
	"evermos-assessment-be/service/order"
	"evermos-assessment-be/service/product"
	"evermos-assessment-be/service/user"
	"github.com/gin-gonic/gin"
)

func main() {

	config := app.New()
	db := app.NewDB(config)

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userController := user.NewController(userService)

	productRepository := product.NewRepository(db)
	productService := product.NewService(productRepository)
	productController := product.NewController(productService)

	orderRepository := order.NewRepository(db)
	orderService := order.NewService(orderRepository, userService, productService)
	orderControler := order.NewController(orderService)

	// Setup Gin
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.CustomRecovery(exception.ErrorHandler))

	// Setup Route
	userController.Route(router)
	productController.Route(router)
	orderControler.Route(router, middleware.DBTransactionMiddleware(db))

	// Start App
	err := router.Run(":8000")
	helper.PanicIfError(err)

}
