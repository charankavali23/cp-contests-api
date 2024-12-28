package models

type Contest struct {
	Platform string `json:"platform"`
	Id string `json:"id"`
	Name string `json:"name"`
	StartTime string `json:"start_time"`
	Duration int `json:"duration"`
}

type RequestBody struct {
	Platforms []string `json:"services"`
	Phase []string `json:"phase"`
	DurationRange []int `json:"duration_range"`
	FromDate string `json:"from_date"`
	ToDate string `json:"to_date"`
}