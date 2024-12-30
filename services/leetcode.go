package services

import (
	"encoding/json"
	"net/http"
	"io"
	"log"
	"time"

	"github.com/charankavali23/cp-contests-api/models"
	"github.com/charankavali23/cp-contests-api/utils"

	"github.com/spf13/viper"
)

func formateLeetcodeContest(contest models.LeetcodeContestDetails) (models.Contest, models.ApiError) {
	return models.Contest{
		Platform: "leetcode",
		Id: contest.TitleSlug,
		Name: contest.Title,
		URL: viper.GetString("leetcode.contest_url") + contest.TitleSlug,
		StartTime: time.Unix(int64(contest.StartTime), 0).Format(time.RFC3339),
		EndTime: time.Unix(int64(contest.StartTime+contest.Duration), 0).Format(time.RFC3339),
		Duration: contest.Duration/60,
	},
	models.ApiError{}
}

func processLeetcodeRawData() models.ApiError {
	log.Println("Processing leetcode contests")
	leetcodeData = models.ServiceContests{}
	for _, contest := range leetcodeRawData.Data.AllContests {
		formatedContest, apiError := formateLeetcodeContest(contest)
		if apiError != (models.ApiError{}) {
			return apiError
		}
		leetcodeData.AllContests = append(leetcodeData.AllContests, formatedContest)
	}
	return models.ApiError{}
}

func GetLeetcodeContests(currentDatetime time.Time) (models.ServiceContests, models.ApiError) {
	log.Println("Fetching Leetcode contests")
	if currentDatetime.IsZero() || currentDatetime.Sub(leetcodeLoadDateTime).Hours() >= 12 {
		resp, apiError := utils.FetchAPIResponse(viper.GetString("leetcode.api_url"))
		if apiError != (models.ApiError{}) {
			log.Println("Error fetching Leetcode contests")
			return models.ServiceContests{}, apiError
		}
		resp_body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("Error reading leetcode contests response body")
			return models.ServiceContests{}, utils.NewApiError("Error reading leetcode contests response body", err.Error(), http.StatusInternalServerError)
		}	
		if err := json.Unmarshal(resp_body, &leetcodeRawData); err != nil {
			log.Println("Error unmarshalling leetcode contests response body")
			return models.ServiceContests{}, utils.NewApiError("Error unmarshalling leetcode contests response body", err.Error(), http.StatusInternalServerError)
		}
		if err := processLeetcodeRawData(); err != (models.ApiError{}) {
			return models.ServiceContests{}, err
		}
		leetcodeLoadDateTime = currentDatetime
	}
	return leetcodeData, models.ApiError{}
}