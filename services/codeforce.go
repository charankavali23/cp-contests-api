package services

import (
	"log"
	"strconv"
	"time"

	"github.com/charankavali23/cp-contests-api/models"
	"github.com/charankavali23/cp-contests-api/utils"

	"github.com/spf13/viper"
)

func FormateCodeforcesContest(contest models.CodeforcesContestDetails) (models.Contest, models.ApiError) {
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

func GetCodeforcesContests(currentDatetime time.Time) (models.ServiceContests, models.ApiError) {
	log.Println("Fetching Codeforces contests")
	if currentDatetime.IsZero() || currentDatetime.Sub(codeforcesLoadDateTime).Hours() >= 12 {
		resp, apiError := utils.FetchAPIResponse(viper.GetString("codeforces.api_url"))
		if apiError != (models.ApiError{}) {
			log.Println("Error fetching Codeforces contests")
			return models.ServiceContests{}, apiError
		}
		if apiError := utils.GetJsonBody(resp.Body, &codeforcesRawData); apiError != (models.ApiError{}) {
			return models.ServiceContests{}, apiError
		}
		rawData := [][]models.CodeforcesContestDetails{
			codeforcesRawData.Result,
		}
		if apiError := utils.ProcessRawData(rawData, &codeforcesData, FormateCodeforcesContest); apiError != (models.ApiError{}) {
			return models.ServiceContests{}, apiError
		}
		codeforcesLoadDateTime = currentDatetime
	}
	return codeforcesData, models.ApiError{}
}