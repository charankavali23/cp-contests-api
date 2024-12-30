package models

type LeetcodeContestDetails struct {
	Title string `json:"title"`
	TitleSlug string `json:"titleSlug"`
	StartTime int `json:"startTime"`
	Duration int `json:"duration"`
}

type LeetcodeContests struct {
	Data struct {
		AllContests []LeetcodeContestDetails `json:"allContests"`
	} `json:"data"`
}