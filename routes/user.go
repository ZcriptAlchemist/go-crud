package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/suhailmshaik/go-crud/controller"
)

// *gin.Engine is a pointer to the instance of the Gin router aka a web server

func UserRoute(router *gin.Engine) {
	// router.GET("/", controller.UserController)
	router.GET("/users/get", controller.GetUsers)
	router.POST("/users/create", controller.CreateUser)
	router.DELETE("/users/delete/:id", controller.DeleteUser)
	router.PUT("/users/update/:id", controller.UpdateUser)
}