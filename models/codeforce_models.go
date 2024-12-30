package models

type CodeforcesContestDetails struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Phase string `json:"phase"`
	Frozen bool `json:"frozen"`
	DurationSeconds int `json:"durationSeconds"`
	StartTimeSeconds int `json:"startTimeSeconds"`
	RelativeTimeSeconds int `json:"relativeTimeSeconds"`
}

type CodeforcesContests struct {
	Status string `json:"status"`
	Result []CodeforcesContestDetails `json:"result"`
}