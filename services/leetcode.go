package services

import (
	"log"
	"time"

	"github.com/charankavali23/cp-contests-api/models"
	"github.com/charankavali23/cp-contests-api/utils"

	"github.com/spf13/viper"
)

func FormateLeetcodeContest(contest models.LeetcodeContestDetails) (models.Contest, models.ApiError) {
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

func GetLeetcodeContests() (models.ServiceContests, models.ApiError) {
	log.Println("Fetching Leetcode contests")
	currentDatetime := time.Now()
	if currentDatetime.IsZero() || currentDatetime.Sub(leetcodeLoadDateTime).Hours() >= 12 {
		resp, apiError := utils.FetchAPIResponse(viper.GetString("leetcode.api_url"))
		if apiError != (models.ApiError{}) {
			log.Println("Error fetching Leetcode contests")
			return models.ServiceContests{}, apiError
		}
		defer resp.Body.Close()
		if apiError := utils.GetJsonBody(resp.Body, &leetcodeRawData); apiError != (models.ApiError{}) {
			return models.ServiceContests{}, apiError
		}
		rawData := [][]models.LeetcodeContestDetails{
			leetcodeRawData.Data.AllContests,
		}
		if apiError := utils.ProcessRawData(rawData, &leetcodeData, FormateLeetcodeContest); apiError != (models.ApiError{}) {
			return models.ServiceContests{}, apiError
		}
		leetcodeLoadDateTime = currentDatetime
	}
	return leetcodeData, models.ApiError{}
}