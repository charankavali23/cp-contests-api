package main

import (
	"log"

	"github.com/charankavali23/cp-contests-api/config"
	"github.com/charankavali23/cp-contests-api/middleware"
	"github.com/charankavali23/cp-contests-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Competitive Programming Contest API")
	// load configuration
	config.LoadConfig()
	// set and initilizeing the router and middleware
	router := gin.Default()
	router.Use(middleware.APILoggingMiddleware(), middleware.RequestValidatorMiddleware())
	routes.InitRouter(router)
	// start the server
	router.Run(":8000")
}