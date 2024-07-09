package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/suhailmshaik/go-crud/DAO"
	"github.com/suhailmshaik/go-crud/models"
	"github.com/suhailmshaik/go-crud/service"
)

// CreateUser: This function creates a new user in the database. It binds the request body to a User struct, creates a new user, and returns the created user as a JSON response.
func CreateUser (context *gin.Context) {
	var user models.User
	context.BindJSON(&user)
	// passes the address of user to the AddKey function in service layer
	service.AddKey(&user)
	context.JSON(200, &user)
}

// GetUsers: This function retrieves all users from the database and returns them as a JSON response.
func GetUsers (context *gin.Context) {
	var users []models.User
	// Pasing the address of the users slice to the ReadUsers function in the DAO layer.
	DAO.ReadUsers(&users)
	context.JSON(200, &users)
}

// UpdateUser: This function updates a user in the database based on the user's ID. It retrieves the user by ID, binds the request body to the user struct, saves the updated user, and returns the updated user as a JSON response.
func UpdateUser (context *gin.Context) {
	var user models.User
    context.BindJSON(&user)
	DAO.UpdateUser(&user, context)
	context.JSON(200, &user)
}

// DeleteUser: This function deletes a user from the database based on the user's ID. It retrieves the user by ID, deletes the user, and returns the deleted user as a JSON response.
func DeleteUser (context *gin.Context) {
	var user models.User
	DAO.DeleteUser(&user, context)
	context.JSON(200, &user)
}