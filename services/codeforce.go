package services

import (
	"encoding/json"
	"net/http"
	"io"
	"log"
	"strconv"
	"time"

	"github.com/charankavali23/cp-contests-api/models"
	"github.com/charankavali23/cp-contests-api/utils"

	"github.com/spf13/viper"
)

func formateCodeforcesContest(contest models.CodeforcesContestDetails) (models.Contest, models.ApiError) {
	var url string
	if contest.Phase == "CODING" || contest.Phase == "BEFORE" {
		url = viper.GetString("codeforces.contest_reg_url") + strconv.Itoa(contest.Id)
	} else if contest.Phase == "FINISHED" {
		url = viper.GetString("codeforces.contest_url") + strconv.Itoa(contest.Id)
	}
	return models.Contest{
		Platform: "codeforces",
		Id: strconv.Itoa(contest.Id),
		Name: contest.Name,
		URL: url,
		StartTime: time.Unix(int64(contest.StartTimeSeconds), 0).Format(time.RFC3339),
		EndTime: time.Unix(int64(contest.StartTimeSeconds+contest.DurationSeconds), 0).Format(time.RFC3339),
		Duration: contest.DurationSeconds/60,
	},
	models.ApiError{}
}

func processCodeforcesRawData() models.ApiError {
	log.Println("Processing Codeforces contests")
	codeforcesData = models.ServiceContests{}
	for _, contest := range codeforcesRawData.Result {
		formatedContest, apiError := formateCodeforcesContest(contest)
		if apiError != (models.ApiError{}) {
			return apiError
		}
		codeforcesData.AllContests = append(codeforcesData.AllContests, formatedContest)
	}
	return models.ApiError{}
}

func GetCodeforceContests(currentDatetime time.Time) (models.ServiceContests, models.ApiError) {
	log.Println("Fetching Codeforces contests")
	if currentDatetime.IsZero() || currentDatetime.Sub(codeforcesLoadDateTime).Hours() >= 12 {
		resp, apiError := utils.FetchAPIResponse(viper.GetString("codeforces.api_url"))
		if apiError != (models.ApiError{}) {
			log.Println("Error fetching Codeforces contests")
			return models.ServiceContests{}, apiError
		}
		resp_body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("Error reading Codeforces contests response body")
			return models.ServiceContests{}, utils.NewApiError("Error reading Codeforces contests response body", err.Error(), http.StatusInternalServerError)
		}
		if err := json.Unmarshal(resp_body, &codeforcesRawData); err != nil {
			log.Println("Error unmarshalling Codeforces contests response body")
			return models.ServiceContests{}, utils.NewApiError("Error unmarshalling Codeforces contests response body", err.Error(), http.StatusInternalServerError)
		}
		if err := processCodeforcesRawData(); err != (models.ApiError{}) {
			return models.ServiceContests{}, err
		}
		codeforcesLoadDateTime = currentDatetime
	}
	return codeforcesData, models.ApiError{}
}