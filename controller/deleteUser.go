package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/suhailmshaik/go-crud/DAO"

	// "github.com/suhailmshaik/go-crud/config"
	"github.com/suhailmshaik/go-crud/models"
)

// DeleteUser: This function deletes a user from the database based on the user's ID. It retrieves the user by ID, deletes the user, and returns the deleted user as a JSON response.

func DeleteUser (context *gin.Context) {
	var user models.User
	DAO.DeleteUser(&user, context)
}