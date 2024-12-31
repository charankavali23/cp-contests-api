package services

import (
	// "encoding/json"
	// "io"
	"log"
	"net/http"
	"time"
	"strconv"

	"github.com/charankavali23/cp-contests-api/models"
	"github.com/charankavali23/cp-contests-api/utils"

	"github.com/spf13/viper"
)

// FormateCodeChefContest formats CodeChef contest to unified format
func FormateCodeChefContest(contest models.CodeChefContestDetails) (models.Contest, models.ApiError) {
	duration, err := strconv.Atoi(contest.ContestDuration)
	if err != nil {
		return models.Contest{}, utils.NewApiError("Error converting contest duration", err.Error(), http.StatusInternalServerError)
	}
	return models.Contest{
		Platform: "codechef",
		Id: contest.ContestCode,
		Name: contest.ContestName,
		URL: viper.GetString("code_chef.site_url") + contest.ContestCode,
		StartTime: contest.ContestStartDateIso,
		EndTime: contest.ContestEndDateIso,
		Duration: duration,
	},
	models.ApiError{}
}

// CodeChefContests fetches CodeChef contests
func GetCodeChefContests(currentDatetime time.Time) (models.ServiceContests, models.ApiError) {
	log.Println("CodeChef contests")
	if codeChefLoadDateTime.IsZero() || currentDatetime.Sub(codeChefLoadDateTime).Hours() >= 12 {
		resp, apiError := utils.FetchAPIResponse(viper.GetString("code_chef.api_url"))
		if apiError != (models.ApiError{}) {
			log.Println("Error fetching CodeChef contests")
			return models.ServiceContests{}, apiError
		}
		if apiError := utils.GetJsonBody(resp.Body, &codeChefRawData); apiError != (models.ApiError{}) {
			return models.ServiceContests{}, apiError
		}
		rawData := [][]models.CodeChefContestDetails{
			codeChefRawData.PresentContests,
			codeChefRawData.FutureContests,
			codeChefRawData.PastContests,
		}
		if apiError := utils.ProcessRawData(rawData, &codeChefData, FormateCodeChefContest); apiError != (models.ApiError{}) {
			return models.ServiceContests{}, apiError
		}
		codeChefLoadDateTime = currentDatetime
	}
	return codeChefData, models.ApiError{}
}