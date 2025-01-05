package models

type DurationRange struct {
	MinMinutes interface{} `json:"min_minutes"`
	MaxMinutes interface{} `json:"max_minutes"`
}

type ByStartTime struct {
	Ongoing interface{} `json:"ongoing"`
	Upcoming interface{} `json:"upcoming"`
	Completed interface{} `json:"completed"`
}

type SortOrder struct {
	ByStartTime `json:"by_start_time"`
	ByDuration interface{} `json:"by_duration"`
}

type RequestBody struct {
	Usage string `json:"usage"`
	Platforms []string `json:"platforms"`
	Phases []string `json:"phases"`
	FromDateTime interface{} `json:"from_date_time"`
	ToDateTime interface{} `json:"to_date_time"`
	DurationRange `json:"duration_range"`
	SortOrder  `json:"sort_order"`
}

type ApiError struct {
	Message string `json:"message"`
	Error string `json:"error"`
	StatusCode int `json:"status_code"`
}

type Contest struct {
	Platform string `json:"platform"`
	Id string `json:"id"`
	Name string `json:"name"`
	URL string `json:"url"`
	StartTime string `json:"start_time"`
	EndTime string `json:"end_time"`
	Duration int `json:"duration"`
}

type ServiceContests struct {
	Status string `json:"status"`
	AllContests []Contest `json:"all_contests"`
}

type ResponseBody struct {
	Status string `json:"status"`
	OngoingContests []Contest `json:"ongoing_contests"`
	UpcomingContests []Contest `json:"upcoming_contests"`
	CompletedContests []Contest `json:"completed_contests"`
}

type ResponseBodyError struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Error string `json:"error"`
}
