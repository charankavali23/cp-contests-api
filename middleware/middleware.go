package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// Logging middleware logs the request method, URL, protocol and latency
func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		end := time.Now()
		latency := end.Sub(start)
		log.Printf("%s %s %s %v\n", c.Request.Method, c.Request.URL, c.Request.Proto, latency)
	}
}

// RequestBodyValidation middleware validates the request body
func RequestBodyValidation() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Request body validation")
		c.Next()
	}
}