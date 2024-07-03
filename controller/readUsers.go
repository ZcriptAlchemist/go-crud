package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/suhailmshaik/go-crud/DAO"
	"github.com/suhailmshaik/go-crud/models"
)

// GetUsers: This function retrieves all users from the database and returns them as a JSON response.

func GetUsers (context *gin.Context) {
	var users []models.User
	// Pasing the address of the users slice to the ReadUsers function in the DAO layer.
	DAO.ReadUsers(&users)
	context.JSON(200, &users)
}