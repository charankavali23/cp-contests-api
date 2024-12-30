package handlers

import (
	"log"
	"fmt"
	"net/http"

	"github.com/charankavali23/cp-contests-api/utils"
	"github.com/charankavali23/cp-contests-api/models"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context) {
	log.Println("Error handler")
	r := recover()
	if r != nil {
		log.Println("Recovered from panic: ", r)
		utils.UpdateResponse(c, models.ResponseJSON{}, utils.NewApiError("Internal server error", fmt.Sprint(r), http.StatusInternalServerError))
	}
}