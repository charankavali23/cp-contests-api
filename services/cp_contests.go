package services

import (
	"log"

	"github.com/charankavali23/cp-contests-api/models"

	"github.com/gin-gonic/gin"
)

func FetchAllContests() ([]models.Contest, error) {
	log.Println("Fetch all contests")
	contests := []models.Contest{
		{
			Platform: "Codeforces",
			Id: "123",
			Name: "Codeforces Round #123",
			StartTime: "2021-07-01T12:00:00Z",
			Duration: 120,
		},
		{
			Platform: "Codeforces",
			Id: "124",
			Name: "Codeforces Round #124",
			StartTime: "2021-07-02T12:00:00Z",
			Duration: 120,
		},
	}
	return contests, nil
}

func FetchContests(c *gin.Context) ([]models.Contest, error) {
	log.Println("Fetch contests")
	contests := []models.Contest{
		{
			Platform: "Codeforces",
			Id: "125",
			Name: "Codeforces Round #125",
			StartTime: "2021-07-01T12:00:00Z",
			Duration: 120,
		},
		{
			Platform: "Codeforces",
			Id: "126",
			Name: "Codeforces Round #126",
			StartTime: "2021-07-02T12:00:00Z",
			Duration: 120,
		},
	}
	return contests, nil
}