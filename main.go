package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/suhailmshaik/go-crud/config"
	"github.com/suhailmshaik/go-crud/routes"
)

func main() {
	fmt.Println("welcome to go-crud")
	// initializing Gin router or a Web Server // gin.default
	router := gin.New()
	// connecting to the database
	config.Connect()
	// passing router or web server into user defined package of routes
	routes.UserRoute(router)
	// running the server on port 5000
	router.Run(":5000")
}