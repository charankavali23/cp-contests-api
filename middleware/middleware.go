package middleware

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/charankavali23/cp-contests-api/models"
	"github.com/charankavali23/cp-contests-api/utils"

	"github.com/gin-gonic/gin"
	"github.com/xeipuuv/gojsonschema"
)

var requestBodySchema = map[string]interface{}{
	"$schema": "http://json-schema.org/draft-07/schema#",
	"type": "object",
	"required": []string{"usage", "platforms", "phases"},
	"properties": map[string]interface{}{
		"usage": map[string]interface{}{
			"type": "string",
			"description": "Purpose of the request, e.g., 'testing', 'production'.",
		},
		"platforms": map[string]interface{}{
			"type": "array",
			"items": map[string]interface{}{
				"type": "string",
				"enum": []string{"codechef", "codeforces", "leetcode",},
			},
			"uniqueItems": true,
			"minItems": 1,
			"description": "List of platforms to fetch contests from.",
		},
		"phases": map[string]interface{}{
			"type": "array",
			"items": map[string]interface{}{
				"type": "string",
				"enum": []string{"ongoing", "upcoming", "completed",},
			},
			"uniqueItems": true,
			"minItems": 1,
			"description": "Phases of contests to include.",
		},
		"from_date_time": map[string]interface{}{
			"type": []string{"string", "null",},
			"format": "date-time",
			"pattern": "^\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}(?:\\+\\d{2}:\\d{2}|-\\d{2}:\\d{2})$",
			"description": "Start of the date-time range (ISO 8601 format).",
		},
		"to_date_time": map[string]interface{}{
			"type": []string{"string", "null",},
			"format": "date-time",
			"pattern": "^\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}(?:\\+\\d{2}:\\d{2}|-\\d{2}:\\d{2})$",
			"description": "End of the date-time range (ISO 8601 format).",
		},
		"duration_range": map[string]interface{}{
			"type": "object",
			"required": []string{},
			"properties": map[string]interface{}{
				"min_minutes": map[string]interface{}{
					"type": []string{"integer", "null",},
					"minimum": 0,
					"description": "Minimum duration of contests in minutes.",
				},
				"max_minutes": map[string]interface{}{
					"type": []string{"integer", "null",},
					"minimum": 0,
					"description": "Maximum duration of contests in minutes.",
				},
			},
			"additionalProperties": false,
			"description": "Range for contest durations in minutes.",
		},
		"sort_order": map[string]interface{}{
			"type": "object",
			"required": []string{},
			"properties": map[string]interface{}{
				"by_duration": map[string]interface{}{
					"type": []string{"string", "null",},
					"enum": []interface{}{"asc", "desc", nil,},
					"description": "Sort order for contest durations.",
				},
				"by_start_time": map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"ongoing": map[string]interface{}{
							"type": []string{"string", "null",},
							"enum": []interface{}{"asc", "desc", nil,},
						},
						"upcoming": map[string]interface{}{
							"type": []string{"string", "null",},
							"enum": []interface{}{"asc", "desc", nil,},
						},
						"completed": map[string]interface{}{
							"type": []string{"string", "null",},
							"enum": []interface{}{"asc", "desc", nil,},
						},
					},
					"additionalProperties": false,
					"description": "Sort order for each phase.",
				},
			},
			"additionalProperties": false,
			"description": "Sort order for contests.",
		},
	},
	"additionalProperties": false,
}

// APILoggingMiddleware logs the request method, URL, protocol and latency
func APILoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		end := time.Now()
		latency := end.Sub(start)
		log.Printf("%s %s %s %v\n", c.Request.Method, c.Request.URL, c.Request.Proto, latency)
	}
}

func validateRequestBody(requestBody map[string]interface{}) models.ApiError {
	jsonSchemaLoader := gojsonschema.NewGoLoader(requestBodySchema)
	jsonBodyLoader := gojsonschema.NewGoLoader(requestBody)
	result, err := gojsonschema.Validate(jsonSchemaLoader, jsonBodyLoader)
	if err != nil {
		return utils.NewApiError("Error validating request body", err.Error(), http.StatusBadRequest)
	}
	if !result.Valid() {
		errorMsg := ""
		for _, desc := range result.Errors() {
			errorMsg += desc.String() + "\n"
		}
		return utils.NewApiError("Invalid request body", errorMsg, http.StatusBadRequest)
	}
	return models.ApiError{}
}

// RequestValidatorMiddleware validates the request body
func RequestValidatorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Request body validation")
		if c.Request.Method == "POST" {
			log.Println("POST request")
			var validationRequestBody map[string]interface{}
			jsonBodyBytes, err := io.ReadAll(c.Request.Body)
			if err != nil {
				log.Println("Error reading contests request body")
				utils.UpdateResponse(c, models.ResponseBody{}, utils.NewApiError("Error reading contests request body", err.Error(), http.StatusInternalServerError))
				c.Abort()
				return
			}
			if err := json.Unmarshal(jsonBodyBytes, &validationRequestBody); err != nil {
				log.Println("Error unmarshalling contests request body")
				utils.UpdateResponse(c, models.ResponseBody{}, utils.NewApiError("Error unmarshalling contests request body", err.Error(), http.StatusInternalServerError))
				c.Abort()
				return
			}
			log.Printf("Validation Request body: %+v", validationRequestBody)
			if apiError := validateRequestBody(validationRequestBody); apiError != (models.ApiError{}) {
				utils.UpdateResponse(c, models.ResponseBody{}, apiError)
				c.Abort()
				return
			}
			log.Println("Request body validated")
			var validatedRequestBody models.RequestBody
			if apiError := utils.MapToStruct(validationRequestBody, &validatedRequestBody); apiError != (models.ApiError{}) {
				utils.UpdateResponse(c, models.ResponseBody{}, apiError)
				c.Abort()
				return
			}
			log.Printf("Validated Request body: %+v", validatedRequestBody)
			c.Set("validatedRequestBody", validatedRequestBody)
		} else {
			log.Println("GET request")
		}
		c.Next()
	}
}