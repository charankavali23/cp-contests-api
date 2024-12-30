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

// FormateCodeChefContest formats CodeChef contest to unified format
func FormateCodeChefContest(contest models.CodeChefContestDetails) (models.Contest, models.ApiError) {
	duration, err := strconv.Atoi(contest.ContestDuration)
	if err != nil {
		return models.Contest{}, utils.NewApiError("Error converting contest duration", err.Error(), http.StatusInternalServerError)
	}
	url := viper.GetString("code_chef.site_url") + contest.ContestCode
	return models.Contest{
		Platform: "CodeChef",
		Id: contest.ContestCode,
		Name: contest.ContestName,
		URL: url,
		StartTime: contest.ContestStartDateIso,
		EndTime: contest.ContestEndDateIso,
		Duration: duration,
	},
	models.ApiError{}
}

// ProcessCodeChefRawData processes raw data from CodeChef
func ProcessCodeChefRawData() models.ApiError {
	log.Println("Processing CodeChef contests")
	codeChefData = models.ServiceContests{}
	for _, contest := range codeChefRawData.PresentContests {
		formatedContest, apiError := FormateCodeChefContest(contest)
		if apiError != (models.ApiError{}) {
			return apiError
		}
		codeChefData.OngoingContests = append(codeChefData.OngoingContests, formatedContest)
	}
	for _, contest := range codeChefRawData.FutureContests {
		formatedContest, apiError := FormateCodeChefContest(contest)
		if apiError != (models.ApiError{}) {
			return apiError
		}
		codeChefData.UpcomingContests = append(codeChefData.UpcomingContests, formatedContest)
	}
	for _, contest := range codeChefRawData.PastContests {
		formatedContest, apiError := FormateCodeChefContest(contest)
		if apiError != (models.ApiError{}) {
			return apiError
		}
		codeChefData.CompletedContests = append(codeChefData.CompletedContests, formatedContest)
	}
	return models.ApiError{}
}

// CodeChefContests fetches CodeChef contests
func CodeChefContests(current_datetime time.Time) (models.ServiceContests, models.ApiError) {
	log.Println("CodeChef contests")
	if codeChefLoadDateTime.IsZero() || current_datetime.Sub(codeChefLoadDateTime).Hours() >= 12 {
		codeChefLoadDateTime = time.Now()
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
		if err := ProcessCodeChefRawData(); err != (models.ApiError{}) {
			return models.ServiceContests{}, err
		}
	}
	arr := []int{1, 2, 3}
	log.Println(arr[10])
	return codeChefData, models.ApiError{}
}