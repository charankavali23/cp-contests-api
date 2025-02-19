package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"reflect"

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

// All data in jsonBodyRC will be read and parses the JSON-encoded data and stores the result in the value pointed to by v. If v is nil or not a pointer
func GetJsonBody(jsonBodyRC io.ReadCloser, v any) models.ApiError {
	jsonBodyBytes, err := io.ReadAll(jsonBodyRC)
	if err != nil {
		log.Println("Error reading contests response body")
		return NewApiError("Error reading contests response body", err.Error(), http.StatusInternalServerError)
	}
	if err := json.Unmarshal(jsonBodyBytes, &v); err != nil {
		log.Println("Error unmarshalling contests response body")
		return NewApiError("Error unmarshalling contests response body", err.Error(), http.StatusInternalServerError)
	}
	return models.ApiError{}
}

func ProcessRawData[serviceContestDetails any](rawData [][]serviceContestDetails, processedData *models.ServiceContests, formatContest func(serviceContestDetails) (models.Contest, models.ApiError)) models.ApiError {
	for _, contestsArray := range rawData {
		for _, rawContest := range contestsArray {
			formatedContest, apiError := formatContest(rawContest)
			if apiError != (models.ApiError{}) {
				return apiError
			}
			processedData.AllContests = append(processedData.AllContests, formatedContest)
		}
	}
	return models.ApiError{}
}

func IsAvailable[dataType any](object dataType, array []dataType) bool {
	for _, element := range array {
		if reflect.DeepEqual(element, object) {
			return true
		}
	}
	return false
}

func MapToStruct[structType any](data map[string]interface{}, object *structType) models.ApiError {
	jsonDate, err := json.Marshal(data)
	if err != nil {
		return NewApiError("Error marshalling data", err.Error(), http.StatusInternalServerError)
	}
	if err := json.Unmarshal(jsonDate, &object); err != nil {
		return NewApiError("Error unmarshalling data", err.Error(), http.StatusInternalServerError)
	}
	return models.ApiError{}
}