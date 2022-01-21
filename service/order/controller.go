package order

import (
	"evermos-assessment-be/helper"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type orderController struct {
	service Service
}

func NewController(service Service) *orderController {
	return &orderController{service: service}
}

func (controller *orderController) Route(app *gin.Engine, handler gin.HandlerFunc) {
	route := app.Group("api/orders")
	route.GET("/", handler, controller.Get)
	route.GET("/:id", handler, controller.FindById)
	route.POST("/", handler, controller.Create)
}

func (controller *orderController) Get(c *gin.Context) {
	txHandle := c.MustGet("db_trx").(*gorm.DB)

	orders := controller.service.WithTrx(txHandle).FindAll()

	c.JSON(http.StatusOK, helper.ApiResponse("List Order", http.StatusOK, "success", OrdersFormat(orders)))
	return
}

func (controller *orderController) Create(c *gin.Context) {
	var input CreateOrderRequest

	txHandle := c.MustGet("db_trx").(*gorm.DB)

	err := c.ShouldBindJSON(&input)
	helper.PanicIfError(err)

	order, err := controller.service.WithTrx(txHandle).Save(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiResponse(err.Error(), http.StatusBadRequest, "failed", map[string]interface{}{}))
		return
	}
	c.JSON(http.StatusOK, helper.ApiResponse("Create Order Successfully", http.StatusOK, "success", OrderFormat(order)))
	return
}

func (controller *orderController) FindById(c *gin.Context) {
	var inputParam DetailOrderRequest
	err := c.ShouldBindUri(&inputParam)
	helper.PanicIfError(err)

	txHandle := c.MustGet("db_trx").(*gorm.DB)

	order, err := controller.service.WithTrx(txHandle).FindById(inputParam.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiResponse(err.Error(), http.StatusBadRequest, "failed", map[string]interface{}{}))
		return
	}

	c.JSON(http.StatusOK, helper.ApiResponse("Detail Order", http.StatusOK, "success", OrderFormat(order)))
	return
}
