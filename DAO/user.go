package DAO

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/suhailmshaik/go-crud/config"
	"github.com/suhailmshaik/go-crud/models"
)

// AddUser is a function that adds a new user to the database.

func AddUser (user *models.User) {
	config.DB.Create(user)
	fmt.Printf("%v successfully added as user",user.Name)
}

// ReadUsers is a function that retrieves all users from the database.

func ReadUsers (users *[]models.User) {
	config.DB.Find(&users)
	fmt.Println("All users successfully retrieved")
}

// UpdateUser is a function that updates a user in the database based on the user's ID.

func UpdateUser (updatedUser *models.User, context *gin.Context) {

	// retrieve the previous user by ID
	var user models.User
	result:=config.DB.First(&user, context.Param("id"))
	if result.Error != nil {
		fmt.Println(result.Error) // prints error if the user is not found
	}

	// updating the user with the new user data
	result = config.DB.Model(&user).Updates(updatedUser)
    if result.Error != nil {
        fmt.Println(result.Error) // prints error if the update operation fails
    }
	fmt.Println("Record updated successfully")
}

// DeleteUser is a function that deletes a user from the database based on the user's ID.

func DeleteUser (user *models.User, context *gin.Context) {
	config.DB.Where("id= ?", context.Param("id")).Delete(&user)
	fmt.Println("successfully deleted user")
}