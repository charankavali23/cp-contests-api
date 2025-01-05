package services

import (
	"log"
	"sort"
	"time"

	"github.com/charankavali23/cp-contests-api/models"

	"github.com/gin-gonic/gin"
)

var codeChefLoadDateTime time.Time
var codeChefRawData models.CodeChefContests
var codeChefData models.ServiceContests

var codeforcesLoadDateTime time.Time
var codeforcesRawData models.CodeforcesContests
var codeforcesData models.ServiceContests

var leetcodeLoadDateTime time.Time
var leetcodeRawData models.LeetcodeContests
var leetcodeData models.ServiceContests

func getResponseBody(allContests []models.Contest) (models.ResponseBody) {
	var resp_body models.ResponseBody
	currentDatetime := time.Now().Format(time.RFC3339)
	for _, contest := range allContests {
		if contest.StartTime <= currentDatetime && currentDatetime <= contest.EndTime {
			resp_body.OngoingContests = append(resp_body.OngoingContests, contest)
		} else if currentDatetime < contest.StartTime {
			resp_body.UpcomingContests = append(resp_body.UpcomingContests, contest)
		} else if currentDatetime > contest.EndTime {
			resp_body.CompletedContests = append(resp_body.CompletedContests, contest)
		}
	}
	// Sort OngoingContests by StartTime and Duration
    sort.Slice(resp_body.OngoingContests, func(i, j int) bool {
        if resp_body.OngoingContests[i].StartTime == resp_body.OngoingContests[j].StartTime {
            return resp_body.OngoingContests[i].Duration < resp_body.OngoingContests[j].Duration
        }
        return resp_body.OngoingContests[i].StartTime < resp_body.OngoingContests[j].StartTime
    })
    // Sort UpcomingContests by StartTime and Duration
    sort.Slice(resp_body.UpcomingContests, func(i, j int) bool {
        if resp_body.UpcomingContests[i].StartTime == resp_body.UpcomingContests[j].StartTime {
            return resp_body.UpcomingContests[i].Duration < resp_body.UpcomingContests[j].Duration
        }
        return resp_body.UpcomingContests[i].StartTime < resp_body.UpcomingContests[j].StartTime
    })
	// Sort CompletedContests by StartTime and Duration
    sort.Slice(resp_body.CompletedContests, func(i, j int) bool {
        if resp_body.CompletedContests[i].StartTime == resp_body.CompletedContests[j].StartTime {
            return resp_body.CompletedContests[i].Duration < resp_body.CompletedContests[j].Duration
        }
        return resp_body.CompletedContests[i].StartTime > resp_body.CompletedContests[j].StartTime
    })
	resp_body.Status = "success"
	return resp_body
}

func FetchAllContests() (models.ResponseBody, models.ApiError) {
	log.Println("Fetch all contests")
	currentDatetime := time.Now()
	codeChefcontests, apiError := GetCodeChefContests(currentDatetime)
	if apiError != (models.ApiError{}) {
		return models.ResponseBody{}, apiError
	}
	codeforcesContests, apiError := GetCodeforceContests(currentDatetime)
	if apiError != (models.ApiError{}) {
		return models.ResponseBody{}, apiError
	}
	leetcodeContests, apiError := GetLeetcodeContests(currentDatetime)
	if apiError != (models.ApiError{}) {
		return models.ResponseBody{}, apiError
	}
	var allContests []models.Contest
	allContests = append(allContests, codeChefcontests.AllContests...)
	allContests = append(allContests, codeforcesContests.AllContests...)
	allContests = append(allContests, leetcodeContests.AllContests...)
	resp_body := getResponseBody(allContests)
	return resp_body, models.ApiError{}
}

func FetchContests(c *gin.Context) (models.ResponseBody, models.ApiError) {
	log.Println("Fetch contests")
	currentDatetime := time.Now()
	codeChefcontests, apiError := GetCodeChefContests(currentDatetime)
	if apiError != (models.ApiError{}) {
		return models.ResponseBody{}, apiError
	}
	codeforcesContests, apiError := GetCodeforceContests(currentDatetime)
	if apiError != (models.ApiError{}) {
		return models.ResponseBody{}, apiError
	}
	leetcodeContests, apiError := GetLeetcodeContests(currentDatetime)
	if apiError != (models.ApiError{}) {
		return models.ResponseBody{}, apiError
	}
	var allContests []models.Contest
	allContests = append(allContests, codeChefcontests.AllContests...)
	allContests = append(allContests, codeforcesContests.AllContests...)
	allContests = append(allContests, leetcodeContests.AllContests...)
	resp_body := getResponseBody(allContests)
	return resp_body, models.ApiError{}
}