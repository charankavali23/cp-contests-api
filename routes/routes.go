package routes

import (
	"log"

	"github.com/charankavali23/cp-contests-api/handlers"

	"github.com/gin-gonic/gin"
)

// InitRouter initializes the API routes
func InitRouter(router *gin.Engine) {
	log.Println("Initializing routes")
	api := router.Group("/cp-contests") 
	{
		api.GET("/", handlers.GetAllContests)
		api.POST("/", handlers.GetContests)
	}
	log.Println("Routes initialized")
}