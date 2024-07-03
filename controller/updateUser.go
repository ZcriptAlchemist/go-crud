package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/suhailmshaik/go-crud/DAO"
	"github.com/suhailmshaik/go-crud/models"
)

// UpdateUser: This function updates a user in the database based on the user's ID. It retrieves the user by ID, binds the request body to the user struct, saves the updated user, and returns the updated user as a JSON response.

func UpdateUser (context *gin.Context) {
	var user models.User
    context.BindJSON(&user)
	DAO.UpdateUser(&user, context)
	context.JSON(200, &user)
}