package main

import (
	"log"

	"github.com/charankavali23/cp-contests-api/config"
	"github.com/charankavali23/cp-contests-api/middleware"
	"github.com/charankavali23/cp-contests-api/routes"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	log.Println("Competitive Programming Contest API")

	// load configuration
	config.LoadConfig()

	log.Println(viper.Get("time_zone"))

	// set up the router
	router := gin.Default()

	// set up middleware
	router.Use(middleware.Logging(), middleware.RequestBodyValidation())

	// initialize routes
	routes.InitRouter(router)

	// start the server
	router.Run(":8000")
}