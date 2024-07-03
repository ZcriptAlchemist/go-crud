package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/suhailmshaik/go-crud/config"
	"github.com/suhailmshaik/go-crud/routes"
)

// execution starts at main function - this application can perform CRUD operations on a user entity - configure your postgres database in the config package and run the application
func main() {
	// initializing Gin router or a Web Server // gin.default
	router := gin.New()
	// connecting to the database
	config.Connect()
	// passing router or web server into user defined package of routes
	routes.UserRoute(router)
	// running the server on port 5000
	fmt.Println("welcome to go-crud! \nuse Postman or Thunder Client to test the API's")
	router.Run(":5000")
}