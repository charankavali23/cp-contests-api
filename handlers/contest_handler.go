package handlers

import (
	"log"

	"github.com/charankavali23/cp-contests-api/services"
	"github.com/charankavali23/cp-contests-api/utils"

	"github.com/gin-gonic/gin"
)

// GetAllContests handles the request to fetch all contests
func GetAllContests(c *gin.Context) {
	log.Println("Get all contests")
	defer ErrorHandler(c)
	resp, apiError := services.FetchAllContests()
	utils.UpdateResponse(c, resp, apiError)
}

// GetContests handles the request to fetch contests
func GetContests(c *gin.Context) {
	log.Println("Get contests")
	defer ErrorHandler(c)
	resp, apiError := services.FetchContests(c)
	utils.UpdateResponse(c, resp, apiError)
}