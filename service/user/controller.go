package user

import (
	"evermos-assessment-be/helper"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type userController struct {
	service Service
}

func NewController(service Service) *userController {
	return &userController{service: service}
}

func (controller *userController) Route(app *gin.Engine) {
	route := app.Group("api/users")
	route.GET("/", controller.Get)
	route.POST("/", controller.Save)
	route.GET("/:id", controller.findById)
}

func (controller *userController) Get(c *gin.Context) {
	users := controller.service.FindAll()

	c.JSON(http.StatusOK, helper.ApiResponse("List User", http.StatusOK, "success", UsersFormat(users)))
	return
}

func (controller *userController) Save(c *gin.Context) {
	var input CreateUserRequest
	err := c.ShouldBindJSON(&input)
	helper.PanicIfError(err)

	user := User{
		Name:      input.Name,
		Email:     input.Email,
		Password:  input.Password,
		Phone:     input.Phone,
		Address:   input.Address,
		City:      input.City,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	user = controller.service.Save(user)

	c.JSON(http.StatusOK, helper.ApiResponse("Create User Successfully", http.StatusOK, "success", UserFormat(user)))
	return
}

func (controller *userController) findById(c *gin.Context) {
	var inputParam DetailUserRequest
	err := c.ShouldBindUri(&inputParam)
	helper.PanicIfError(err)

	user, err := controller.service.FindById(inputParam.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiResponse(err.Error(), http.StatusBadRequest, "failed", map[string]interface{}{}))
		return
	}

	c.JSON(http.StatusOK, helper.ApiResponse("Detail User", http.StatusOK, "success", UserFormat(user)))
	return
}
