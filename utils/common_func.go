package utils

import (
	"log"
	"net/http"

	"github.com/charankavali23/cp-contests-api/models"

	"github.com/gin-gonic/gin"
)

func UpdateResponse(c *gin.Context, resp_body models.ResponseBody, apiError models.ApiError) {
	if apiError != (models.ApiError{}) {
		errResp := models.ResponseBodyError{
			Status: "failed",
			Message: apiError.Message,
			Error: apiError.Error,
		}
		c.JSON(apiError.StatusCode, errResp)
		return
	}
	c.JSON(http.StatusOK, resp_body)
}

func FetchAPIResponse(url string) (*http.Response, models.ApiError) {
	log.Println("Fetching API response")
	resp, err := http.Get(url)
	if err != nil {
		return nil, NewApiError("Error fetching API response", err.Error(), http.StatusServiceUnavailable)
	} else if resp.StatusCode != http.StatusOK {
		return nil, NewApiError("Error fetching API response", "API response is not OK", http.StatusServiceUnavailable)
	}
	return resp, models.ApiError{}
}

func NewApiError(message string, err string, statusCode int) models.ApiError {
	apiError := models.ApiError{
		Message: message,
		Error: err,
		StatusCode: statusCode,
	}
	log.Println("API Error: ", apiError)
	return apiError
}

