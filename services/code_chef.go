package services

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
	"strconv"

	"github.com/charankavali23/cp-contests-api/models"
	"github.com/charankavali23/cp-contests-api/utils"

	"github.com/spf13/viper"
)

// formateCodeChefContest formats CodeChef contest to unified format
func formateCodeChefContest(contest models.CodeChefContestDetails) (models.Contest, models.ApiError) {
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

func addContestInCodeChefData(codeChefRawContests []models.CodeChefContestDetails) (models.ApiError) {
	for _, contest := range codeChefRawContests {
		formatedContest, apiError := formateCodeChefContest(contest)
		if apiError != (models.ApiError{}) {
			return apiError
		}
		codeChefData.AllContests = append(codeChefData.AllContests, formatedContest)
	}
	return models.ApiError{}
}

// processCodeChefRawData processes raw data from CodeChef
func processCodeChefRawData() models.ApiError {
	log.Println("Processing CodeChef contests")
	codeChefData = models.ServiceContests{}
	if apiError := addContestInCodeChefData(codeChefRawData.PresentContests); apiError != (models.ApiError{}) {
		return apiError
	}
	if apiError := addContestInCodeChefData(codeChefRawData.FutureContests); apiError != (models.ApiError{}) {
		return apiError
	}
	if apiError := addContestInCodeChefData(codeChefRawData.PastContests); apiError != (models.ApiError{}) {
		return apiError
	}
	return models.ApiError{}
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
		resp_body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("Error reading CodeChef contests response body")
			return models.ServiceContests{}, utils.NewApiError("Error reading CodeChef contests response body", err.Error(), http.StatusInternalServerError)
		}
		if err := json.Unmarshal(resp_body, &codeChefRawData); err != nil {
			log.Println("Error unmarshalling CodeChef contests response body")
			return models.ServiceContests{}, utils.NewApiError("Error unmarshalling CodeChef contests response body", err.Error(), http.StatusInternalServerError)
		}
		if apiError := processCodeChefRawData(); apiError != (models.ApiError{}) {
			return models.ServiceContests{}, apiError
		}
		codeChefLoadDateTime = currentDatetime
	}
	return codeChefData, models.ApiError{}
}