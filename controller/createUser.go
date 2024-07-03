package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/suhailmshaik/go-crud/models"
	"github.com/suhailmshaik/go-crud/service"
)

// CreateUser: This function creates a new user in the database. It binds the request body to a User struct, creates a new user, and returns the created user as a JSON response.

func CreateUser (context *gin.Context) {
	var user models.User
	context.BindJSON(&user)
	// passes the address of user to the GenerateKey function in service layer
	service.GenerateKey(&user)
	context.JSON(200, &user)
}