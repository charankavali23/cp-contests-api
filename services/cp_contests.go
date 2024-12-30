package services

import (
	"log"
	"time"

	"github.com/charankavali23/cp-contests-api/models"

	"github.com/gin-gonic/gin"
)

var codeChefLoadDateTime time.Time
var codeChefRawData models.CodeChefContests
var codeChefData models.ServiceContests

func FetchAllContests() (models.ResponseJSON, models.ApiError) {
	log.Println("Fetch all contests")
	currentDatetime := time.Now()
	contests, apiError := CodeChefContests(currentDatetime)
	if apiError != (models.ApiError{}) {
		return models.ResponseJSON{}, apiError
	}
	resp := models.ResponseJSON{
		Status:  "success",
		ServiceContests: contests,
	}
	return resp, models.ApiError{}
}

func FetchContests(c *gin.Context) (models.ResponseJSON, models.ApiError) {
	log.Println("Fetch contests")
	currentDatetime := time.Now()
	contests, apiError := CodeChefContests(currentDatetime)
	if apiError != (models.ApiError{}) {
		return models.ResponseJSON{}, apiError
	}
	resp := models.ResponseJSON{
		Status:  "success",
		ServiceContests: contests,
	}
	return resp, models.ApiError{}
}