package services

import (
	"log"
	"sort"
	"time"
	"net/http"

	"github.com/charankavali23/cp-contests-api/models"
	"github.com/charankavali23/cp-contests-api/utils"

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

var defaultRequestBody = models.RequestBody{
    Platforms: []string{
		"codechef",
        "codeforces",
        "leetcode",
	},
    Phases: []string{
		"ongoing",
        "upcoming",
        "completed",
	},
    FromDateTime: "2000-01-01T00:00:00+05:30",
    ToDateTime: "2099-12-31T23:59:59+05:30",
    DurationRange: models.DurationRange{
        MinMinutes: 0.0,
        MaxMinutes: 129600.0,
    },
    SortOrder: models.SortOrder{
        ByStartTime: models.ByStartTime{
            Ongoing: "desc",
            Upcoming: "asc",
            Completed: "desc",
        },
		ByDuration: "asc",
    },
}

var platformsFunc = map[string]func()(models.ServiceContests, models.ApiError){
	"codechef": GetCodeChefContests,
	"codeforces": GetCodeforcesContests,
	"leetcode": GetLeetcodeContests,
}

func UpdateRequestBody(requestBody *models.RequestBody) {
	if requestBody.FromDateTime == nil {
		requestBody.FromDateTime = defaultRequestBody.FromDateTime
	}
	if requestBody.ToDateTime == nil {
		requestBody.ToDateTime = defaultRequestBody.ToDateTime
	}
	if requestBody.DurationRange.MinMinutes == nil {
		requestBody.DurationRange.MinMinutes = defaultRequestBody.DurationRange.MaxMinutes
	}
	if requestBody.DurationRange.MaxMinutes == nil {
		requestBody.DurationRange.MaxMinutes = defaultRequestBody.DurationRange.MaxMinutes
	}
	if requestBody.SortOrder.ByDuration == nil {
		requestBody.SortOrder.ByDuration = defaultRequestBody.SortOrder.ByDuration
	}
	if requestBody.SortOrder.ByStartTime.Ongoing == nil {
		requestBody.SortOrder.ByStartTime.Ongoing = defaultRequestBody.SortOrder.ByStartTime.Ongoing
	}
	if requestBody.SortOrder.ByStartTime.Upcoming == nil {
		requestBody.SortOrder.ByStartTime.Upcoming = defaultRequestBody.SortOrder.ByStartTime.Upcoming
	}
	if requestBody.SortOrder.ByStartTime.Completed == nil {
		requestBody.SortOrder.ByStartTime.Completed = defaultRequestBody.SortOrder.ByStartTime.Completed
	}
}

func compareContestsByDuration(contestA models.Contest, contestB models.Contest, byDuration interface{}) bool {
	if byDuration == "" {
		return contestA.Duration < contestB.Duration
	} else if byDuration.(string) == "asc" {
		return contestA.Duration < contestB.Duration
	}
	return contestA.Duration > contestB.Duration
}

func compareContestsByStartTime(contestA models.Contest, contestB models.Contest, byStartTime interface{}, byDuration interface{}) bool {
	if byStartTime == "" {	
		return compareContestsByDuration(contestA, contestB, byDuration)
	} else if byStartTime.(string) == "asc" {
		if contestA.StartTime == contestB.StartTime {
			return compareContestsByDuration(contestA, contestB, byDuration)
		}
		return contestA.StartTime < contestB.StartTime
	} else if byStartTime.(string) == "desc" {
		if contestA.StartTime == contestB.StartTime {
			return compareContestsByDuration(contestA, contestB, byDuration)
		}
		return contestA.StartTime > contestB.StartTime
	}
	return compareContestsByDuration(contestA, contestB, byDuration)
}

func getResponseBody(allContests []models.Contest, requestBody models.RequestBody) (models.ResponseBody) {
	var resp_body models.ResponseBody
	currentDatetime := time.Now().Format(time.RFC3339)
	for _, contest := range allContests {
		if (requestBody.FromDateTime.(string) <= contest.StartTime && contest.StartTime <= requestBody.ToDateTime.(string)) && (uint(requestBody.DurationRange.MinMinutes.(float64)) <= uint(contest.Duration)) && uint(contest.Duration) <= uint(requestBody.DurationRange.MaxMinutes.(float64)){
			if utils.IsAvailable("ongoing", requestBody.Phases) && contest.StartTime <= currentDatetime && currentDatetime <= contest.EndTime {
				resp_body.OngoingContests = append(resp_body.OngoingContests, contest)
			} else if utils.IsAvailable("upcoming", requestBody.Phases) && currentDatetime < contest.StartTime {
				resp_body.UpcomingContests = append(resp_body.UpcomingContests, contest)
			} else if utils.IsAvailable("completed", requestBody.Phases) && currentDatetime > contest.EndTime {
				resp_body.CompletedContests = append(resp_body.CompletedContests, contest)
			}
		}
	}
	// Sort OngoingContests by StartTime and Duration
	if !(requestBody.SortOrder.ByStartTime.Ongoing == "" && requestBody.SortOrder.ByDuration == "") {
		log.Println("Sort OngoingContests by StartTime and Duration")
		sort.Slice(resp_body.OngoingContests, func(i, j int) bool {
			return compareContestsByStartTime(resp_body.OngoingContests[i], resp_body.OngoingContests[j], requestBody.SortOrder.ByStartTime.Ongoing, requestBody.SortOrder.ByDuration)
		})
	}
    // Sort UpcomingContests by StartTime and Duration
	if !(requestBody.SortOrder.ByStartTime.Upcoming == "" && requestBody.SortOrder.ByDuration == "") {
		sort.Slice(resp_body.UpcomingContests, func(i, j int) bool {
			return compareContestsByStartTime(resp_body.UpcomingContests[i], resp_body.UpcomingContests[j], requestBody.SortOrder.ByStartTime.Upcoming, requestBody.SortOrder.ByDuration)
		})
	}
	// Sort CompletedContests by StartTime and Duration
	if !(requestBody.SortOrder.ByStartTime.Completed == "" && requestBody.SortOrder.ByDuration == "") {
		sort.Slice(resp_body.CompletedContests, func(i, j int) bool {
			return compareContestsByStartTime(resp_body.CompletedContests[i], resp_body.CompletedContests[j], requestBody.SortOrder.ByStartTime.Completed, requestBody.SortOrder.ByDuration)
		})
	}
	resp_body.Status = "success"
	return resp_body
}

func fetchPlatformsContests(platforms []string) ([]models.Contest, models.ApiError) {
	type result struct {
		Contests models.ServiceContests
		ApiError models.ApiError
	}
	resultsChan := make(chan result, len(platforms))
	for _, platform := range platforms {
		go func(fetcher func() (models.ServiceContests, models.ApiError)) {
			contests, apiError := fetcher()
			resultsChan <- result{
				Contests: contests,
				ApiError: apiError,
			}
		}(platformsFunc[platform])
	}
	var allContests []models.Contest
	for i := 0; i < len(platforms); i++ {
		res := <- resultsChan
		if res.ApiError != (models.ApiError{}) {
			return nil, res.ApiError
		}
		allContests = append(allContests, res.Contests.AllContests...)
	}
	return allContests, models.ApiError{}
}

func FetchAllContests() (models.ResponseBody, models.ApiError) {
	log.Println("Fetch all contests")
	allContests, apiError := fetchPlatformsContests(defaultRequestBody.Platforms)
	if apiError != (models.ApiError{}) {
		return models.ResponseBody{}, models.ApiError{}
	}
	resp_body := getResponseBody(allContests, defaultRequestBody)
	return resp_body, models.ApiError{}
}

func FetchContests(c *gin.Context) (models.ResponseBody, models.ApiError) {
	log.Println("Fetch contests")
	validateRequestBody, check := c.Get("validatedRequestBody")
	if !check {
		return models.ResponseBody{}, utils.NewApiError("Internal Server Error", "requestBody doesn't exits in context", http.StatusInternalServerError)
	}
	requestBody := validateRequestBody.(models.RequestBody)
	UpdateRequestBody(&requestBody)
	log.Printf("Updated Request body: %+v", requestBody)
	allContests, apiError := fetchPlatformsContests(requestBody.Platforms)
	if apiError != (models.ApiError{}) {
		return models.ResponseBody{}, models.ApiError{}
	}
	resp_body := getResponseBody(allContests, requestBody)
	return resp_body, models.ApiError{}
}