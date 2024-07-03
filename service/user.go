package service

import (
	"github.com/suhailmshaik/go-crud/DAO"
	"github.com/suhailmshaik/go-crud/models"
	"github.com/suhailmshaik/go-crud/utils"
)

// AddKey: This function generates a unique key for the user and adds it to the user object.
func AddKey(user *models.User) {
	// AddKey function implementation
	generatedKey := utils.GenerateKey()
	user.Key = generatedKey

	// 	The DAO package is imported to access the AddUser function, which adds the user to the database.
	DAO.AddUser(user)
}