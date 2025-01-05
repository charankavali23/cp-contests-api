package models

type DurationRange struct {
	MinMinutes uint `json:"min_minutes"`
	MaxMinutes uint `json:"max_minutes"`
}

type ByPhase struct {
	Ongoing string `json:"ongoing"`
	Upcoming string `json:"upcoming"`
	Completed string `json:"completed"`
}

type SortOrder struct {
	ByDuration string `json:"by_duration"`
	ByPhase `json:"by_phase"`
}

type RequestBody struct {
	Platforms []string `json:"platforms"`
	Phases []string `json:"phases"`
	FromDateTime string `json:"from_date_time"`
	ToDateTime string `json:"to_date_time"`
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
