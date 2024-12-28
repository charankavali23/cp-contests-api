package handlers

import (
	"log"
	"net/http"

	"github.com/charankavali23/cp-contests-api/services"

	"github.com/gin-gonic/gin"
)

// GetAllContests handles the request to fetch all contests
func GetAllContests(c *gin.Context) {
	log.Println("Get all contests")
	response, err := services.FetchAllContests()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{ "error": err.Error() })
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetContests handles the request to fetch contests
func GetContests(c *gin.Context) {
	log.Println("Get contests")
	response, err := services.FetchContests(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{ "error": err.Error() })
		return
	}

	c.JSON(http.StatusOK, response)
}