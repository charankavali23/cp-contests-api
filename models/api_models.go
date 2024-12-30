package models

type RequestBody struct {
	Platforms []string `json:"services"`
	Phase []string `json:"phase"`
	DurationRange []int `json:"duration_range"`
	FromDate string `json:"from_date"`
	ToDate string `json:"to_date"`
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
	OngoingContests []Contest `json:"ongoing_contests"`
	UpcomingContests []Contest `json:"upcoming_contests"`
	CompletedContests []Contest `json:"completed_contests"`
}

type ResponseJSON struct {
	Status string `json:"status"`
	ServiceContests
}

type ResponseJSONError struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Error string `json:"error"`
}
